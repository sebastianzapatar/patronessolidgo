package main

import (
	"fmt"
	"strings"
)

// PasosMineria define los pasos primitivos y ganchos (hooks) del algoritmo.
// En Go no hay herencia ni métodos abstractos: modelamos los pasos variables con una interfaz.
type PasosMineria interface {
	Abrir(ruta string) string
	ExtraerDatos(contenido string) []string
	Parsear(crudos []string) map[string]int
	Cerrar(ruta string)
	AntesDeAnalizar() // Gancho (hook) opcional
}

// BaseMineria aporta implementaciones por defecto para pasos opcionales y hooks
// mediante composición (struct embedding).
type BaseMineria struct{}

func (BaseMineria) AntesDeAnalizar() {}
func (BaseMineria) Cerrar(ruta string) {
	fmt.Printf("  [Base] Cerrando descriptor de %s\n", ruta)
}

// MineroDatos contiene el "Método Plantilla" (Ejecutar) que fija el esqueleto del algoritmo.
type MineroDatos struct{}

func (m MineroDatos) Ejecutar(p PasosMineria, ruta string) {
	fmt.Printf("=== Procesando %s ===\n", ruta)
	raw := p.Abrir(ruta)
	crudos := p.ExtraerDatos(raw)
	datos := p.Parsear(crudos)
	p.AntesDeAnalizar() // Hook
	m.analizarYReportar(datos)
	p.Cerrar(ruta)
}

// Paso común fijo en el esqueleto (no sobrescribible).
func (MineroDatos) analizarYReportar(datos map[string]int) {
	fmt.Printf("  [Análisis común] Total claves=%d -> %v\n", len(datos), datos)
}

// Implementación Concreta 1: CSV (embebe BaseMineria y usa el hook por defecto).
type MineroCSV struct{ BaseMineria }

func (MineroCSV) Abrir(ruta string) string {
	fmt.Println("  [CSV] Leyendo texto separado por comas")
	return " compilador,3 ; lexer,2 ; parser,5 "
}
func (MineroCSV) ExtraerDatos(c string) []string { return strings.Split(c, ";") }
func (MineroCSV) Parsear(crudos []string) map[string]int {
	res := make(map[string]int)
	for _, fila := range crudos {
		partes := strings.Split(strings.TrimSpace(fila), ",")
		res[partes[0]] = len(partes[0])
	}
	return res
}

// Implementación Concreta 2: PDF (sobrescribe el hook AntesDeAnalizar).
type MineroPDF struct{ BaseMineria }

func (MineroPDF) Abrir(ruta string) string {
	fmt.Println("  [PDF] Descomprimiendo flujos de objetos PDF")
	return "AST|SSA|IR"
}
func (MineroPDF) ExtraerDatos(c string) []string { return strings.Split(c, "|") }
func (MineroPDF) Parsear(crudos []string) map[string]int {
	res := make(map[string]int)
	for i, tok := range crudos {
		res[tok] = (i + 1) * 10
	}
	return res
}
func (MineroPDF) AntesDeAnalizar() {
	fmt.Println("  [Hook PDF] Limpiando metadatos OCR antes del análisis")
}

func main() {
	plantilla := MineroDatos{}
	plantilla.Ejecutar(MineroCSV{}, "metricas.csv")
	plantilla.Ejecutar(MineroPDF{}, "articulo.pdf")
}
