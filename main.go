package main

import (
	"fmt"
	"os"
	"winsole/pefile"
	"winsole/x86"
)

func main() {
	fmt.Println("Let's build a exe")

	//executableBytes := []byte{0x6a, 0x00, 0x68, 0x00, 0x30, 0x40, 0x00, 0x68, 0x17, 0x30, 0x40, 0x00, 0x6a, 0x00, 0xff, 0x15, 0x70, 0x20, 0x40, 0x00, 0x6a, 0x00, 0xff, 0x15, 0x68, 0x20, 0x40, 0x00}
	e := x86.NewExecutable()

	if e != nil {
		println("Done")

		e.MoveI32(5)
		e.AddI32(1)
		e.AddI32(1)
		e.AddI32(1)
		e.PushEax()
		b := e.GetBytes()

		if b != nil {
			println("Made")

			f, _ := os.Create("example.bin")
			f.Write(b)

			p := pefile.NewPeFile(b)

			pefile.Save(p, "test.exe")
		}
	}
}
