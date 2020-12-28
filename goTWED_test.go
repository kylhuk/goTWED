package goTWED

import "testing"

func TestTWED(t *testing.T) {
	type args struct {
		ta     []float64
		tsa    []float64
		tb     []float64
		tsb    []float64
		nu     float64
		lambda float64
		degree int32
	}
	tests := []struct {
		name         string
		args         args
		wantDistance float64
	}{
		// TODO: Add test cases.
		{"Basic Test", args{
			[]float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1, -0.1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1, -0.1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1.1,
			2.2,
			1,
		}, 2,
		},
		{"Time Series ta too short", args{
			[]float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1, -0.1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1.1,
			2.2,
			1,
		}, -1,
		},
		{"Time Series tb too short", args{
			[]float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1.1,
			2.2,
			1,
		}, -1,
		},
		{"Time Series tsa too short", args{
			[]float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8},
			[]float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1, -0.1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1.1,
			2.2,
			1,
		}, -1,
		},
		{"Time Series tsb too short", args{
			[]float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1.1,
			2.2,
			1,
		}, -1,
		},
		{"Parameter degree not negative", args{
			[]float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1, 1},
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1.1,
			2.2,
			-1,
		}, -1,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := TWED(tt.args.ta, tt.args.tsa, tt.args.tb, tt.args.tsb, tt.args.nu, tt.args.lambda, tt.args.degree); gotDistance != tt.wantDistance {
				t.Errorf("TWED() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
