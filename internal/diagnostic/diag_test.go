package diagnostic

import "testing"

func TestReport(t *testing.T) {
	type args struct {
		rep string
	}
	type returns struct {
		gamma   uint
		epsilon uint
	}
	tests := []struct {
		name string
		args args
		want returns
	}{
		{
			name: "website",
			args: args{rep: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`},
			want: returns{gamma: 22, epsilon: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotG, gotE := Report(tt.args.rep)
			got := returns{gotG, gotE}
			if got != tt.want {
				t.Errorf("Report() got = %v, want %v", got, tt.want)
			}
		})
	}
}
