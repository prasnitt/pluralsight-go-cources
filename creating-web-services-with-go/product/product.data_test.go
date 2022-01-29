package product

import "testing"

func Test_nextId(t *testing.T) {
	tests := []struct {
		name     string
		products map[int]Product
		want     int
	}{
		{"empty product map", map[int]Product{}, 1},
		{"one product in map", map[int]Product{0: Product{}}, 2},
	}
	for _, tt := range tests {
		productMap.m = tt.products
		t.Run(tt.name, func(t *testing.T) {
			if got := nextId(); got != tt.want {
				t.Errorf("nextId() = %v, want %v", got, tt.want)
			}
		})
	}
}
