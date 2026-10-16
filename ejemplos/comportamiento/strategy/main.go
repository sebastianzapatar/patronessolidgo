package main

import (
	"fmt"
	"strings"
)

// Estrategia (Strategy): interfaz común para toda la familia de algoritmos.
type EstrategiaRuta interface {
	Nombre() string
	CalcularRuta(origen, destino string) ([]string, int)
}

// Estrategia Concreta 1: Automóvil (prioriza autopistas y velocidad).
type RutaAutomovil struct{}

func (r RutaAutomovil) Nombre() string { return "Automóvil" }
func (r RutaAutomovil) CalcularRuta(origen, destino string) ([]string, int) {
	return []string{origen, "Autopista Central", "Peaje Norte", destino}, 25
}

// Estrategia Concreta 2: Transporte Público (combina metro y bus).
type RutaTransportePublico struct{}

func (r RutaTransportePublico) Nombre() string { return "Transporte Público" }
func (r RutaTransportePublico) CalcularRuta(origen, destino string) ([]string, int) {
	return []string{origen, "Estación Metro A", "Transbordo Línea B", destino}, 40
}

// Estrategia Concreta 3: Bicicleta (prioriza ciclovías y parques).
type RutaBicicleta struct{}

func (r RutaBicicleta) Nombre() string { return "Bicicleta" }
func (r RutaBicicleta) CalcularRuta(origen, destino string) ([]string, int) {
	return []string{origen, "Ciclovía del Río", "Parque Lineal", destino}, 55
}

// Contexto (Context): delega el cálculo en la estrategia activa sin usar switch.
type Navegador struct {
	estrategia EstrategiaRuta
}

func NuevoNavegador(e EstrategiaRuta) *Navegador {
	return &Navegador{estrategia: e}
}

func (n *Navegador) CambiarEstrategia(e EstrategiaRuta) {
	n.estrategia = e
}

func (n *Navegador) MostrarRuta(origen, destino string) {
	tramos, minutos := n.estrategia.CalcularRuta(origen, destino)
	fmt.Printf("[%s] (%d min): %s\n",
		n.estrategia.Nombre(), minutos, strings.Join(tramos, " -> "))
}

func main() {
	nav := NuevoNavegador(RutaAutomovil{})
	nav.MostrarRuta("Campus", "Aeropuerto")

	nav.CambiarEstrategia(RutaTransportePublico{})
	nav.MostrarRuta("Campus", "Aeropuerto")

	nav.CambiarEstrategia(RutaBicicleta{})
	nav.MostrarRuta("Campus", "Aeropuerto")
}
