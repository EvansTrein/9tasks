package main

import (
	"reflect"
	"testing"
)

func TestSetIntersection(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name         string
		args         args
		wantIsInter  bool
		wantInterSet []int
	}{
		{
			name:         "пересечение множеств",
			args:         args{[]int{1, 2, 3, 4, 5}, []int{1, 2}},
			wantIsInter:  true,
			wantInterSet: []int{1, 2},
		},
		{
			name:         "пересечение множеств 2",
			args:         args{[]int{1, 1, 1, 2, 3, 4, 5, 5}, []int{1, 1, 2, 5, 5}},
			wantIsInter:  true,
			wantInterSet: []int{1, 1, 2, 5, 5},
		},
		{
			name:         "нет пересечения",
			args:         args{[]int{1, 2, 3, 4}, []int{5, 6, 7}},
			wantIsInter:  false,
			wantInterSet: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsInter, gotInterSet := SetIntersection(tt.args.slice1, tt.args.slice2)
			if gotIsInter != tt.wantIsInter {
				t.Errorf("SetIntersection() gotIsInter = %v, want %v", gotIsInter, tt.wantIsInter)
			}
			if !reflect.DeepEqual(gotInterSet, tt.wantInterSet) && len(gotInterSet) != 0 && len(tt.wantInterSet) != 0 {
				t.Errorf("SetIntersection() gotInterSet = %v, want %v", gotInterSet, tt.wantInterSet)
			}
		})
	}
}
