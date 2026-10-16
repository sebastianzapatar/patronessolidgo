package main

import "fmt"

// 1. Producto (interfaz común a todos los transportes)
type Transporte interface {
	Entregar(carga string) string
}

// 2. Productos Concretos
type Camion struct{ Placa string }

func (c *Camion) Entregar(carga string) string {
	return fmt.Sprintf("🚚 Camión [%s] entrega %q por carretera", c.Placa, carga)
}

type Barco struct{ Nombre string }

func (b *Barco) Entregar(carga string) string {
	return fmt.Sprintf("🚢 Barco %q entrega %q por vía marítima", b.Nombre, carga)
}

// 3. Creador (interfaz con el Factory Method)
type Logistica interface {
	CrearTransporte() Transporte // Factory Method
}

// Lógica de negocio común que trabaja con la interfaz Creador
// sin conocer el tipo concreto de transporte que se creará.
func PlanificarEntrega(l Logistica, carga string) {
	transporte := l.CrearTransporte()
	fmt.Println("📦 Planificando:", transporte.Entregar(carga))
}

// 4. Creadores Concretos (sobrescriben el Factory Method)
type LogisticaTerrestre struct{ PlacaAsignada string }

func (lt *LogisticaTerrestre) CrearTransporte() Transporte {
	return &Camion{Placa: lt.PlacaAsignada}
}

type LogisticaMaritima struct{ BuqueAsignado string }

func (lm *LogisticaMaritima) CrearTransporte() Transporte {
	return &Barco{Nombre: lm.BuqueAsignado}
}

func main() {
	terrestre := &LogisticaTerrestre{PlacaAsignada: "ABC-123"}
	maritima := &LogisticaMaritima{BuqueAsignado: "El Navegante"}

	PlanificarEntrega(terrestre, "Electrónicos")
	PlanificarEntrega(maritima, "Contenedores de grano")
}
