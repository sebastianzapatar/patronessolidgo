package main

import "fmt"

// 1. Productos Abstractos (familia de objetos relacionados)
type Boton interface {
	Dibujar() string
}

type Casilla interface {
	Marcar(activo bool) string
}

// 2. Productos Concretos — Variante Windows
type BotonWindows struct{}

func (BotonWindows) Dibujar() string { return "[ 🪟 Botón estilo Windows ]" }

type CasillaWindows struct{}

func (CasillaWindows) Marcar(ok bool) string {
	return fmt.Sprintf("[🪟 Casilla Windows: %v]", ok)
}

// 2. Productos Concretos — Variante Mac
type BotonMac struct{}

func (BotonMac) Dibujar() string { return "( 🍎 Botón estilo macOS )" }

type CasillaMac struct{}

func (CasillaMac) Marcar(ok bool) string {
	return fmt.Sprintf("(🍎 Casilla macOS: %v)", ok)
}

// 3. Fábrica Abstracta (crea una familia completa de productos)
type FabricaGUI interface {
	CrearBoton() Boton
	CrearCasilla() Casilla
}

// 4. Fábricas Concretas (garantizan compatibilidad de la familia)
type FabricaWindows struct{}

func (FabricaWindows) CrearBoton() Boton     { return BotonWindows{} }
func (FabricaWindows) CrearCasilla() Casilla { return CasillaWindows{} }

type FabricaMac struct{}

func (FabricaMac) CrearBoton() Boton     { return BotonMac{} }
func (FabricaMac) CrearCasilla() Casilla { return CasillaMac{} }

// 5. Cliente: usa únicamente las interfaces abstractas
func RenderizarPantalla(f FabricaGUI) {
	boton := f.CrearBoton()
	casilla := f.CrearCasilla()
	fmt.Println(boton.Dibujar(), "->", casilla.Marcar(true))
}

func main() {
	fmt.Println("=== Entorno Windows ===")
	RenderizarPantalla(FabricaWindows{})

	fmt.Println("=== Entorno macOS ===")
	RenderizarPantalla(FabricaMac{})
}
