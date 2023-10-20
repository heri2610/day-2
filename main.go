package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Masukkan bilangan: ")
	fmt.Scanf("%d", &n)

	isPrime := isPrime(n)
	if isPrime {
		fmt.Println(n, "adalah bilangan prima")
	} else {
		fmt.Println(n, "bukan bilangan prima")
	}
	if n < 7 {
		fmt.Println(n, "kurang dari 7")
	} else {
		kelipatan7 := kelipatan7(n)
		fmt.Println("Bilangan kelipatan 7 dari 7 hingga", n, "adalah:")
		for _, k := range kelipatan7 {
			fmt.Println(k)
		}
	}

	luas := luasTrapesium(10, 20, 15)
	fmt.Println("Luas trapesium dengan alas pertama 10, alas kedua 20, dan tinggi 15 adalah:", luas)
}

func isPrime(n int) bool {
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}

func kelipatan7(n int) []int {
	kelipatan7 := []int{}
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			kelipatan7 = append(kelipatan7, i)
		}
	}
	return kelipatan7
}

func luasTrapesium(alas1, alas2, tinggi int) int {
	luas := (alas1 + alas2) * tinggi / 2
	return luas
}
