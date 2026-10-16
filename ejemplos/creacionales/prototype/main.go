package main

import "fmt"

// 1. Interfaz Prototipo (declara el método de clonación)
type Prototipo interface {
	Clonar() Prototipo
	Mostrar(indent string)
}

// 2. Prototipo Concreto A: Archivo
type Archivo struct {
	Nombre   string
	Etiqueta []string // slice para demostrar copia profunda (deep copy)
}

func (a *Archivo) Mostrar(indent string) {
	fmt.Printf("%s📄 %s (tags: %v)\n", indent, a.Nombre, a.Etiqueta)
}

func (a *Archivo) Clonar() Prototipo {
	copiaTags := make([]string, len(a.Etiqueta))
	copy(copiaTags, a.Etiqueta)
	return &Archivo{Nombre: a.Nombre + "_copia", Etiqueta: copiaTags}
}

// 2. Prototipo Concreto B: Carpeta (contiene otros Prototipos)
type Carpeta struct {
	Nombre    string
	Elementos []Prototipo
}

func (c *Carpeta) Mostrar(indent string) {
	fmt.Printf("%s📁 %s/\n", indent, c.Nombre)
	for _, elem := range c.Elementos {
		elem.Mostrar(indent + "   ")
	}
}

func (c *Carpeta) Clonar() Prototipo {
	clon := &Carpeta{
		Nombre:    c.Nombre + "_clon",
		Elementos: make([]Prototipo, len(c.Elementos)),
	}
	for i, elem := range c.Elementos {
		clon.Elementos[i] = elem.Clonar() // clona recursivamente sin saber el tipo concreto
	}
	return clon
}

func main() {
	doc := &Archivo{Nombre: "informe.pdf", Etiqueta: []string{"v1", "oficial"}}
	orig := &Carpeta{
		Nombre:    "Proyecto",
		Elementos: []Prototipo{doc, &Archivo{Nombre: "datos.csv", Etiqueta: []string{"raw"}}},
	}

	// 3. El cliente clona la estructura completa a través de la interfaz Prototipo
	copia := orig.Clonar()
	doc.Etiqueta[0] = "MODIFICADO" // no afecta al clon gracias a la copia profunda

	fmt.Println("--- Original (tras modificar doc) ---")
	orig.Mostrar("")
	fmt.Println("--- Clon independiente ---")
	copia.Mostrar("")
}
