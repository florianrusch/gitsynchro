package internal_test

import (
	"testing"

	"github.com/florianrusch/gitsynchro/internal"
)

func TestNestedJoin(t *testing.T) {
	t.Parallel()

	type T2 struct {
		theString string
	}

	type T struct {
		nestedStruct T2
	}

	type S2 struct {
		theSecondString string
	}

	type S struct {
		theSecondNestedStruct S2
	}

	want := "a,b,c"
	joinChar := ","

	objT := []T{
		{nestedStruct: T2{theString: "a"}},
		{nestedStruct: T2{theString: "b"}},
		{nestedStruct: T2{theString: "c"}},
	}
	fnT := func(t T) string {
		return t.nestedStruct.theString
	}

	t.Run("First structs", func(t *testing.T) {
		t.Parallel()

		if got := internal.NestedJoin(joinChar, objT, fnT); got != want {
			t.Errorf("NestedJoin() = %v, want %v", got, want)
		}
	})

	objS := []S{
		{theSecondNestedStruct: S2{theSecondString: "a"}},
		{theSecondNestedStruct: S2{theSecondString: "b"}},
		{theSecondNestedStruct: S2{theSecondString: "c"}},
	}
	fnS := func(t S) string {
		return t.theSecondNestedStruct.theSecondString
	}

	t.Run("Second structs", func(t *testing.T) {
		t.Parallel()

		if got := internal.NestedJoin(joinChar, objS, fnS); got != want {
			t.Errorf("NestedJoin() = %v, want %v", got, want)
		}
	})
}
