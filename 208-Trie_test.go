package leetcode

import (
	"testing"
)

func TestConstructorTrie(t *testing.T) {
	obj := ConstructorTrie()
	obj.Insert("apple")
	exist := obj.Search("apple")
	if exist != true{
		t.Fatal()
	}
	exist = obj.Search("appl")
	if exist != false{
		t.Fatal()
	}
	exist= obj.StartsWith("appl")
	if exist != true{
		t.Fatal()
	}
	obj.Insert("appl")
	exist= obj.StartsWith("appl")
	if exist != true{
		t.Fatal()
	}
}
