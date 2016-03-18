package lazy

import "testing"

type Thing struct {
	Name string
}
var count = 0


func Test_LazyThing(t *testing.T){

	lazy:= NewLazy(func() interface{}{
		count++
		return &Thing{"x"}
	})

	if lazyThing, ok  := lazy.Value().(*Thing) ; !ok {
		t.Error("Can't convert to Thing")
	} else {
		if lazyThing.Name != "x" {
			t.Errorf("Bad Name %v", lazyThing.Name)
		}
		if count > 1 {
			t.Errorf("Too manycall %v", count)
		}
	}

	if lazyThing, ok  := lazy.Value().(*Thing) ; !ok {
		// nil ... Why ?
		t.Error("Can't convert to Thing")
	} else {
		if lazyThing.Name != "x" {
			t.Errorf("Bad Name %v", lazyThing.Name)
		}
		if count > 1 {
			t.Errorf("Too manycall %v", count)
		}
	}
	// Let it explode in pieces
	{
		lazyThing, _ := lazy.Value().(*Thing)
		// nil ... Why ?
		if lazyThing.Name != "x" {
			t.Error(lazyThing)
		}
		if count > 1 {
			t.Errorf("Too manycall %v", count)
		}
	}

}
