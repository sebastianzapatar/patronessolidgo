package main

import "fmt"

// 1. Componente (interfaz común para hojas simples y contenedores complejos).
type NodoSistema interface {
	Nombre() string
	TamanoKB() int
	Mostrar(indent string)
}

// 2. Hoja (Leaf): elemento básico del árbol que no tiene subelementos.
type Archivo struct {
	nombre   string
	tamanoKB int
}

func NuevoArchivo(nombre string, tamanoKB int) *Archivo {
	return &Archivo{nombre: nombre, tamanoKB: tamanoKB}
}

func (a *Archivo) Nombre() string   { return a.nombre }
func (a *Archivo) TamanoKB() int    { return a.tamanoKB }
func (a *Archivo) Mostrar(ind string) {
	fmt.Printf("%s📄 %s (%d KB)\n", ind, a.nombre, a.tamanoKB)
}

// 3. Contenedor / Compuesto (Composite): agrupa hojas u otros contenedores.
type Carpeta struct {
	nombre string
	hijos  []NodoSistema
}

func NuevaCarpeta(nombre string) *Carpeta {
	return &Carpeta{nombre: nombre}
}

func (c *Carpeta) Agregar(nodos ...NodoSistema) {
	c.hijos = append(c.hijos, nodos...)
}

func (c *Carpeta) Nombre() string { return c.nombre }

// TamanoKB delega recursivamente el cálculo a cada hijo del subárbol.
func (c *Carpeta) TamanoKB() int {
	total := 0
	for _, h := range c.hijos {
		total += h.TamanoKB()
	}
	return total
}

func (c *Carpeta) Mostrar(ind string) {
	fmt.Printf("%s📁 %s/ [total: %d KB]\n", ind, c.nombre, c.TamanoKB())
	for _, h := range c.hijos {
		h.Mostrar(ind + "   ")
	}
}

// 4. Cliente: trata igual a un archivo suelto que a todo un árbol de carpetas.
func ImprimirResumen(n NodoSistema) {
	fmt.Printf("Resumen de '%s' -> %d KB\n", n.Nombre(), n.TamanoKB())
}

func main() {
	src := NuevaCarpeta("src")
	src.Agregar(
		NuevoArchivo("lexer.go", 18),
		NuevoArchivo("parser.go", 42),
	)

	raiz := NuevaCarpeta("compilador")
	raiz.Agregar(
		NuevoArchivo("go.mod", 2),
		NuevoArchivo("README.md", 8),
		src,
	)

	raiz.Mostrar("")
	fmt.Println("---")
	ImprimirResumen(NuevoArchivo("lexer.go", 18)) // Hoja individual
	ImprimirResumen(raiz)                         // Árbol compuesto completo
}
