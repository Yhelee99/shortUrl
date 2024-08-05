package basex

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	tests := []struct {
		name string
		s    uint64
		want string
	}{
		{
			name: "case:0",
			s:    0,
			want: "0",
		},
		{
			name: "case:6347",
			s:    6347,
			want: "1En",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.s); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want uint64
	}{
		{
			name: "case:1En",
			s:    "1En",
			want: 6347,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.s); got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
