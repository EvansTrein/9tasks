package main

import "testing"

func TestDefinitionType(t *testing.T) {
	type args struct {
		variable interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "int",
			args: args{variable: 2},
			want: "Тип переменной: int",
		},
		{
			name: "float64",
			args: args{variable: 5.0},
			want: "Тип переменной: float64",
		},
		{
			name: "string",
			args: args{variable: "Golang"},
			want: "Тип переменной: string",
		},
		{
			name: "bool",
			args: args{variable: true},
			want: "Тип переменной: bool",
		},
		{
			name: "complex128",
			args: args{variable: 1 + 2i},
			want: "Тип переменной: complex128",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefinitionType(tt.args.variable); got != tt.want {
				t.Errorf("DefinitionType() = %v, want %v", got, tt.want)
			}
		})
	}
}
