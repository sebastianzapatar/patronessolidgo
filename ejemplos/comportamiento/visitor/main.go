package main

import "fmt"

// Visitante (Visitor): declara un método de visita por cada tipo concreto de nodo del AST.
type VisitanteExpr interface {
	VisitarNumero(n *Numero)
	VisitarSuma(s *Suma)
	VisitarMultiplicacion(m *Multiplicacion)
}

// Elemento (Element): nodo del AST que acepta un visitante (Double Dispatch).
type Expr interface {
	Aceptar(v VisitanteExpr)
}

// Elementos Concretos: nodos del árbol sintáctico abstracto (AST).
type Numero struct{ Valor int }

func (n *Numero) Aceptar(v VisitanteExpr) { v.VisitarNumero(n) }

type Suma struct{ Izq, Der Expr }

func (s *Suma) Aceptar(v VisitanteExpr) { v.VisitarSuma(s) }

type Multiplicacion struct{ Izq, Der Expr }

func (m *Multiplicacion) Aceptar(v VisitanteExpr) { v.VisitarMultiplicacion(m) }

// Visitante Concreto 1: Evaluador de expresiones (acumula el resultado en una pila/campo).
type Evaluador struct{ Resultado int }

func (e *Evaluador) VisitarNumero(n *Numero) { e.Resultado = n.Valor }
func (e *Evaluador) VisitarSuma(s *Suma) {
	s.Izq.Aceptar(e)
	izq := e.Resultado
	s.Der.Aceptar(e)
	e.Resultado = izq + e.Resultado
}
func (e *Evaluador) VisitarMultiplicacion(m *Multiplicacion) {
	m.Izq.Aceptar(e)
	izq := e.Resultado
	m.Der.Aceptar(e)
	e.Resultado = izq * e.Resultado
}

// Visitante Concreto 2: Impresor de árboles en notación infija parentizada.
type ImpresorAST struct{ Salida string }

func (i *ImpresorAST) VisitarNumero(n *Numero) {
	i.Salida = fmt.Sprintf("%d", n.Valor)
}
func (i *ImpresorAST) VisitarSuma(s *Suma) {
	s.Izq.Aceptar(i)
	izq := i.Salida
	s.Der.Aceptar(i)
	i.Salida = fmt.Sprintf("(%s + %s)", izq, i.Salida)
}
func (i *ImpresorAST) VisitarMultiplicacion(m *Multiplicacion) {
	m.Izq.Aceptar(i)
	izq := i.Salida
	m.Der.Aceptar(i)
	i.Salida = fmt.Sprintf("(%s * %s)", izq, i.Salida)
}

func main() {
	// Árbol AST para la expresión: (3 + 5) * 2
	ast := &Multiplicacion{
		Izq: &Suma{Izq: &Numero{3}, Der: &Numero{5}},
		Der: &Numero{2},
	}

	imp := &ImpresorAST{}
	ast.Aceptar(imp)
	fmt.Println("Expresión AST :", imp.Salida)

	eval := &Evaluador{}
	ast.Aceptar(eval)
	fmt.Println("Resultado eval:", eval.Resultado)
}
