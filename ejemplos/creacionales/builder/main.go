package main

import "fmt"

// 3. Producto complejo resultante
type Casa struct {
	Paredes  string
	Techo    string
	Ventanas int
	ConPiscina bool
}

func (c Casa) String() string {
	return fmt.Sprintf("🏠 Casa [Paredes: %s | Techo: %s | Ventanas: %d | Piscina: %v]",
		c.Paredes, c.Techo, c.Ventanas, c.ConPiscina)
}

// 1. Interfaz Constructora (define los pasos comunes)
type ConstructorCasa interface {
	Reiniciar()
	ConstruirParedes()
	ConstruirTecho()
	ConstruirVentanas()
	ConstruirPiscina()
	ObtenerCasa() Casa
}

// 2. Constructor Concreto A: Cabaña de madera
type ConstructorCabana struct{ casa Casa }

func (b *ConstructorCabana) Reiniciar()         { b.casa = Casa{} }
func (b *ConstructorCabana) ConstruirParedes()  { b.casa.Paredes = "Troncos de pino" }
func (b *ConstructorCabana) ConstruirTecho()    { b.casa.Techo = "Tejas de madera" }
func (b *ConstructorCabana) ConstruirVentanas() { b.casa.Ventanas = 2 }
func (b *ConstructorCabana) ConstruirPiscina()  { b.casa.ConPiscina = false }
func (b *ConstructorCabana) ObtenerCasa() Casa  { return b.casa }

// 2. Constructor Concreto B: Mansión moderna
type ConstructorMansion struct{ casa Casa }

func (b *ConstructorMansion) Reiniciar()         { b.casa = Casa{} }
func (b *ConstructorMansion) ConstruirParedes()  { b.casa.Paredes = "Hormigón y cristal" }
func (b *ConstructorMansion) ConstruirTecho()    { b.casa.Techo = "Terraza solar" }
func (b *ConstructorMansion) ConstruirVentanas() { b.casa.Ventanas = 12 }
func (b *ConstructorMansion) ConstruirPiscina()  { b.casa.ConPiscina = true }
func (b *ConstructorMansion) ObtenerCasa() Casa  { return b.casa }

// 4. Director (define el orden de los pasos de construcción)
type Director struct{ constructor ConstructorCasa }

func (d *Director) CambiarConstructor(c ConstructorCasa) { d.constructor = c }

func (d *Director) ConstruirCompleta() Casa {
	d.constructor.Reiniciar()
	d.constructor.ConstruirParedes()
	d.constructor.ConstruirTecho()
	d.constructor.ConstruirVentanas()
	d.constructor.ConstruirPiscina()
	return d.constructor.ObtenerCasa()
}

func main() {
	director := &Director{}

	cabana := &ConstructorCabana{}
	director.CambiarConstructor(cabana)
	fmt.Println(director.ConstruirCompleta())

	mansion := &ConstructorMansion{}
	director.CambiarConstructor(mansion)
	fmt.Println(director.ConstruirCompleta())
}
