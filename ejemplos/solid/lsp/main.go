package main

import "fmt"

// ❌ VIOLACIÓN: una interfaz con setters que Cuadrado no puede cumplir.
type RectanguloMutable interface {
	SetAncho(w float64)
	SetAlto(h float64)
	Area() float64
}

type Rectangulo struct{ ancho, alto float64 }

func (r *Rectangulo) SetAncho(w float64) { r.ancho = w }
func (r *Rectangulo) SetAlto(h float64)  { r.alto = h }
func (r *Rectangulo) Area() float64      { return r.ancho * r.alto }

type Cuadrado struct{ lado float64 }

func (c *Cuadrado) SetAncho(w float64) { c.lado = w } // cambia ambos lados
func (c *Cuadrado) SetAlto(h float64)  { c.lado = h } // cambia ambos lados
func (c *Cuadrado) Area() float64      { return c.lado * c.lado }

// El cliente confía en el contrato: "ancho y alto son independientes".
func PruebaArea(r RectanguloMutable) {
	r.SetAncho(5)
	r.SetAlto(4)
	fmt.Printf("%-16T área esperada 20, obtenida %.0f\n", r, r.Area())
}

// ✅ SOLUCIÓN: un contrato que TODOS pueden cumplir.
type Forma interface{ Area() float64 }

type Rect struct{ Ancho, Alto float64 }
type Cuad struct{ Lado float64 }

func (r Rect) Area() float64 { return r.Ancho * r.Alto }
func (c Cuad) Area() float64 { return c.Lado * c.Lado }

func AreaTotal(formas ...Forma) (t float64) {
	for _, f := range formas {
		t += f.Area()
	}
	return
}

func main() {
	PruebaArea(&Rectangulo{})
	PruebaArea(&Cuadrado{}) // 💥 rompe la expectativa del cliente

	fmt.Println("Área total:", AreaTotal(Rect{5, 4}, Cuad{3}))
}
