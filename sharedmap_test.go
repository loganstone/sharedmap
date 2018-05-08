package sharedmap

import (
	"testing"
)

func testSharedMapGet(t *testing.T, sm SharedMap, key, val interface{}) {
	value, ok := sm.Get(key)

	if ok {
		if value != val {
			t.Errorf("%v != %v", value, val)
		}
	} else {
		t.Errorf("!ok(%t)", ok)
	}
}

func TestSharedMapSet(t *testing.T) {
	sm := New()
	sm.Set("foo", 100)
	sm.Set(1, "100")

	testSharedMapGet(t, sm, "foo", 100)
	testSharedMapGet(t, sm, 1, "100")
}

func TestSharedMapGet(t *testing.T) {
	sm := New()
	sm.Set("foo", "bar")

	testSharedMapGet(t, sm, "foo", "bar")
}

func TestSharedMapRemove(t *testing.T) {
	sm := New()
	sm.Set("foo", "bar")

	sm.Remove("foo")

	value, ok := sm.Get("foo")
	if value != nil || ok {
		t.Errorf("value(%#v), ok(%t)", value, ok)
	}
}

func TestSharedMapCount(t *testing.T) {
	sm := New()
	sm.Set("foo1", "bar1")
	sm.Set("foo2", "bar2")

	if 2 != sm.Count() {
		t.Errorf("%d != %d", 2, sm.Count())
	}
}
