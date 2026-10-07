package main

import (
	"fmt"
	"time"
)

func funcionAnalisis(n int) {
	operaciones := 0
	for i := 1; i <= n/3; i++ {
		for j := 1; j <= n; j += 4 {
			operaciones++
		}
	}
	_ = operaciones 
}

func main() {
	tamanosInput := []int{1, 10, 100, 1000, 10000, 100000, 1000000}

	fmt.Println("==================================================")
	fmt.Printf("%-15s | %-20s\n", "Tamaño (n)", "Tiempo de Ejecución")
	fmt.Println("==================================================")

	for _, n := range tamanosInput {
		inicio := time.Now()
		
		funcionAnalisis(n)
		
		tiempoEjecucion := time.Since(inicio)
		
		fmt.Printf("%-15d | %-20v\n", n, tiempoEjecucion)
	}
	fmt.Println("==================================================")
}