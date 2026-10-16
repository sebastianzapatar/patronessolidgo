package main

import "fmt"

type Coche struct{ Placa string }

// ✅ Interfaces pequeñas y enfocadas (en vez de una "Estacionamiento" gigante).
type Estacionamiento interface {
	Aparcar(c Coche)
	Sacar(c Coche)
	Capacidad() int
}

type Pagable interface {
	CalcularTarifa(c Coche, horas int) float64
}

// base reutilizable por composición (embedding)
type base struct {
	cupos   int
	ocupado map[string]bool
}

func (b *base) Aparcar(c Coche) { b.ocupado[c.Placa] = true; fmt.Println("  entra", c.Placa) }
func (b *base) Sacar(c Coche)   { delete(b.ocupado, c.Placa); fmt.Println("  sale ", c.Placa) }
func (b *base) Capacidad() int  { return b.cupos - len(b.ocupado) }

// Gratis: implementa SOLO lo que necesita. Nada de métodos de pago vacíos.
type EstacionamientoGratis struct{ *base }

// Pagado: implementa Estacionamiento + Pagable.
type EstacionamientoPagado struct {
	*base
	precioHora float64
}

func (e EstacionamientoPagado) CalcularTarifa(_ Coche, horas int) float64 {
	return float64(horas) * e.precioHora
}

func nuevaBase(n int) *base { return &base{cupos: n, ocupado: map[string]bool{}} }

func main() {
	auto := Coche{"ABC-123"}
	lugares := []Estacionamiento{
		EstacionamientoGratis{nuevaBase(10)},
		EstacionamientoPagado{nuevaBase(5), 3000},
	}
	for _, e := range lugares {
		fmt.Printf("%T\n", e)
		e.Aparcar(auto)
		// Solo preguntamos por el pago a quien lo soporta
		if p, ok := e.(Pagable); ok {
			fmt.Printf("  tarifa 2h: $%.0f\n", p.CalcularTarifa(auto, 2))
		}
		e.Sacar(auto)
		fmt.Println("  cupos libres:", e.Capacidad())
	}
}
