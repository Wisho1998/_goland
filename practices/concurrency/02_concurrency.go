package main
//
//import (
//	"fmt"
//	"sync"
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
//func calcCargos(c *carrito, wg *sync.WaitGroup) {
//	defer wg.Done()
//	time.Sleep(3000 * time.Millisecond)
//	c.impuestos = c.subtotal * 0.70
//	c.envio = c.subtotal * 0.10
//}
//
//func calcTotal(c *carrito, wg *sync.WaitGroup) {
//	defer wg.Done()
//	time.Sleep(3000 * time.Millisecond)
//	c.total = c.subtotal + c.impuestos + c.envio
//}
//
//func main() {
//	subtotales := []float64{100.00, 110.00, 105.00, 110.00}
//	for _, st := range subtotales {
//		var wg = sync.WaitGroup{}
//		c := carrito{subtotal: st}
//		wg.Add(2)
//		// las go routine no nos garantizan el orden de ejecución.
//		go calcCargos(&c, &wg)
//		go calcTotal(&c, &wg)
//		wg.Wait()
//		fmt.Printf("\n %#v \n", c)
//	}
//}
