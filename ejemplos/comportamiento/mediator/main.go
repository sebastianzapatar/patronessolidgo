package main

import "fmt"

// 2. Mediador (interfaz de comunicación centralizada)
type MediadorTorre interface {
	SolicitarPista(a Aeronave) bool
	NotificarDespegue(a Aeronave)
}

// 1. Componente (interfaz de las aeronaves que colaboran sin conocerse)
type Aeronave interface {
	Identificador() string
	Aterrizar()
	Despegar()
	PermitirAterrizaje()
}

// 3. Mediador Concreto: coordina la pista y la cola de espera
type TorreControl struct {
	pistaLibre bool
	colaEspera []Aeronave
}

func NuevaTorreControl() *TorreControl {
	return &TorreControl{pistaLibre: true}
}

func (t *TorreControl) SolicitarPista(a Aeronave) bool {
	if t.pistaLibre {
		t.pistaLibre = false
		return true
	}
	t.colaEspera = append(t.colaEspera, a)
	return false
}

func (t *TorreControl) NotificarDespegue(a Aeronave) {
	if len(t.colaEspera) > 0 {
		siguiente := t.colaEspera[0]
		t.colaEspera = t.colaEspera[1:]
		siguiente.PermitirAterrizaje()
	} else {
		t.pistaLibre = true
	}
}

// 4. Componente Concreto: solo se comunica con el mediador
type Avion struct {
	codigo string
	torre  MediadorTorre
}

func (av *Avion) Identificador() string { return av.codigo }

func (av *Avion) Aterrizar() {
	if av.torre.SolicitarPista(av) {
		fmt.Printf("  ✈️  [%s] Pista asignada → aterrizando ahora\n", av.codigo)
	} else {
		fmt.Printf("  ⏳ [%s] Pista ocupada → esperando turno en el aire\n", av.codigo)
	}
}

func (av *Avion) Despegar() {
	fmt.Printf("  🛫 [%s] Ha liberado la pista → notificando a la torre\n", av.codigo)
	av.torre.NotificarDespegue(av)
}

func (av *Avion) PermitirAterrizaje() {
	fmt.Printf("  🗼 [Torre → %s] Pista libre, autorizado para aterrizar\n", av.codigo)
	fmt.Printf("  ✈️  [%s] Aterrizando tras espera\n", av.codigo)
}

func main() {
	torre := NuevaTorreControl()
	vuelo1 := &Avion{codigo: "AV-101", torre: torre}
	vuelo2 := &Avion{codigo: "LA-205", torre: torre}
	vuelo3 := &Avion{codigo: "CM-309", torre: torre}

	vuelo1.Aterrizar()
	vuelo2.Aterrizar()
	vuelo3.Aterrizar()
	fmt.Println()
	vuelo1.Despegar()
	vuelo2.Despegar()
}
