package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(4, 2)
	if err != nil {
		t.Errorf("se esperaba nil, pero se obtuvo error  : %v ", err)
	}

	if result != 2 {
		t.Errorf("se esperaba 2, pero se obtuvo %f ", result)
	}
}

func TestDivide2(t *testing.T) {
	test := []struct {
		a, b   float64
		want   float64
		hasErr bool
	}{
		{4, 2, 2, false},
		{10, 5, 2, false},
		{10, 0, 0, true},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("%f / %f", tt.a, tt.b), func(t *testing.T) {
			t.Parallel() //funcion paralelistmo
			result, err := divide(tt.a, tt.b)
			if (err != nil) != tt.hasErr {
				t.Errorf("error esperado : %v, error obtenido : %v", tt.hasErr, err)
			}
			if result != tt.want {
				t.Errorf("se esperaba %f , pero se obtuvo %f", tt.want, result)
			}
		})
	}
}

/**
 * Funcion benchmark
 */
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = divide(4, 2)
	}
}
