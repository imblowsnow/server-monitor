package utils

import "testing"

func TestCopyObject(t *testing.T) {
	type Source struct {
		Name string
		Age  int
	}
	type Dest struct {
		Name string
		Age  int
	}

	src := Source{Name: "Tom", Age: 18}
	var dest Dest
	CopyObject(src, &dest)

	if src.Name != dest.Name || src.Age != dest.Age {
		t.Errorf("CopyObject failed, src: %v, dest: %v", src, dest)
	}
}
