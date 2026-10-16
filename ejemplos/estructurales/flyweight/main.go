package main

import "fmt"

// 1. Flyweight (Estado Intrínseco): inmutable y compartido entre miles de instancias.
type TipoArbol struct {
	especie string
	color   string
	textura string // Simula un recurso pesado en RAM (malla 3D / textura HD)
}

func (t *TipoArbol) Dibujar(x, y int) {
	fmt.Printf("  🌳 [%s | %s | %s] en (%d, %d) (ptr=%p)\n",
		t.especie, t.color, t.textura, x, y, t)
}

// 2. Fábrica Flyweight: gestiona el pool/caché de objetos compartidos.
type FabricaTiposArbol struct {
	tipos map[string]*TipoArbol
}

func NuevaFabricaTipos() *FabricaTiposArbol {
	return &FabricaTiposArbol{tipos: make(map[string]*TipoArbol)}
}

func (f *FabricaTiposArbol) ObtenerTipo(especie, color, textura string) *TipoArbol {
	clave := especie + "_" + color
	if tipo, existe := f.tipos[clave]; existe {
		return tipo // Reutiliza el flyweight existente en memoria
	}
	fmt.Printf("➕ Creando nuevo Flyweight en RAM para: %s (%s)\n", especie, color)
	nuevo := &TipoArbol{especie: especie, color: color, textura: textura}
	f.tipos[clave] = nuevo
	return nuevo
}

// 3. Contexto (Estado Extrínseco): liviano, único para cada árbol del mapa.
type Arbol struct {
	x, y int
	tipo *TipoArbol // Referencia al flyweight compartido
}

// 4. Cliente / Contenedor: almacena los contextos livianos.
type Bosque struct {
	arboles []Arbol
	fabrica *FabricaTiposArbol
}

func (b *Bosque) Plantar(x, y int, especie, color, textura string) {
	tipoCompartido := b.fabrica.ObtenerTipo(especie, color, textura)
	b.arboles = append(b.arboles, Arbol{x: x, y: y, tipo: tipoCompartido})
}

func main() {
	bosque := &Bosque{fabrica: NuevaFabricaTipos()}

	bosque.Plantar(10, 20, "Pino", "Verde Oscuro", "textura_pino_4k.png")
	bosque.Plantar(15, 25, "Pino", "Verde Oscuro", "textura_pino_4k.png")
	bosque.Plantar(50, 80, "Cerezo", "Rosa", "textura_cerezo_4k.png")
	bosque.Plantar(12, 30, "Pino", "Verde Oscuro", "textura_pino_4k.png")
	bosque.Plantar(55, 85, "Cerezo", "Rosa", "textura_cerezo_4k.png")

	fmt.Printf("\nRenderizando %d árboles usando solo %d objetos TipoArbol en RAM:\n",
		len(bosque.arboles), len(bosque.fabrica.tipos))
	for _, a := range bosque.arboles {
		a.tipo.Dibujar(a.x, a.y)
	}
}
