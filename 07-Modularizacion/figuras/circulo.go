package figuras

type Circulo struct {
	Radio float64
}

func (c *Circulo) obtenerArea() float64 {
	return 3.1415 * c.Radio * c.Radio
}
func (c *Circulo) obtenerPerimetro() float64 {
	return 2 * 3.1415 * c.Radio
}
