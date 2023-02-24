package main

import (
	"fmt"
)

type arbol struct {
	dato       int
	izdo, dcho *arbol
}

func CrearNodo(valor int) *arbol {
	arb := new(arbol)
	arb.dato = valor
	arb.izdo = nil
	arb.dcho = nil
	return arb
}
func InsertarNodo(arbolito *arbol, dato int) *arbol {
	if arbolito == nil {
		return CrearNodo(dato)

	} else if dato == arbolito.dato {
		fmt.Println("El valor ya existe")
		return arbolito

	} else if dato > arbolito.dato {
		arbolito.dcho = InsertarNodo(arbolito.dcho, dato)
		return arbolito
	} else {
		arbolito.izdo = InsertarNodo(arbolito.izdo, dato)
		return arbolito
	}
}
func Preorden(arbolito *arbol) {
	if arbolito != nil {
		fmt.Print(arbolito.dato, ", ")
		Preorden(arbolito.izdo)
		Preorden(arbolito.dcho)
	}
}

func main() {
	var nodo *arbol

	nodo = InsertarNodo(nodo, 5)
	nodo = InsertarNodo(nodo, 3)
	nodo = InsertarNodo(nodo, 7)
	nodo = InsertarNodo(nodo, 2)
	nodo = InsertarNodo(nodo, 4)
	nodo = InsertarNodo(nodo, 6)
	nodo = InsertarNodo(nodo, 8)

	fmt.Println("Preorden: ")
	Preorden(nodo)
}
