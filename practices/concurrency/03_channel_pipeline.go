package main

import (
	"fmt"
	"time"
)

type carrito struct {
	subtotal  float64
	impuestos float64
	envio     float64
	total     float64
}

func calcCargos(entrada <-chan *carrito, salida chan<- *carrito) {
	for c := range entrada {
		time.Sleep(3000 * time.Millisecond)
		c.impuestos = c.subtotal * 0.70
		c.envio = c.subtotal * 0.10
		salida <- c
	}
	close(salida)
}

func calcTotal(entrada <-chan *carrito, salida chan<- *carrito) {
	for c := range entrada {
		time.Sleep(3000 * time.Millisecond)
		c.total = c.subtotal + c.impuestos + c.envio
		salida <- c
	}
	close(salida)
}

func main() {
	subtotales := []float64{100.00, 110.00, 105.00, 110.00}
	carritos := make(chan *carrito)
	conCargos := make(chan *carrito)
	conTotal := make(chan *carrito)

	go publicar(subtotales, carritos)
	go calcCargos(carritos, conCargos)
	go calcTotal(conCargos, conTotal)
	suscribir(conTotal)
}

func suscribir(conTotal <-chan *carrito) {
	for c := range conTotal {
		fmt.Printf("\n %#v \n", c)
	}
}

func publicar(subtotales []float64, carritos chan<- *carrito) {
	for _, st := range subtotales {
		carritos <- &carrito{subtotal: st}
	}
	close(carritos)
}
