package sharedmap

import (
	"testing"
)

func testSharedMapGet(t *testing.T, sm SharedMap, k, v interface{}, exist bool) {
	value, ok := sm.Get(k)

	if value != v || ok != exist {
		err := "value(%v) != v(%v)|| ok(%t) != exist(%t)"
		t.Errorf(err, value, v, ok, exist)
	}
}

func TestSharedMapSet(t *testing.T) {
	sm := New()
	sm.Set("foo", 100)
	sm.Set(1, "100")

	testSharedMapGet(t, sm, "foo", 100, true)
	testSharedMapGet(t, sm, 1, "100", true)
}

func TestSharedMapGet(t *testing.T) {
	sm := New()
	sm.Set("foo", "bar")

	testSharedMapGet(t, sm, "foo", "bar", true)
}

func TestSharedMapRemove(t *testing.T) {
	sm := New()
	sm.Set("foo", "bar")

	sm.Remove("foo")

	testSharedMapGet(t, sm, "foo", nil, false)
}

func TestSharedMapCount(t *testing.T) {
	sm := New()
	sm.Set("foo1", "bar1")
	sm.Set("foo2", "bar2")

	if 2 != sm.Count() {
		t.Errorf("%d != %d", 2, sm.Count())
	}
}
