package main

import "fmt"

// Suscriptor (Observer): interfaz común para todos los observadores.
type Suscriptor interface {
	ID() string
	Actualizar(evento, archivo string)
}

// Notificador (Subject / Publisher): gestiona suscripciones y emite eventos.
type GestorEventos struct {
	oyentes map[string][]Suscriptor
}

func NuevoGestorEventos() *GestorEventos {
	return &GestorEventos{oyentes: make(map[string][]Suscriptor)}
}

func (g *GestorEventos) Suscribir(evento string, s Suscriptor) {
	g.oyentes[evento] = append(g.oyentes[evento], s)
}

func (g *GestorEventos) Desuscribir(evento string, id string) {
	lista := g.oyentes[evento]
	for i, s := range lista {
		if s.ID() == id {
			g.oyentes[evento] = append(lista[:i], lista[i+1:]...)
			return
		}
	}
}

func (g *GestorEventos) Notificar(evento, archivo string) {
	for _, s := range g.oyentes[evento] {
		s.Actualizar(evento, archivo)
	}
}

// Notificador concreto: Editor de código que delega la suscripción por composición.
type EditorCodigo struct {
	Eventos *GestorEventos
	archivo string
}

func NuevoEditor() *EditorCodigo {
	return &EditorCodigo{Eventos: NuevoGestorEventos()}
}

func (e *EditorCodigo) Abrir(ruta string) {
	e.archivo = ruta
	e.Eventos.Notificar("abrir", e.archivo)
}

func (e *EditorCodigo) Guardar() {
	e.Eventos.Notificar("guardar", e.archivo)
}

// Suscriptor Concreto 1: registra eventos en bitácora.
type LoggerListener struct{ nombre string }

func (l *LoggerListener) ID() string { return l.nombre }
func (l *LoggerListener) Actualizar(ev, arch string) {
	fmt.Printf("[LOG:%s] evento=%q sobre %s\n", l.nombre, ev, arch)
}

// Suscriptor Concreto 2: ejecuta análisis estático al guardar.
type LinterListener struct{ regla string }

func (l *LinterListener) ID() string { return "linter-" + l.regla }
func (l *LinterListener) Actualizar(ev, arch string) {
	fmt.Printf("[LINTER:%s] verificando sintaxis de %s tras %q\n", l.regla, arch, ev)
}

func main() {
	editor := NuevoEditor()
	logger := &LoggerListener{nombre: "auditoria"}
	linter := &LinterListener{regla: "gofmt"}

	editor.Eventos.Suscribir("abrir", logger)
	editor.Eventos.Suscribir("guardar", logger)
	editor.Eventos.Suscribir("guardar", linter)

	editor.Abrir("parser.go")
	editor.Guardar()

	fmt.Println("--- Dando de baja al linter ---")
	editor.Eventos.Desuscribir("guardar", linter.ID())
	editor.Guardar()
}
