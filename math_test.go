package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é invalido: resultado%d. esperado: %d", total, 30)
	}
}
