package main

import (
	"fmt"
	"time"
)

// Ejercicio 2, 
func funcionAnalisis(n int) {
	if n <= 1 {
		return
	}
	
	operaciones := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			operaciones++ // Representa la operación de costo O(1) del printf
			break         // Trampa estructural que reduce el costo a O(n)
		}
	}
	
	_ = operaciones // Evita errores de variable no utilizada en Go
}

func main() {
	// Tamaños de input requeridos por el laboratorio
	tamanosInput := []int{1, 10, 100, 1000, 10000, 100000, 1000000}

	fmt.Println("==================================================")
	fmt.Printf("%-15s | %-20s\n", "Tamaño (n)", "Tiempo de Ejecución")
	fmt.Println("==================================================")

	for _, n := range tamanosInput {
		// Punto de inicio del profiling
		inicio := time.Now()
		
		funcionAnalisis(n)
		
		// Cálculo del delta de tiempo
		tiempoEjecucion := time.Since(inicio)
		
		fmt.Printf("%-15d | %-20v\n", n, tiempoEjecucion)
	}
	fmt.Println("==================================================")
}