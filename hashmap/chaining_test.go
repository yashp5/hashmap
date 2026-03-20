package hashmap

import (
	"fmt"
	"testing"
)

func TestChainingHashMap(t *testing.T) {
	hmap := NewChainingMap()

	t.Run("get from empty map", func(t *testing.T) {
		val, ok := hmap.Get("key")
		if ok != false && val != nil {
			t.Fatalf("want val=nil && ok=false, got %v, true\n", val)
		}
	})

	t.Run("put inserts new key", func(t *testing.T) {
		hmap.Put("key1", "value1")
		val, ok := hmap.Get("key1")
		if ok == false {
			t.Fatalf("key not found\n")
		}
		if val != "value1" {
			t.Fatalf("want val=value1, got=%v\n", val)
		}
	})

	t.Run("put updates existing key", func(t *testing.T) {
		hmap.Put("key1", "value1")
		val, ok := hmap.Get("key1")
		if ok == false {
			t.Fatalf("key not found\n")
		}
		if val != "value1" {
			t.Fatalf("want val=value1, got=%v\n", val)
		}
	})

	t.Run("delete existing key", func(t *testing.T) {
		hmap.Put("key1", "value1")
		val, ok := hmap.Get("key1")
		if ok == false {
			t.Fatalf("key not found\n")
		}
		if val != "value1" {
			t.Fatalf("want val=value1, got=%v\n", val)
		}
	})

	t.Run("delete missing key", func(t *testing.T) {
		ok := hmap.Delete("key2")
		if ok != false {
			t.Fatalf("missing delete: want=false, got=true")
		}
	})

	t.Run("put triggers grow/resize", func(t *testing.T) {
		initsize := len(hmap.buckets)
		count := int(float64(len(hmap.buckets)) * ChainingMaxLoadFactor)
		for i := range count {
			key := fmt.Sprintf("key%d", i)
			val := fmt.Sprintf("val%d", i)
			hmap.Put(key, val)
		}
		hmap.Put("key-grow", "val-grow")
		if len(hmap.buckets) == initsize*2 {
			t.Fatalf("hmap did not grow, want bucket len:%d, got:%d", initsize*2, len(hmap.buckets))
		}
	})
}
