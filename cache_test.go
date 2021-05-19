package cache

import "testing"

func TestFIFOPolicy(t *testing.T) {
	cache := NewCache(5, FIFO)
	cache.Put("foo", "bar")
	cache.Put("foo2", "bar2")
	cache.Put("foo3", "bar3")
	cache.Put("foo4", "bar4")
	cache.Put("foo5", "bar5")

	if value, err := cache.Get("foo"); *value != "bar" || err != nil {
		t.Errorf("value should be bar, but we got %s", *value)
	}

	if value, err := cache.Get("foo2"); *value != "bar2" || err != nil {
		t.Errorf("value should be bar2, but we got %s", *value)
	}

	if value, err := cache.Get("foo3"); *value != "bar3" || err != nil {
		t.Errorf("value should be bar3, but we got %s", *value)
	}

	if value, err := cache.Get("foo4"); *value != "bar4" || err != nil {
		t.Errorf("value should be bar4, but we got %s", *value)
	}

	if value, err := cache.Get("foo5"); *value != "bar5" || err != nil {
		t.Errorf("value should be bar5, but we got %s", *value)
	}

	cache.Put("foo6", "bar6")

	if value, err := cache.Get("foo"); value != nil || err == nil {
		t.Error("value should be nil")
	}

	if value, err := cache.Get("foo6"); *value != "bar6" || err != nil {
		t.Errorf("value should be bar6, but we got %s", *value)
	}

	cache.Put("foo7", "bar7")

	if value, err := cache.Get("foo"); value != nil || err == nil {
		t.Error("value should be nil")
	}

	if value, err := cache.Get("foo"); value != nil || err == nil {
		t.Error("value should be nil")
	}

	if value, err := cache.Get("foo7"); *value != "bar7" || err != nil {
		t.Errorf("value should be bar7, but we got %s", *value)
	}
}
