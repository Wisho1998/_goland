package main
//
//import (
//	"fmt"
//	"time"
//)
//
//type carrito struct {
//	subtotal  float64
//	impuestos float64
//	envio     float64
//	total     float64
//}
//
//func calcCargos(c *carrito) {
//	time.Sleep(3000 * time.Millisecond)
//	c.impuestos = c.subtotal * 0.70
//	c.envio = c.subtotal * 0.10
//}
//
//func calcTotal(c *carrito) {
//	time.Sleep(3000 * time.Millisecond)
//	c.total = c.subtotal + c.impuestos + c.envio
//}
//
//func main() {
//	subtotales := []float64{100.00, 110.00, 105.00, 110.00}
//
//	for _, st := range subtotales {
//		c := carrito{subtotal: st}
//		calcCargos(&c)
//		calcTotal(&c)
//		fmt.Printf("\n %#v \n", c)
//	}
//}
