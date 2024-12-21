package main

import (
	"reflect"
	"testing"
)

func TestSetDifference(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name          string
		args          args
		wantDiffSlice []string
	}{
		{
			name:          "разность множеств",
			args:          args{[]string{"one", "two", "three", "four", "five"}, []string{"four", "five"}},
			wantDiffSlice: []string{"one", "two", "three"},
		},
		{
			name:          "разность множеств 2",
			args:          args{[]string{"case1", "case2", "case3", "case4"}, []string{"case1", "case4", "case10"}},
			wantDiffSlice: []string{"case2", "case3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiffSlice := SetDifference(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(gotDiffSlice, tt.wantDiffSlice) {
				t.Errorf("SetDifference() = %v, want %v", gotDiffSlice, tt.wantDiffSlice)
			}
		})
	}
}

func TestSetDifferenceInvalid(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	diff := SetDifference(slice1, slice2)
	wantInvResult := []string{"apple", "cherry", "43", "fig"} // некорректный результат

	if reflect.DeepEqual(diff, wantInvResult) {
		t.Errorf("ожидаемый некорректный результат: %v, но получено: %v", wantInvResult, diff)
	}
}
