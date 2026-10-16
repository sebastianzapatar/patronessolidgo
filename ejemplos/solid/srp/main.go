package main

import "fmt"

type Libro struct {
	Nombre, Autor string
	Precio        float64
}

// Factura: SOLO sabe calcular. Su única razón para cambiar
// es un cambio en las reglas de negocio de la facturación.
type Factura struct {
	Libro         Libro
	Cantidad      int
	TasaDescuento float64
	TasaImpuesto  float64
}

func (f Factura) Total() float64 {
	precio := (f.Libro.Precio - f.Libro.Precio*f.TasaDescuento) * float64(f.Cantidad)
	return precio * (1 + f.TasaImpuesto)
}

// FacturaImpresion: única responsabilidad -> presentar la factura.
type FacturaImpresion struct{ factura Factura }

func (fi FacturaImpresion) Imprimir() {
	f := fi.factura
	fmt.Printf("%dx %s %.2f$\n", f.Cantidad, f.Libro.Nombre, f.Libro.Precio)
	fmt.Printf("Descuento: %.0f%%  Impuesto: %.0f%%\n", f.TasaDescuento*100, f.TasaImpuesto*100)
	fmt.Printf("Total: %.2f$\n", f.Total())
}

// FacturaPersistencia: única responsabilidad -> guardar la factura.
type FacturaPersistencia struct{ factura Factura }

func (fp FacturaPersistencia) GuardarArchivo(nombre string) {
	fmt.Printf("Guardando factura de %q en %s\n", fp.factura.Libro.Nombre, nombre)
}

func main() {
	libro := Libro{Nombre: "El Quijote", Autor: "Cervantes", Precio: 50}
	f := Factura{Libro: libro, Cantidad: 2, TasaDescuento: 0.10, TasaImpuesto: 0.19}

	FacturaImpresion{f}.Imprimir()
	FacturaPersistencia{f}.GuardarArchivo("factura.txt")
}
