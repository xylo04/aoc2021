package depth

import "testing"

func Test_Increases(t *testing.T) {
	type args struct {
		d []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "website",
			args: args{d: []uint{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Increases(tt.args.d); got != tt.want {
				t.Errorf("Increases() = %v, want %v", got, tt.want)
			}
		})
	}
}
