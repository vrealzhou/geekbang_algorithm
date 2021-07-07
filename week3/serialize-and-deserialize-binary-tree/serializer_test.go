package main

import (
	"testing"
)

func TestParseString(t *testing.T) {
	str := "[1,2,3,null,null,4,5,null,null,null,null]"
	codec := Constructor()
	root := (&codec).deserialize(str)
	t.Log("deserialize done")
	serialized := (&codec).serialize(root)
	if str != serialized {
		t.Errorf("incorrect serialized: %s", serialized)
	}
}
