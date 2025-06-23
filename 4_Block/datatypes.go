package main

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

func main() {
	var (
		a int        = -1234567890
		b int8       = -128
		c int16      = -32768
		d int32      = -2147483648
		e int64      = -9223372036854775808
		f uint       = 1234567890
		g uint8      = 255
		h uint16     = 65535
		i uint32     = 4294967295
		j uint64     = 18446744073709551615
		k uintptr    = 0xdeadbeef
		m float32    = 3.1415926535
		n float64    = 3.14159265358979323846
		o complex64  = 3 + 2i
		p complex128 = cmplx.Sqrt(-5 + 12i)
		q byte       = 65
		r string     = "Hello, world!"
		s bool       = true
	)

	fmt.Println("=== Целые числа ===")
	fmt.Printf("int:    %d (размер: %d байт)\n", a, unsafe.Sizeof(a))
	fmt.Printf("int8:   %d (размер: %d байт)\n", b, unsafe.Sizeof(b))
	fmt.Printf("int16:  %d (размер: %d байт)\n", c, unsafe.Sizeof(c))
	fmt.Printf("int32:  %d (размер: %d байт)\n", d, unsafe.Sizeof(d))
	fmt.Printf("int64:  %d (размер: %d байт)\n", e, unsafe.Sizeof(e))
	fmt.Printf("uint:   %d (размер: %d байт)\n", f, unsafe.Sizeof(f))
	fmt.Printf("uint8:  %d (размер: %d байт)\n", g, unsafe.Sizeof(g))
	fmt.Printf("uint16: %d (размер: %d байт)\n", h, unsafe.Sizeof(h))
	fmt.Printf("uint32: %d (размер: %d байт)\n", i, unsafe.Sizeof(i))
	fmt.Printf("uint64: %d (размер: %d байт)\n", j, unsafe.Sizeof(j))
	fmt.Printf("uintptr: %#x (размер: %d байт)\n", k, unsafe.Sizeof(k))

	fmt.Println("\n=== Числа с плавающей точкой ===")
	fmt.Printf("float32: %.7f (размер: %d байт)\n", m, unsafe.Sizeof(m))
	fmt.Printf("float64: %.15f (размер: %d байт)\n", n, unsafe.Sizeof(n))

	fmt.Println("\n=== Комплексные числа ===")
	fmt.Printf("complex64:  %v (размер: %d байт)\n", o, unsafe.Sizeof(o))
	fmt.Printf("complex128: %v (размер: %d байт)\n", p, unsafe.Sizeof(p))

	fmt.Println("\n=== Байт и строка ===")
	fmt.Printf("byte:   %d (символ: '%c') (размер: %d байт)\n", q, q, unsafe.Sizeof(q))
	fmt.Printf("string: %q (размер: %d байт)\n", r, unsafe.Sizeof(r))

	fmt.Println("\n=== Логический тип ===")
	fmt.Printf("bool: %t (размер: %d байт)\n", s, unsafe.Sizeof(s))
}
