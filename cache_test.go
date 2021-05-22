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

func TestLRUPolicy(t *testing.T) {
	cache := NewCache(5, LRU)
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

	cache.Put("foo6", "bar6") // foo3 should be removed

	if value, err := cache.Get("foo3"); value != nil || err == nil {
		t.Error("value should be nil")
	}

	if value, err := cache.Get("foo"); *value != "bar" || err != nil {
		t.Errorf("value should be bar, but we got %s", *value)
	}

	if value, err := cache.Get("foo6"); *value != "bar6" || err != nil {
		t.Errorf("value should be bar6, but we got %s", *value)
	}

	cache.Put("foo7", "bar7")

	if value, err := cache.Get("foo4"); value != nil || err == nil {
		t.Error("value should be nil")
	}
	if value, err := cache.Get("foo7"); *value != "bar7" || err != nil {
		t.Errorf("value should be bar7, but we got %s", *value)
	}
}

func TestLFUPolicy(t *testing.T) {
	cache := NewCache(2, LFU)
	cache.Put("foo1", "bar1")
	cache.Put("foo2", "bar2")

	if value, err := cache.Get("foo1"); *value != "bar1" || err != nil {
		t.Errorf("value should be bar, but we got %s", *value)
	}

	cache.Put("foo3", "bar3")

	if value, err := cache.Get("foo2"); value != nil || err == nil {
		t.Error("value should be nil")
	}

}

func TestClockPolicy(t *testing.T) {
	cache := NewCache(5, CLOCK)
	cache.Put("foo1", "bar1")
	cache.Put("foo2", "bar2")
	cache.Put("foo3", "bar3")
	cache.Put("foo4", "bar4")
	cache.Put("foo5", "bar5")

	if value, err := cache.Get("foo1"); *value != "bar1" || err != nil {
		t.Errorf("value should be bar1, but we got %s", *value)
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

	cache.Put("foo6", "bar6") // foo1 is evicted

	if value, err := cache.Get("foo1"); value != nil || err == nil {
		t.Error("value should be nil")
	}

	if value, err := cache.Get("foo2"); *value != "bar2" || err != nil {
		t.Errorf("value should be bar2, but we got %s", *value)
	}

	if value, err := cache.Get("foo3"); *value != "bar3" || err != nil {
		t.Errorf("value should be bar3, but we got %s", *value)
	}

	cache.Put("foo7", "bar7") // foo4 is evicted

	if value, err := cache.Get("foo4"); value != nil || err == nil {
		t.Error("value should be nil")
	}
}
