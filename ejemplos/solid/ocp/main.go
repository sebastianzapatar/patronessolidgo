package main

import "fmt"

type Factura struct {
	ID    int
	Total float64
}

// Abstracción estable: CERRADA a modificación.
type FacturaPersistencia interface {
	Guardar(f Factura) error
}

// Extensiones: cada nueva forma de guardar es un tipo NUEVO.
type ArchivoPersistencia struct{ Ruta string }

func (a ArchivoPersistencia) Guardar(f Factura) error {
	fmt.Printf("[archivo] factura #%d -> %s\n", f.ID, a.Ruta)
	return nil
}

type BaseDeDatosPersistencia struct{ Tabla string }

func (b BaseDeDatosPersistencia) Guardar(f Factura) error {
	fmt.Printf("[postgres] INSERT INTO %s VALUES (%d, %.2f)\n", b.Tabla, f.ID, f.Total)
	return nil
}

// Agregado meses después SIN tocar nada de lo anterior.
type MongoPersistencia struct{ Coleccion string }

func (m MongoPersistencia) Guardar(f Factura) error {
	fmt.Printf("[mongo] %s.insertOne({id: %d})\n", m.Coleccion, f.ID)
	return nil
}

// Este código nunca cambia aunque aparezcan nuevos almacenamientos.
func GuardarTodas(p FacturaPersistencia, facturas []Factura) {
	for _, f := range facturas {
		if err := p.Guardar(f); err != nil {
			fmt.Println("error:", err)
		}
	}
}

func main() {
	facturas := []Factura{{1, 107.1}, {2, 59.5}}
	destinos := []FacturaPersistencia{
		ArchivoPersistencia{"facturas.txt"},
		BaseDeDatosPersistencia{"facturas"},
		MongoPersistencia{"facturas"},
	}
	for _, d := range destinos {
		GuardarTodas(d, facturas)
	}
}
