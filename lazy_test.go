package lazy

import (
	"testing"
)

func Test_lazy(t *testing.T) {
	count := 0
	lazy := NewLazy(func() interface{} {
		count++;
		return "x"
	})
	x := lazy.Value()
	if x != "x" {
		t.Error("doens't work, 1st")
	}
	x = lazy.Value()
	if r, ok := x.(string); ok && ( r != "x" || count > 1 ) {
		t.Error("doens't work 2nd")
	}
	x = lazy.Value()
	if r, ok := x.(string); ok && ( r != "x" || count > 1 ) {
		t.Error("doens't work 2nd")
	}

	for w:=0 ; w <= 1000000 ; w++ {
		go func(){
			x = lazy.Value()
			if r, ok := x.(string); ok && ( r != "x" || count > 1 ) {
				t.Error("doens't work 2nd")
			}
		}()
	}

	if count > 1 {
		t.Errorf("bad count %v", count)
	}
}

