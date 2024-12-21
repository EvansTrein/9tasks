package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceExample(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name          string
		args          args
		wantEvenSlice []int
	}{
		{
			name:          "четные числа",
			args:          args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			wantEvenSlice: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEvenSlice := SliceExample(tt.args.s); !reflect.DeepEqual(gotEvenSlice, tt.wantEvenSlice) {
				t.Errorf("SliceExample() = %v, want %v", gotEvenSlice, tt.wantEvenSlice)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	type args struct {
		s   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "добавить 666",
			args: args{[]int{1, 2, 3}, 666},
			want: []int{1, 2, 3, 666},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddElements(tt.args.s, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "копирование",
			args: args{[]int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CopySlice(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopySlice() = %v, want %v", got, tt.want)
			}
			origPtr := *(*unsafe.Pointer)(unsafe.Pointer(&tt.args.s))
			gotPtr := *(*unsafe.Pointer)(unsafe.Pointer(&got))
			if origPtr == gotPtr {
				t.Errorf("Слайсы имеют один и тот же адрес памяти")
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args struct {
		s    []int
		indx int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "удаление по индексу",
			args:    args{[]int{1, 2, 3, 4}, 0},
			want:    []int{2, 3, 4},
			wantErr: false,
		},
		{
			name:    "удаление по несуществующему индексу",
			args:    args{[]int{1, 2, 3, 4}, 10},
			want:    []int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveElement(tt.args.s, tt.args.indx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
