package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x string
	x = "sdfrgsdfgdfgsdfgsdfgsdfgsdfgsdfgsdfgsdfgsdfg34kjghkhjk3h2j1k3h2j1k32hj13k2g13hj2k1h32jk13h2j1k3h2j3k2hj32k1h3j21kh32j1k3h2j13h2j13h2j13h2j1k3h2j1k3h2j1k3h2j1k3h2j1k32q465436756u7tuye5567eujghjcgv56w43tsr66j6j6j6j6jhjhjhjhjgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg6j56k5l5l3l3lk5k6j7j6k54l4l34k5k5j76j7k6l54"

	fmt.Println("Address of x:", &x)
	fmt.Println("Size of int:", unsafe.Sizeof(x))

	// Получаем указатель на первый байт переменной x
	ptr := uintptr(unsafe.Pointer(&x))

	// Выводим адреса каждого байта переменной x
	for i := 0; i < int(unsafe.Sizeof(x)); i++ {
		fmt.Printf("Byte %d address: 0x%x\n", i, ptr+uintptr(i))
	}
}
