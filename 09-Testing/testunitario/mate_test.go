package testunitario

import "testing"

// func TestSuma(t *testing.T) {
// 	total := Suma(5, 5)

// 	if total != 10 {
// 		t.Errorf("Suma incorrecta, tiene %d se espearaba %d", total, 10)
// 	}
// }

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	} {
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se espearaba %d", total, 10)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{5, 3, 5},
		{2, 3, 3},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("GetMax incorrecta, tiene %d se espearaba %d", max, item.c)
		}
	}
}
