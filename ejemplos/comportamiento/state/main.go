package main

import "fmt"

// Estado (State): interfaz con las acciones dependientes del estado.
type EstadoDocumento interface {
	Nombre() string
	Revisar(d *Documento)
	Publicar(d *Documento, esAdmin bool)
}

// Contexto (Context): mantiene referencia al estado actual y delega en él.
type Documento struct {
	Titulo string
	estado EstadoDocumento
}

func NuevoDocumento(titulo string) *Documento {
	doc := &Documento{Titulo: titulo}
	doc.CambiarEstado(&EstadoBorrador{})
	return doc
}

func (d *Documento) CambiarEstado(nuevo EstadoDocumento) {
	d.estado = nuevo
	fmt.Printf("  -> [%s] ahora está en estado: %s\n", d.Titulo, nuevo.Nombre())
}

func (d *Documento) Revisar()                 { d.estado.Revisar(d) }
func (d *Documento) Publicar(esAdmin bool)    { d.estado.Publicar(d, esAdmin) }

// Estado Concreto 1: Borrador
type EstadoBorrador struct{}

func (e *EstadoBorrador) Nombre() string { return "Borrador" }
func (e *EstadoBorrador) Revisar(d *Documento) {
	fmt.Println("[Borrador] Enviando documento a moderación...")
	d.CambiarEstado(&EstadoModeracion{})
}
func (e *EstadoBorrador) Publicar(d *Documento, esAdmin bool) {
	fmt.Println("[Borrador] No se puede publicar directamente sin revisión.")
}

// Estado Concreto 2: Moderación
type EstadoModeracion struct{}

func (e *EstadoModeracion) Nombre() string { return "Moderación" }
func (e *EstadoModeracion) Revisar(d *Documento) {
	fmt.Println("[Moderación] El documento ya está en revisión.")
}
func (e *EstadoModeracion) Publicar(d *Documento, esAdmin bool) {
	if !esAdmin {
		fmt.Println("[Moderación] Permiso denegado: solo un admin puede aprobar. Vuelve a borrador.")
		d.CambiarEstado(&EstadoBorrador{})
		return
	}
	fmt.Println("[Moderación] Aprobado por administrador.")
	d.CambiarEstado(&EstadoPublicado{})
}

// Estado Concreto 3: Publicado
type EstadoPublicado struct{}

func (e *EstadoPublicado) Nombre() string { return "Publicado" }
func (e *EstadoPublicado) Revisar(d *Documento) {
	fmt.Println("[Publicado] Ya está publicado; no requiere revisión.")
}
func (e *EstadoPublicado) Publicar(d *Documento, esAdmin bool) {
	fmt.Println("[Publicado] Operación ignorada: el documento ya es público.")
}

func main() {
	doc := NuevoDocumento("Especificación Gramática LL(1)")
	doc.Publicar(false)
	doc.Revisar()
	doc.Publicar(false) // Falla y regresa a Borrador
	doc.Revisar()
	doc.Publicar(true)  // Admin aprueba -> Publicado
	doc.Revisar()
}
