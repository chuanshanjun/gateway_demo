package main

import "fmt"

func main() {
	//var a [4]int
	//a[0] = 1
	//i := a[0]
	//fmt.Printf("i is %d\n", i)
	//
	//for index, v := range a {
	//	fmt.Printf("index is %d\n", index)
	//	fmt.Printf("value is %d\n", v)
	//}
	//
	//b := [2]string{"penn", "teller"}
	//
	//fmt.Println(b[0])
	//
	//c := [...]string{"jay", "young"}
	//
	//fmt.Println(c[0])
	//
	//letters := []string{"a", "b", "c", "d"}
	//fmt.Println(letters[0])
	//
	////var s []byte
	////s = make([]byte, 5, 5)
	//// s == []byte{0, 0, 0, 0, 0}
	//
	//s := make([]byte, 5)
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//
	//// 切片零值
	//var v []byte
	//fmt.Println(len(v))
	//fmt.Println(cap(v))
	//
	//fmt.Println("===>>>")
	//
	//b2 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//c2 := b2[1:4]
	//for k, v := range c2 {
	//	fmt.Printf("index is %c\n", k)
	//	fmt.Printf("value is %c\n", v)
	//}
	//
	//c3 := b2[:2]
	//for _, v := range c3 {
	//	fmt.Printf("value is %c", v)
	//}
	//fmt.Println("")
	//
	//fmt.Println("===>>>")
	//
	//c4 := b2[2:]
	//for _, v := range c4 {
	//	fmt.Printf("%c ", v)
	//}
	//fmt.Println("")
	//
	//c5 := b2[:]
	//for _, v := range c5 {
	//	fmt.Printf("%c ", v)
	//}
	//fmt.Println("")
	//
	//fmt.Println("===>>>")
	//x := [3]string{"Лайка", "Белка", "Стрелка"}
	//c6 := x[:]
	//for _, v := range c6 {
	//	fmt.Printf("%s ", v)
	//}
	//fmt.Println("")
	//
	//d := []byte{'r', 'o', 'a', 'd'}
	//e := d[2:]
	//for _, v := range e {
	//	fmt.Printf("%c ", v)
	//}
	//fmt.Println("")
	//
	//e[1] = 'm'
	//for _, v := range e {
	//	fmt.Printf("%c ", v)
	//}
	//fmt.Println("")
	//
	//for _, v := range d {
	//	fmt.Printf("%c ", v)
	//}
	//fmt.Println("")

	var s []byte
	s = make([]byte, 5)
	s = s[2:4]
	fmt.Printf("len %d", len(s))
	fmt.Printf("cap %d", cap(s))
	fmt.Println("")

	s = s[:cap(s)]
	fmt.Printf("len %d", len(s))
	fmt.Printf("cap %d", cap(s))

	fmt.Println("===>>> 切片生长")

	// 切片的生长
	//t := make([]byte, len(s), (cap(s)+1)*2)
	//for i := range s {
	//	t[i] = s[i]
	//}
	//s = t
	//
	//for _, v := range s {
	//	fmt.Println(v)
	//}
	//
	//fmt.Printf("t.length %d", len(t))
	//fmt.Printf("t.cap %d", cap(t))

	t := make([]byte, len(s), (cap(s)+1)*2)
	copy(t, s)
	s = t

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)

	fmt.Println("===>>>")

	for _, v := range p {
		fmt.Printf("%d", v)
	}
	fmt.Println("")

	fmt.Println("===>>>")
	a := make([]int, 1)
	a = append(a, 1, 2, 3)
	for _, v := range a {
		fmt.Printf("%d", v)
	}
	fmt.Println("")

	fmt.Println("===>>>")

	a1 := []string{"John", "Paul"}
	b1 := []string{"George", "Ringo", "Pete"}
	a1 = append(a1, b1...)
	for _, v := range a1 {
		fmt.Printf("%s ", v)
	}
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
