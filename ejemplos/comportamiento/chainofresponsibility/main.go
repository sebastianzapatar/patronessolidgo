package main

import "fmt"

// Solicitud contiene los datos que viajan a lo largo de la cadena.
type Solicitud struct {
	Usuario     string
	Autenticado bool
	Intentos    int
	Datos       string
}

// 1. Manejador (interfaz común a todos los eslabones)
type Manejador interface {
	EstablecerSiguiente(m Manejador) Manejador
	Manejar(s *Solicitud)
}

// 2. ManejadorBase (evita duplicar el enlace al siguiente eslabón)
type ManejadorBase struct {
	siguiente Manejador
}

func (b *ManejadorBase) EstablecerSiguiente(m Manejador) Manejador {
	b.siguiente = m
	return m
}

func (b *ManejadorBase) PasarAlSiguiente(s *Solicitud) {
	if b.siguiente != nil {
		b.siguiente.Manejar(s)
	}
}

// 3a. Manejador Concreto: protege contra fuerza bruta
type FiltroFuerzaBruta struct{ ManejadorBase }

func (f *FiltroFuerzaBruta) Manejar(s *Solicitud) {
	if s.Intentos >= 3 {
		fmt.Printf("  [FuerzaBruta] ✘ IP bloqueada para %q (%d intentos fallidos)\n", s.Usuario, s.Intentos)
		return // Corta la cadena
	}
	fmt.Println("  [FuerzaBruta] ✔ Intentos dentro del límite")
	f.PasarAlSiguiente(s)
}

// 3b. Manejador Concreto: verifica credenciales
type FiltroAutenticacion struct{ ManejadorBase }

func (a *FiltroAutenticacion) Manejar(s *Solicitud) {
	if !s.Autenticado {
		fmt.Printf("  [Autenticación] ✘ Acceso denegado a %q: sesión no iniciada\n", s.Usuario)
		return // Corta la cadena
	}
	fmt.Printf("  [Autenticación] ✔ Usuario %q autenticado\n", s.Usuario)
	a.PasarAlSiguiente(s)
}

// 3c. Manejador Concreto: valida y sanea el payload
type FiltroValidacion struct{ ManejadorBase }

func (v *FiltroValidacion) Manejar(s *Solicitud) {
	if s.Datos == "" {
		fmt.Println("  [Validación] ✘ Solicitud rechazada: datos del pedido vacíos")
		return
	}
	fmt.Println("  [Validación] ✔ Datos del pedido saneados")
	v.PasarAlSiguiente(s)
}

// 3d. Manejador Concreto final: procesa el pedido
type ProcesadorPedidos struct{ ManejadorBase }

func (p *ProcesadorPedidos) Manejar(s *Solicitud) {
	fmt.Printf("  [Pedidos] 🎉 Pedido procesado con éxito para %q: %s\n", s.Usuario, s.Datos)
}

func main() {
	// 4. Cliente: ensambla la cadena de responsabilidad
	fuerzaBruta := &FiltroFuerzaBruta{}
	autenticacion := &FiltroAutenticacion{}
	validacion := &FiltroValidacion{}
	procesador := &ProcesadorPedidos{}

	fuerzaBruta.EstablecerSiguiente(autenticacion).
		EstablecerSiguiente(validacion).
		EstablecerSiguiente(procesador)

	casos := []*Solicitud{
		{Usuario: "atacante", Autenticado: false, Intentos: 5, Datos: "2x Laptop"},
		{Usuario: "invitado", Autenticado: false, Intentos: 1, Datos: "1x Teclado"},
		{Usuario: "ana", Autenticado: true, Intentos: 0, Datos: "1x Monitor 4K"},
	}

	for i, req := range casos {
		fmt.Printf("Solicitud #%d (usuario=%q):\n", i+1, req.Usuario)
		fuerzaBruta.Manejar(req)
	}
}
