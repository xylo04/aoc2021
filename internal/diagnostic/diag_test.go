package diagnostic

import "testing"

func TestReport(t *testing.T) {
	type args struct {
		rep string
	}
	tests := []struct {
		name string
		args args
		want DiagReport
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
			want: DiagReport{
				Gamma:        22,
				Epsilon:      9,
				PowerConsume: 198,
				O2Gen:        23,
				CO2Scrub:     10,
				LifeSupport:  230,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcReport(tt.args.rep)
			if got != tt.want {
				t.Errorf("CalcReport() got = %v, want %v", got, tt.want)
			}
		})
	}
}
