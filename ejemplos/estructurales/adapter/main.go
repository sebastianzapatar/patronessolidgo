package main

import "fmt"

// 1. Interfaz con el Cliente: protocolo que el código cliente espera usar.
type ProcesadorPago interface {
	Cobrar(montoUSD float64) string
}

// Servicio compatible existente en nuestra aplicación.
type PasarelaLocal struct{}

func (p PasarelaLocal) Cobrar(montoUSD float64) string {
	return fmt.Sprintf("[PasarelaLocal] Cobro directo procesado: $%.2f USD", montoUSD)
}

// 2. Servicio (incompatible): SDK externo de un tercero que no podemos modificar.
// Usa otro nombre de método, recibe centavos enteros y exige código de moneda.
type SDKStripeExterno struct {
	APIKey string
}

func (s *SDKStripeExterno) CreateChargeInCents(cents int, currency string) string {
	return fmt.Sprintf("[Stripe SDK key=%s] Charge OK: %d cents (%s)", s.APIKey, cents, currency)
}

// 3. Adaptador: implementa ProcesadorPago y envuelve (compone) *SDKStripeExterno.
type AdaptadorStripe struct {
	servicio *SDKStripeExterno
	moneda   string
}

func NuevoAdaptadorStripe(sdk *SDKStripeExterno, moneda string) *AdaptadorStripe {
	return &AdaptadorStripe{servicio: sdk, moneda: moneda}
}

// Cobrar traduce la llamada del cliente al formato y firma que espera el SDK externo.
func (a *AdaptadorStripe) Cobrar(montoUSD float64) string {
	centavos := int(montoUSD * 100) // Conversión de formato: dólares (float64) -> centavos (int)
	respuestaSDK := a.servicio.CreateChargeInCents(centavos, a.moneda)
	return fmt.Sprintf("[Adaptador -> Stripe] %s", respuestaSDK)
}

// 4. Cliente: trabaja únicamente con la interfaz ProcesadorPago sin acoplarse al SDK.
func FinalizarCompra(carrito string, total float64, pasarela ProcesadorPago) {
	fmt.Printf("Cliente pagando %s ($%.2f):\n", carrito, total)
	fmt.Println("  ", pasarela.Cobrar(total))
}

func main() {
	// Uso con el servicio local que ya cumple la interfaz
	local := PasarelaLocal{}
	FinalizarCompra("Curso de Compiladores", 29.99, local)

	// Uso con el servicio incompatible a través del Adaptador
	sdkExterno := &SDKStripeExterno{APIKey: "sk_live_987"}
	adaptador := NuevoAdaptadorStripe(sdkExterno, "USD")
	FinalizarCompra("Suscripción Anual Pro", 149.50, adaptador)
}
