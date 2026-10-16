package main

import "fmt"

// 1. Originador: objeto cuyo estado interno queremos salvaguardar
type Editor struct {
	texto  string
	cursor int
}

func (e *Editor) Escribir(fragmento string) {
	e.texto += fragmento
	e.cursor = len(e.texto)
}

func (e *Editor) Mostrar() string {
	return fmt.Sprintf("texto=%q (cursor=%d)", e.texto, e.cursor)
}

// Crea una instantánea inmutable de su estado actual
func (e *Editor) Guardar(etiqueta string) *Memento {
	return &Memento{
		etiqueta: etiqueta,
		texto:    e.texto,
		cursor:   e.cursor,
	}
}

// Restaura su estado interno a partir de un memento
func (e *Editor) Restaurar(m *Memento) {
	if m == nil {
		return
	}
	e.texto = m.texto
	e.cursor = m.cursor
}

// 2. Memento: almacena el estado del originador en campos privados
type Memento struct {
	etiqueta string
	texto    string
	cursor   int
}

// El cuidador solo puede consultar metadatos públicos, nunca alterar el estado
func (m *Memento) Etiqueta() string { return m.etiqueta }

// 3. Cuidador (Caretaker): administra la pila de instantáneas
type Historial struct {
	instantaneas []*Memento
}

func (h *Historial) Registrar(m *Memento) {
	fmt.Printf("  💾 Guardada instantánea [%s]\n", m.Etiqueta())
	h.instantaneas = append(h.instantaneas, m)
}

func (h *Historial) ExtraerUltima() *Memento {
	if len(h.instantaneas) == 0 {
		return nil
	}
	idx := len(h.instantaneas) - 1
	m := h.instantaneas[idx]
	h.instantaneas = h.instantaneas[:idx]
	return m
}

func main() {
	editor := &Editor{}
	historial := &Historial{}

	editor.Escribir("package main; ")
	historial.Registrar(editor.Guardar("v1: encabezado"))

	editor.Escribir("func main() { ")
	historial.Registrar(editor.Guardar("v2: firma main"))

	editor.Escribir("syntax_error??? }")
	fmt.Println("  Estado actual:", editor.Mostrar())

	fmt.Println("\nRestaurando estados previos desde el historial:")
	snap := historial.ExtraerUltima()
	editor.Restaurar(snap)
	fmt.Printf("  ⮌ Restaurado [%s] → %s\n", snap.Etiqueta(), editor.Mostrar())

	snap = historial.ExtraerUltima()
	editor.Restaurar(snap)
	fmt.Printf("  ⮌ Restaurado [%s] → %s\n", snap.Etiqueta(), editor.Mostrar())
}
