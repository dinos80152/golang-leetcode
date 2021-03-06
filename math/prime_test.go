package math

import "testing"

func TestIsPrimeSimple(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1}, false},
		{"2", args{2}, true},
		{"3", args{3}, true},
		{"4", args{4}, false},
		{"7", args{7}, true},
		{"10", args{10}, false},
		{"11", args{11}, true},
		{"13", args{13}, true},
		{"25", args{25}, false},
		{"8191", args{8191}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeSimple(tt.args.n); got != tt.want {
				t.Errorf("IsPrimeSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
