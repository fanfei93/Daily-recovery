package main

import "testing"

func Test_freqAlphabets(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "10#11#12"}, want: "jkab"},
		{name: "2", args: args{s: "1326#"}, want: "acz"},
		{name: "3", args: args{s: "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"}, want: "abcdefghijklmnopqrstuvwxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := freqAlphabets(tt.args.s); got != tt.want {
				t.Errorf("freqAlphabets() = %v, want %v", got, tt.want)
			}
		})
	}
}
