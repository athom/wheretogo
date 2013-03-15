package wheretogo

import (
	"testing"
)

func TestRouter(t *testing.T) {
	key1 := "1234"
	key2 := "5678"

	path1 := "/white_house"
	path2 := "/black_land"

	AddRoute(key1, path1)
	AddRoute(key2, path2)

	_, ok := DefaultRouter.Pool[key1]
	if !ok {
		t.Errorf("shold has key: " + key1)
	}

	r1 := Where(key1)
	if r1 != path1 {
		t.Errorf("shold redirect to: " + path1)
	}

	_, ok = DefaultRouter.Pool[key1]
	if ok {
		t.Errorf("shold delete key: " + key1)
	}

	_, ok = DefaultRouter.Pool[key2]
	if !ok {
		t.Errorf("shold has key: " + key2)
	}

	r2 := Where(key2)
	if r2 != path2 {
		t.Errorf("shold redirect to: " + path2)
	}

}
