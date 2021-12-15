package main

import (
	"fmt"
	"github.com/fatih/color"
)

type PC struct {
	ram   int
	brand string
	disk  int
}

func (myPC PC) String() string {
	str := color.HiYellowString("My [Pc] is/have:")
	str += color.HiMagentaString("\n * brand %s", myPC.brand)
	str += color.HiCyanString("\n * %d GB of RAM", myPC.ram)
	str += color.HiGreenString("\n * %d GB od disk", myPC.disk)
	return str
}

func main() {
	myPC := PC{ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPC)
}
