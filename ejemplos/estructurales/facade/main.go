package main

import "fmt"

// --- Subsistema Complejo: múltiples tipos con orden de inicialización estricto ---

type ServicioInventario struct{}

func (ServicioInventario) VerificarStock(sku string) bool {
	fmt.Printf("  1. [Inventario] Verificando disponibilidad de %s -> OK\n", sku)
	return true
}

type PasarelaFraude struct{}

func (PasarelaFraude) EvaluarRiesgo(cliente string, monto float64) bool {
	fmt.Printf("  2. [Antifraude] Analizando operación de %s por $%.2f -> APROBADO\n", cliente, monto)
	return true
}

type MotorCobros struct{}

func (MotorCobros) Debitar(cliente string, monto float64) string {
	fmt.Printf("  3. [Cobros] Debitando $%.2f a %s -> TX-9081\n", monto, cliente)
	return "TX-9081"
}

type ServicioLogistica struct{}

func (ServicioLogistica) Despachar(sku, cliente string) string {
	fmt.Printf("  4. [Logística] Generando guía de envío para %s (%s) -> GUIA-441\n", sku, cliente)
	return "GUIA-441"
}

// --- Fachada (Facade): interfaz simplificada de alto nivel para el cliente ---

type FachadaCompra struct {
	inventario ServicioInventario
	antifraude PasarelaFraude
	cobros     MotorCobros
	logistica  ServicioLogistica
}

func NuevaFachadaCompra() *FachadaCompra {
	return &FachadaCompra{}
}

// ComprarConUnClic coordina todo el subsistema en el orden correcto.
func (f *FachadaCompra) ComprarConUnClic(cliente, sku string, precio float64) string {
	fmt.Printf("Iniciando compra rápida: %s adquiere %s\n", cliente, sku)
	if !f.inventario.VerificarStock(sku) || !f.antifraude.EvaluarRiesgo(cliente, precio) {
		return "Compra rechazada"
	}
	tx := f.cobros.Debitar(cliente, precio)
	guia := f.logistica.Despachar(sku, cliente)
	return fmt.Sprintf("Pedido completado (Transacción: %s, Guía: %s)", tx, guia)
}

func main() {
	// El cliente solo conoce la Fachada; ignora los 4 subsistemas internos.
	tienda := NuevaFachadaCompra()
	resultado := tienda.ComprarConUnClic("Ana Gómez", "LAPTOP-PRO-16", 1299.00)
	fmt.Println("Resultado:", resultado)
}
