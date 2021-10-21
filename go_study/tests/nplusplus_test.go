package tests

import "testing"

func Test_nPlusPlus(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "n++",
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPlusPlus(); got != tt.want {
				t.Errorf("nPlusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nPlusPlusRight(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "n++ right",
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPlusPlusRight(); got != tt.want {
				t.Errorf("nPlusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nPlusPlusRight2(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "n++ right",
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPlusPlusRight2(); got != tt.want {
				t.Errorf("nPlusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}
