use super::{single_value_size_tracker::SingleValueSizeTracker, BlockKeyArrowBuilder};
use crate::key::{CompositeKey, KeyWrapper};
use arrow::{
    array::{Array, ArrayRef, Int32Array, Int32Builder, ListBuilder},
    datatypes::Field,
};
use parking_lot::RwLock;
use std::{collections::BTreeMap, sync::Arc};

#[derive(Clone, Debug)]
pub struct Int32ArrayStorage {
    pub(crate) storage: Arc<RwLock<BTreeMap<CompositeKey, Int32Array>>>,
    size_tracker: SingleValueSizeTracker,
}

impl Int32ArrayStorage {
    pub(in crate::arrow) fn new() -> Self {
        Self {
            storage: Arc::new(RwLock::new(BTreeMap::new())),
            size_tracker: SingleValueSizeTracker::new(),
        }
    }

    pub(super) fn get_prefix_size(&self) -> usize {
        return self.size_tracker.get_prefix_size();
    }

    pub(super) fn get_key_size(&self) -> usize {
        return self.size_tracker.get_key_size();
    }

    pub(super) fn get_value_size(&self) -> usize {
        return self.size_tracker.get_value_size();
    }

    /// The count of the total number of values in the storage across all arrays.
    pub(super) fn total_value_count(&self) -> usize {
        let storage = self.storage.read();
        storage.iter().fold(0, |acc, (_, value)| acc + value.len())
    }

    pub(super) fn split(&self, prefix: &str, key: KeyWrapper) -> Int32ArrayStorage {
        let mut storage_guard = self.storage.write();
        let split = storage_guard.split_off(&CompositeKey {
            prefix: prefix.to_string(),
            key,
        });
        Int32ArrayStorage {
            storage: Arc::new(RwLock::new(split)),
            // TODO: make this correct
            size_tracker: SingleValueSizeTracker::new(),
        }
    }

    pub(super) fn build_keys(&self, builder: BlockKeyArrowBuilder) -> BlockKeyArrowBuilder {
        let storage = self.storage.read();
        // TODO: mut ref instead of ownership of builder
        let mut builder = builder;
        for (key, _) in storage.iter() {
            builder.add_key(key.clone());
        }
        builder
    }

    pub(super) fn len(&self) -> usize {
        let storage = self.storage.read();
        storage.len()
    }

    pub(super) fn to_arrow(&self) -> (Field, ArrayRef) {
        let item_capacity = self.storage.read().len();
        let mut value_builder;
        if item_capacity == 0 {
            value_builder = ListBuilder::new(Int32Builder::new());
        } else {
            value_builder = ListBuilder::with_capacity(
                Int32Builder::with_capacity(self.total_value_count()),
                item_capacity,
            );
        }

        let storage = self.storage.read();
        for (_, value) in storage.iter() {
            value_builder.append_value(value);
        }

        let value_field = Field::new(
            "value",
            arrow::datatypes::DataType::List(Arc::new(Field::new(
                "item",
                arrow::datatypes::DataType::Int32,
                true,
            ))),
            true,
        );
        let value_arr = value_builder.finish();
        (
            value_field,
            (&value_arr as &dyn Array).slice(0, value_arr.len()),
        )
    }
}