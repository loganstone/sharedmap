package sharedmap

import (
	"testing"
)

type Expected struct {
	key   interface{}
	value interface{}
	exist bool
}

func testSharedMapGet(t *testing.T, sm *SharedMap, expected *Expected) {
	actual, ok := sm.Get(expected.key)

	if actual != expected.value || ok != expected.exist {
		err := "%v != %v|| ok(%t) != exist(%t)"
		t.Errorf(err, actual, expected.value, ok, expected.exist)
	}
}

func TestSharedMapSet(t *testing.T) {
	sm := New()
	sm.Set("foo", 100)
	sm.Set(1, "100")

	expected := &Expected{"foo", 100, true}
	testSharedMapGet(t, sm, expected)

	expected = &Expected{1, "100", true}
	testSharedMapGet(t, sm, expected)
}

func TestSharedMapGet(t *testing.T) {
	sm := New()
	sm.Set("foo", "bar")

	expected := &Expected{"foo", "bar", true}
	testSharedMapGet(t, sm, expected)
}

func TestSharedMapRemove(t *testing.T) {
	sm := New()
	sm.Set("foo", "bar")

	sm.Remove("foo")

	expected := &Expected{"foo", nil, false}
	testSharedMapGet(t, sm, expected)
}

func TestSharedMapCount(t *testing.T) {
	sm := New()
	sm.Set("foo1", "bar1")
	sm.Set("foo2", "bar2")

	actual := sm.Count()
	expected := 2

	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
