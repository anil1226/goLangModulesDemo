package buffers

import "testing"

func TestPrimeNumber(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "PRIME_NUMBER_7",
			args: args{count: 7},
			want: "2, 3, 5, 7, ",
		},
		{
			name: "PRIME_NUMBER_100",
			args: args{count: 100},
			want: "2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeNumber(tt.args.count); got != tt.want {
				t.Errorf("PrimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
