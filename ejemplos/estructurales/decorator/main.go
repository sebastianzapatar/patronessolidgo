package main

import (
	"fmt"
	"strings"
)

// 1. Componente: interfaz común para el objeto envuelto y para todos los decoradores.
type Notificador interface {
	Enviar(mensaje string) string
}

// 2. Componente Concreto: define el comportamiento base (enviar por correo).
type NotificadorEmail struct {
	Destinatario string
}

func (n NotificadorEmail) Enviar(mensaje string) string {
	return fmt.Sprintf("📧 Email a <%s>: %q", n.Destinatario, mensaje)
}

// 3. Decorador Concreto A: envuelve un Notificador y añade envío por SMS.
type DecoradorSMS struct {
	envuelto Notificador
	telefono string
}

func ConSMS(n Notificador, telefono string) Notificador {
	return DecoradorSMS{envuelto: n, telefono: telefono}
}

func (d DecoradorSMS) Enviar(mensaje string) string {
	previo := d.envuelto.Enviar(mensaje)
	return fmt.Sprintf("%s\n📱 SMS a [%s]: %q", previo, d.telefono, mensaje)
}

// 4. Decorador Concreto B: añade notificación a un canal de Slack.
type DecoradorSlack struct {
	envuelto Notificador
	canal    string
}

func ConSlack(n Notificador, canal string) Notificador {
	return DecoradorSlack{envuelto: n, canal: canal}
}

func (d DecoradorSlack) Enviar(mensaje string) string {
	previo := d.envuelto.Enviar(mensaje)
	return fmt.Sprintf("%s\n💬 Slack [%s]: %q", previo, d.canal, mensaje)
}

// 5. Decorador Concreto C: transforma el mensaje antes de delegarlo a la pila.
type DecoradorUrgente struct {
	envuelto Notificador
}

func ConFormatoUrgente(n Notificador) Notificador {
	return DecoradorUrgente{envuelto: n}
}

func (d DecoradorUrgente) Enviar(mensaje string) string {
	transformado := "[URGENTE] " + strings.ToUpper(mensaje)
	return d.envuelto.Enviar(transformado)
}

func main() {
	// Pila de decoradores: Urgente -> Slack -> SMS -> Email base
	var alerta Notificador = NotificadorEmail{Destinatario: "ops@empresa.co"}
	alerta = ConSMS(alerta, "+57-300-1234567")
	alerta = ConSlack(alerta, "#incidentes-prod")
	alerta = ConFormatoUrgente(alerta)

	fmt.Println(alerta.Enviar("uso de CPU al 98% en nodo principal"))
}
