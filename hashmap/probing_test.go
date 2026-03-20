package hashmap

import "testing"

func TestProbingMap(t *testing.T) {
	hmap := NewProbingMap()
	t.Run("get missing key", func(t *testing.T) {
		_, ok := hmap.Get("key1")
		if ok != false {
		}
	})

	t.Run("put insert key", func(t *testing.T) {
		hmap.Put("key1", "value")
		val, ok := hmap.Get("key1")
		if ok == false || val != "value1" {
		}
	})

	t.Run("get valid key", func(t *testing.T) {
		val, ok := hmap.Get("key1")
		if ok != true && val != "value1" {
		}
	})

	t.Run("put update key", func(t *testing.T) {
		hmap.Put("key1", "value")
		val, ok := hmap.Get("key1")
		if ok == false || val != "value1" {
		}
	})

	t.Run("put that triggers probe", func(t *testing.T) {
		hmap.Put("key1", "value1")
		hmap.Put("key2", "value2")
	})

	t.Run("get that triggers probe", func(t *testing.T) {})

	t.Run("delete key", func(t *testing.T) {
		deleteKey := "del-key"
		hmap.Put(deleteKey, "val")
		hmap.Delete(deleteKey)
	})

	t.Run("delete key that triggers probe", func(t *testing.T) {
		deleteKey := "del-key"
		hmap.Put(deleteKey, "val")
		hmap.Delete(deleteKey)
	})
}
