package pilot

import "testing"

func TestParse(t *testing.T) {
	type args struct {
		cmds []string
	}
	type returns struct {
		x uint
		y uint
	}
	tests := []struct {
		name string
		args args
		want returns
	}{
		{
			name: "website",
			args: args{cmds: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2"}},
			want: returns{15, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := Parse(tt.args.cmds)
			got := returns{gotX, gotY}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
