package main

import (
	"fmt"
	"unsafe"
)

type OrderInfo struct {
	OrderCode   rune     // rune is an alias for int32
	Amount      int      // typically 8 bytes on a 64-bit machine
	OrderNumber uint16   // 2 bytes
	Items       []string // slice header is 3 machine words (24 bytes on 64-bit)
	IsReady     bool     // 1 byte
}

// SmallOrderInfo reorders the fields to minimize padding and overall size.
type SmallOrderInfo struct {
	Amount      int      // 8-byte aligned
	Items       []string // 24 bytes for slice header on 64-bit
	OrderCode   rune     // 4 bytes
	IsReady     bool     // 1 byte
	OrderNumber uint16   // 2 bytes (aligned on 2-byte boundary)
}

func main() {
	fmt.Println("=== OrderInfo ===")
	fmt.Printf("Size: %d\n", unsafe.Sizeof(OrderInfo{}))
	fmt.Printf("Offset of OrderCode:   %d\n", unsafe.Offsetof(OrderInfo{}.OrderCode))
	fmt.Printf("Offset of Amount:      %d\n", unsafe.Offsetof(OrderInfo{}.Amount))
	fmt.Printf("Offset of OrderNumber: %d\n", unsafe.Offsetof(OrderInfo{}.OrderNumber))
	fmt.Printf("Offset of Items:       %d\n", unsafe.Offsetof(OrderInfo{}.Items))
	fmt.Printf("Offset of IsReady:     %d\n", unsafe.Offsetof(OrderInfo{}.IsReady))

	fmt.Println("\n=== SmallOrderInfo ===")
	fmt.Printf("Size: %d\n", unsafe.Sizeof(SmallOrderInfo{}))
	fmt.Printf("Offset of Amount:      %d\n", unsafe.Offsetof(SmallOrderInfo{}.Amount))
	fmt.Printf("Offset of Items:       %d\n", unsafe.Offsetof(SmallOrderInfo{}.Items))
	fmt.Printf("Offset of OrderCode:   %d\n", unsafe.Offsetof(SmallOrderInfo{}.OrderCode))
	fmt.Printf("Offset of IsReady:     %d\n", unsafe.Offsetof(SmallOrderInfo{}.IsReady))
	fmt.Printf("Offset of OrderNumber: %d\n", unsafe.Offsetof(SmallOrderInfo{}.OrderNumber))
}
