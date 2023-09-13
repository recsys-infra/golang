package slice_test

import (
	"testing"

	"github.com/recsys-infra/golang/slice"
)

func Test_Contains(t *testing.T) {
	if !slice.Contains([]string{"1", "2", "3"}, "3") {
		t.Fatalf("string slice failed")
	}

	if slice.Contains([]string{"1", "2", "3"}, "4") {
		t.Fatalf("strings not contains failed")
	}

	if !slice.Contains([]int{1, 2, 3}, 3) {
		t.Fatalf("int slice failed")
	}

	if slice.Contains([]int{1, 2, 3}, 4) {
		t.Fatalf("int slice failed")
	}
}
