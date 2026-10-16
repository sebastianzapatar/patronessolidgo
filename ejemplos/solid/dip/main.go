package main

import "fmt"

type Factura struct{ ID int }
type Libro struct{ ISBN string }

// Abstracciones (las define quien las USA, no quien las implementa)
type FacturaPersistencia interface{ Guardar(Factura) }
type LibroPersistencia interface{ Guardar(Libro) }

// Módulo de ALTO nivel: depende solo de interfaces.
type AdministradorPersistencia struct {
	facturas FacturaPersistencia
	libros   LibroPersistencia
}

// Inyección de dependencias por constructor
func NewAdministrador(f FacturaPersistencia, l LibroPersistencia) *AdministradorPersistencia {
	return &AdministradorPersistencia{facturas: f, libros: l}
}

func (a *AdministradorPersistencia) Respaldar(fs []Factura, ls []Libro) {
	for _, f := range fs {
		a.facturas.Guardar(f)
	}
	for _, l := range ls {
		a.libros.Guardar(l)
	}
}

// Módulos de BAJO nivel: detalles intercambiables.
type MySQLFacturas struct{}

func (MySQLFacturas) Guardar(f Factura) { fmt.Println("MySQL   <- factura", f.ID) }

type MemoriaLibros struct{ datos []Libro }

func (m *MemoriaLibros) Guardar(l Libro) {
	m.datos = append(m.datos, l)
	fmt.Println("Memoria <- libro", l.ISBN)
}

// Un doble de prueba: cambiar el detalle no toca el alto nivel.
type FakeFacturas struct{ n int }

func (f *FakeFacturas) Guardar(Factura) { f.n++ }

func main() {
	admin := NewAdministrador(MySQLFacturas{}, &MemoriaLibros{})
	admin.Respaldar([]Factura{{1}, {2}}, []Libro{{"978-84"}})

	fake := &FakeFacturas{}
	NewAdministrador(fake, &MemoriaLibros{}).Respaldar([]Factura{{1}, {2}, {3}}, nil)
	fmt.Println("En pruebas se guardaron", fake.n, "facturas")
}
