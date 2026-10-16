package main

import "fmt"

// 2. Comando (interfaz con método de ejecución y reversión)
type Comando interface {
	Ejecutar() bool
	Deshacer()
	Nombre() string
}

// 4. Receptor: contiene la lógica de negocio real
type EditorTexto struct {
	Texto string
}

func (e *EditorTexto) Agregar(fragmento string) { e.Texto += fragmento }
func (e *EditorTexto) BorrarUltimos(n int) {
	if n > len(e.Texto) {
		n = len(e.Texto)
	}
	e.Texto = e.Texto[:len(e.Texto)-n]
}

// 3a. Comando Concreto: insertar texto
type ComandoEscribir struct {
	editorRespaldo *EditorTexto
	nuevoTexto     string
	estadoPrevio   string
}

func (c *ComandoEscribir) Ejecutar() bool {
	c.estadoPrevio = c.editorRespaldo.Texto
	c.editorRespaldo.Agregar(c.nuevoTexto)
	return true // modifica el estado -> guardar en historial
}

func (c *ComandoEscribir) Deshacer()      { c.editorRespaldo.Texto = c.estadoPrevio }
func (c *ComandoEscribir) Nombre() string { return "Escribir(" + c.nuevoTexto + ")" }

// 3b. Comando Concreto: borrar los últimos N caracteres
type ComandoBorrar struct {
	editor       *EditorTexto
	cantidad     int
	estadoPrevio string
}

func (c *ComandoBorrar) Ejecutar() bool {
	c.estadoPrevio = c.editor.Texto
	c.editor.BorrarUltimos(c.cantidad)
	return true
}

func (c *ComandoBorrar) Deshacer()      { c.editor.Texto = c.estadoPrevio }
func (c *ComandoBorrar) Nombre() string { return fmt.Sprintf("Borrar(%d)", c.cantidad) }

// 1. Emisor / Invocador: dispara comandos y administra la pila de historial
type Aplicacion struct {
	historial []Comando
}

func (a *Aplicacion) EjecutarComando(c Comando) {
	if c.Ejecutar() {
		a.historial = append(a.historial, c)
	}
	fmt.Printf("  ▶ Ejecutado %-20s →", c.Nombre())
}

func (a *Aplicacion) DeshacerUltimo() {
	if len(a.historial) == 0 {
		return
	}
	ultimo := a.historial[len(a.historial)-1]
	a.historial = a.historial[:len(a.historial)-1]
	ultimo.Deshacer()
	fmt.Printf("  ⮌ Deshecho  %-20s →", ultimo.Nombre())
}

func main() {
	// 5. Cliente: crea el receptor, los comandos y el invocador
	editor := &EditorTexto{}
	app := &Aplicacion{}

	pasos := []Comando{
		&ComandoEscribir{editorRespaldo: editor, nuevoTexto: "Hola "},
		&ComandoEscribir{editorRespaldo: editor, nuevoTexto: "Mundo"},
		&ComandoBorrar{editor: editor, cantidad: 5},
		&ComandoEscribir{editorRespaldo: editor, nuevoTexto: "Patrones en Go!"},
	}

	for _, cmd := range pasos {
		app.EjecutarComando(cmd)
		fmt.Printf(" %q\n", editor.Texto)
	}

	fmt.Println("\nRevirtiendo las últimas 2 operaciones (Ctrl+Z):")
	app.DeshacerUltimo()
	fmt.Printf(" %q\n", editor.Texto)
	app.DeshacerUltimo()
	fmt.Printf(" %q\n", editor.Texto)
}
