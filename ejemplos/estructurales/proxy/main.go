package main

import "fmt"

// 1. Interfaz de Servicio: implementada tanto por el Servicio Real como por el Proxy.
type RepositorioDocumentos interface {
	ObtenerDocumento(id string, rolUsuario string) string
}

// 2. Servicio Real: operación costosa (consulta a BD o descarga remota).
type AlmacenRemotoS3 struct{}

func (a *AlmacenRemotoS3) ObtenerDocumento(id string, _ string) string {
	fmt.Printf("  🌐 [AlmacenRemotoS3] Descargando '%s' desde el servidor remoto...\n", id)
	return fmt.Sprintf("Contenido confidencial del archivo [%s]", id)
}

// 3. Proxy: controla acceso (protección), guarda en caché e inicializa en diferido (virtual).
type ProxyRepositorio struct {
	servicioReal *AlmacenRemotoS3 // Se crea solo cuando realmente hace falta (lazy init)
	cache        map[string]string
}

func NuevoProxyRepositorio() *ProxyRepositorio {
	return &ProxyRepositorio{cache: make(map[string]string)}
}

func (p *ProxyRepositorio) ObtenerDocumento(id string, rolUsuario string) string {
	fmt.Printf("-> Solicitud de '%s' por rol '%s'\n", id, rolUsuario)

	// a) Proxy de protección: control de acceso
	if rolUsuario != "admin" && rolUsuario != "editor" {
		return "⛔ ACCESO DENEGADO: permisos insuficientes"
	}

	// b) Proxy de caché: evita descargas repetidas
	if contenido, ok := p.cache[id]; ok {
		fmt.Printf("  ⚡ [Proxy Cache] Devolviendo '%s' desde memoria caché\n", id)
		return contenido
	}

	// c) Proxy virtual: inicialización diferida del servicio pesado
	if p.servicioReal == nil {
		fmt.Println("  🔧 [Proxy Virtual] Inicializando conexión con AlmacenRemotoS3...")
		p.servicioReal = &AlmacenRemotoS3{}
	}

	resultado := p.servicioReal.ObtenerDocumento(id, rolUsuario)
	p.cache[id] = resultado
	return resultado
}

func main() {
	var repo RepositorioDocumentos = NuevoProxyRepositorio()

	fmt.Println("  Resultado:", repo.ObtenerDocumento("nomina_2026.pdf", "invitado"))
	fmt.Println("  Resultado:", repo.ObtenerDocumento("nomina_2026.pdf", "admin"))
	fmt.Println("  Resultado:", repo.ObtenerDocumento("nomina_2026.pdf", "editor"))
}
