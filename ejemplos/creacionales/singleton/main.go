package main

import (
	"fmt"
	"sync"
)

// 1. Estructura Singleton (no exportada o con inicialización controlada)
type ConexionBD struct {
	DSN       string
	Consultas int
}

var (
	instancia *ConexionBD
	once      sync.Once // garantiza ejecución única y segura entre goroutines
)

// ObtenerBD es el punto de acceso global al Singleton.
func ObtenerBD() *ConexionBD {
	once.Do(func() {
		fmt.Println("🔌 [sync.Once] Creando la ÚNICA conexión a la base de datos...")
		instancia = &ConexionBD{DSN: "postgres://localhost:5432/produccion"}
	})
	return instancia
}

func (db *ConexionBD) Ejecutar(sql string) {
	db.Consultas++
	fmt.Printf("   [Consulta #%d en %s] -> %s\n", db.Consultas, db.DSN, sql)
}

func main() {
	// Primera llamada: inicializa la instancia única
	bd1 := ObtenerBD()
	bd1.Ejecutar("SELECT * FROM usuarios")

	// Segunda llamada: reutiliza exactamente la misma instancia
	bd2 := ObtenerBD()
	bd2.Ejecutar("INSERT INTO logs VALUES ('login')")

	fmt.Printf("¿bd1 y bd2 apuntan al mismo objeto? %v (total consultas: %d)\n",
		bd1 == bd2, bd1.Consultas)
}
