package main

import "testing"

func Test_canReach(t *testing.T) {
	type args struct {
		arr   []int
		start int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name:"1",args:args{arr:[]int{4,2,3,0,3,1,2},start:5},want:true},
		{name:"2",args:args{arr:[]int{4,2,3,0,3,1,2},start:0},want:true},
		{name:"3",args:args{arr:[]int{3,0,2,1,2},start:2},want:false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.arr, tt.args.start); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
