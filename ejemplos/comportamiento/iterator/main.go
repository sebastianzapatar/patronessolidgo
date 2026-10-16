package main

import (
	"fmt"
	"iter"
)

type Libro struct {
	Titulo string
	Tema   string
}

// 1. Iterador (interfaz común para recorrer elementos)
type Iterador interface {
	HayMas() bool
	Siguiente() Libro
}

// 3. Colección (interfaz que fabrica iteradores)
type ColeccionLibros interface {
	CrearIteradorPorTema(tema string) Iterador
}

// 4. Colección Concreta
type Biblioteca struct {
	libros []Libro
}

func (b *Biblioteca) Agregar(l Libro) { b.libros = append(b.libros, l) }

func (b *Biblioteca) CrearIteradorPorTema(tema string) Iterador {
	return &IteradorPorTema{biblio: b, tema: tema, posicion: 0}
}

// Método idiomático en Go 1.23+: iterador funcional compatible con for-range
func (b *Biblioteca) TodosEnOrdenInverso() iter.Seq[Libro] {
	return func(yield func(Libro) bool) {
		for i := len(b.libros) - 1; i >= 0; i-- {
			if !yield(b.libros[i]) {
				return
			}
		}
	}
}

// 2. Iterador Concreto: guarda su propio cursor de recorrido
type IteradorPorTema struct {
	biblio   *Biblioteca
	tema     string
	posicion int
}

func (it *IteradorPorTema) HayMas() bool {
	for it.posicion < len(it.biblio.libros) {
		if it.biblio.libros[it.posicion].Tema == it.tema {
			return true
		}
		it.posicion++
	}
	return false
}

func (it *IteradorPorTema) Siguiente() Libro {
	libro := it.biblio.libros[it.posicion]
	it.posicion++
	return libro
}

func main() {
	// 5. Cliente: usa la interfaz Iterador sin conocer el slice interno
	biblio := &Biblioteca{}
	biblio.Agregar(Libro{Titulo: "Diseño de Compiladores", Tema: "Sistemas"})
	biblio.Agregar(Libro{Titulo: "Patrones de Diseño GoF", Tema: "Arquitectura"})
	biblio.Agregar(Libro{Titulo: "El Lenguaje Go", Tema: "Sistemas"})
	biblio.Agregar(Libro{Titulo: "Arquitectura Limpia", Tema: "Arquitectura"})

	fmt.Println("Recorrido con Iterador clásico (Tema = Sistemas):")
	it := biblio.CrearIteradorPorTema("Sistemas")
	for it.HayMas() {
		l := it.Siguiente()
		fmt.Printf("  📖 %s [%s]\n", l.Titulo, l.Tema)
	}

	fmt.Println("\nRecorrido con iter.Seq de Go 1.23+ (orden inverso):")
	for l := range biblio.TodosEnOrdenInverso() {
		fmt.Printf("  🔄 %s [%s]\n", l.Titulo, l.Tema)
	}
}
