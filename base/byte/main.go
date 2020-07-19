package main

import (
	"encoding/binary"
	"fmt"
)

const Msg_Header = "12345678"

func main() {
	var msg = []byte(Msg_Header)

	for k, v := range msg {
		fmt.Printf("k is %d , v is %c\n", k, v)
	}

	msg[1] = 'A'
	msg[2] = 'B'

	len1 := len(msg)

	fmt.Println("length is", len1)

	len2 := int32(len1)
	fmt.Println("length is", len2)

	fmt.Println("===>>>")

	var a = []byte{0, 1, 2, 3}
	for k, v := range a {
		fmt.Printf("k is %d , v is %v\n", k, v)
	}
	fmt.Println(a)
	fmt.Println(binary.BigEndian.Uint32(a))
	fmt.Println(binary.LittleEndian.Uint32(a))

	//length := binary.BigEndian.Uint32(a)
	//bodyBuf := make([]byte, length)
	//for k, v := range bodyBuf {
	//	fmt.Printf("k is %d , v is %v\n", k, v)
	//}

	var edge = 255
	fmt.Println(edge + 1)
	fmt.Println(uint8(edge + 1))
}
