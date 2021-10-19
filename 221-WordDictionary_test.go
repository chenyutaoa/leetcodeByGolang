package leetcode

import "testing"

func Test_WordDictionaryConstructor(t *testing.T) {
	obj := WordDictionaryConstructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	if obj.Search("pad") {
		t.Fatal("not fount pad")
	}
	if !obj.Search("bad") {
		t.Fatal("exist bad")
	}
	if !obj.Search(".ad") {
		t.Fatal("exist .ad")
	}
	if !obj.Search("bad") {
		t.Fatal("exist b..")
	}
	obj = WordDictionaryConstructor()
	obj.AddWord("a")
	obj.AddWord("ab")
	if !obj.Search("a") {
		t.Fatal("exist a")
	}
	if !obj.Search("a.") {
		t.Fatal("exist bad")
	}
	if !obj.Search("ab") {
		t.Fatal("exist bad")
	}
}
