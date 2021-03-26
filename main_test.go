package main

import "testing"

func TestPrint01(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{"str1"}},
		{"case2", args{"str2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print01(tt.args.str)
		})
	}
}

func TestPrint02(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{"str1"}},
		{"case2", args{"str2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print02(tt.args.str)
		})
	}
}
