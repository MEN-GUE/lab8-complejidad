package main

import (
	"fmt"
	"time"
)

func function(n int) {
	counter := 0
	for i := n / 2; i <= n; i++ {
		for j := 1; j+n/2 <= n; j++ {
			for k := 1; k <= n; k = k * 2 {
				counter++
			}
		}
	}
}

func main() {
	inputs := []int{1, 10, 100, 1000, 10000, 100000, 1000000}

	fmt.Println("==================================================")
	fmt.Printf("%-15s | %-20s\n", "Tamaño (n)", "Tiempo de Ejecución")
	fmt.Println("==================================================")

	for _, n := range inputs {
		inicio := time.Now()
		
		function(n)
		
		tiempoEjecucion := time.Since(inicio)
		
		fmt.Printf("%-15d | %-20v\n", n, tiempoEjecucion)
	}
	fmt.Println("==================================================")
}