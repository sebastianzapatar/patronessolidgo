package main

import "fmt"

// 1. Implementación (interfaz de plataforma): operaciones primitivas de bajo nivel.
type Dispositivo interface {
	EstaEncendido() bool
	Encender()
	Apagar()
	Volumen() int
	FijarVolumen(porcentaje int)
	Nombre() string
}

// 2. Implementación Concreta A: Televisor.
type Televisor struct {
	encendido bool
	volumen   int
}

func (t *Televisor) EstaEncendido() bool     { return t.encendido }
func (t *Televisor) Encender()               { t.encendido = true; fmt.Println("  [TV] Encendido") }
func (t *Televisor) Apagar()                 { t.encendido = false; fmt.Println("  [TV] Apagado") }
func (t *Televisor) Volumen() int            { return t.volumen }
func (t *Televisor) FijarVolumen(v int)      { t.volumen = v; fmt.Printf("  [TV] Volumen = %d%%\n", t.volumen) }
func (t *Televisor) Nombre() string          { return "Smart TV 4K" }

// 2. Implementación Concreta B: Radio.
type Radio struct {
	encendido bool
	volumen   int
}

func (r *Radio) EstaEncendido() bool     { return r.encendido }
func (r *Radio) Encender()               { r.encendido = true; fmt.Println("  [Radio] Encendida") }
func (r *Radio) Apagar()                 { r.encendido = false; fmt.Println("  [Radio] Apagada") }
func (r *Radio) Volumen() int            { return r.volumen }
func (r *Radio) FijarVolumen(v int)      { r.volumen = v; fmt.Printf("  [Radio] Volumen = %d%%\n", r.volumen) }
func (r *Radio) Nombre() string          { return "Radio FM Vintage" }

// 3. Abstracción: control de alto nivel que contiene el "puente" (dispositivo).
type ControlRemoto struct {
	dispositivo Dispositivo
}

func NuevoControlRemoto(d Dispositivo) *ControlRemoto {
	return &ControlRemoto{dispositivo: d}
}

func (c *ControlRemoto) BotonEncendido() {
	fmt.Printf("Control -> alternando energía en %s\n", c.dispositivo.Nombre())
	if c.dispositivo.EstaEncendido() {
		c.dispositivo.Apagar()
	} else {
		c.dispositivo.Encender()
	}
}

func (c *ControlRemoto) SubirVolumen() {
	c.dispositivo.FijarVolumen(c.dispositivo.Volumen() + 15)
}

// 4. Abstracción Refinada: extiende la abstracción por composición (embedding).
type ControlAvanzado struct {
	*ControlRemoto
}

func NuevoControlAvanzado(d Dispositivo) *ControlAvanzado {
	return &ControlAvanzado{ControlRemoto: NuevoControlRemoto(d)}
}

func (ca *ControlAvanzado) Silenciar() {
	fmt.Printf("Control Avanzado -> silenciando %s\n", ca.dispositivo.Nombre())
	ca.dispositivo.FijarVolumen(0)
}

func main() {
	tv := &Televisor{volumen: 30}
	controlBasico := NuevoControlRemoto(tv)
	controlBasico.BotonEncendido()
	controlBasico.SubirVolumen()

	radio := &Radio{volumen: 50}
	controlPro := NuevoControlAvanzado(radio)
	controlPro.BotonEncendido()
	controlPro.SubirVolumen()
	controlPro.Silenciar()
}
