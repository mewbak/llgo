package main

import (
	"testing"
)

func _TestStaticStructInterfaceConversion(t *testing.T) {
	err := runAndCheckMain(testdata("interface.go"), checkStringsEqual)
	if err != nil {
		t.Fatal(err)
	}
}

func _TestInterfaceToInterfaceConversion(t *testing.T) {
	err := runAndCheckMain(testdata("interface_i2i.go"), checkStringsEqual)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStaticBasicTypeToInterfaceConversion(t *testing.T) {
	err := runAndCheckMain(testdata("interfaces/basic.go"), checkStringsEqual)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInterfaceMethods(t *testing.T) {
	err := runAndCheckMain(testdata("interfaces/methods.go"), checkStringsEqual)
	if err != nil {
		t.Fatal(err)
	}
}

// vim: set ft=go:
