package figuras

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) obtenerArea() float64 {
	return c.Lado * c.Lado
}
func (c *Cuadrado) obtenerPerimetro() float64 {
	return 4 * c.Lado
}
