package main

import (
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func main() {
	// Asegurar que la carpeta docs exista
	err := os.MkdirAll("docs", 0755)
	if err != nil {
		log.Fatalf("Error creando directorio docs: %v", err)
	}

	// Datos del Ejercicio 1 (O(n^2 log n)) - Convertidos a milisegundos
	datosE1 := plotter.XYs{
		{X: 1, Y: 0.000711},
		{X: 10, Y: 0.000652},
		{X: 100, Y: 0.048038},
		{X: 1000, Y: 6.217346},
		{X: 10000, Y: 497.096717},
		{X: 100000, Y: 53376.265852}, // 53.37s
	}
	generarGrafico(datosE1, "Ejercicio 1 - Complejidad O(n^2 log n)", "docs/grafica_ejercicio1.png", color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Datos del Ejercicio 2 (O(n)) - Convertidos a milisegundos
	datosE2 := plotter.XYs{
		{X: 1, Y: 0.000581},
		{X: 10, Y: 0.000320},
		{X: 100, Y: 0.000381},
		{X: 1000, Y: 0.001353},
		{X: 10000, Y: 0.011141},
		{X: 100000, Y: 0.071243},
		{X: 1000000, Y: 0.716808},
	}
	generarGrafico(datosE2, "Ejercicio 2 - Complejidad O(n)", "docs/grafica_ejercicio2.png", color.RGBA{R: 0, G: 128, B: 0, A: 255})

	// Datos del Ejercicio 3 (O(n^2)) - Convertidos a milisegundos
	datosE3 := plotter.XYs{
		{X: 1, Y: 0.000571},
		{X: 10, Y: 0.000291},
		{X: 100, Y: 0.000932},
		{X: 1000, Y: 0.067978},
		{X: 10000, Y: 5.975296},
		{X: 100000, Y: 432.299803},
		{X: 1000000, Y: 42593.614156}, // 42.59s
	}
	generarGrafico(datosE3, "Ejercicio 3 - Complejidad O(n^2)", "docs/grafica_ejercicio3.png", color.RGBA{R: 0, G: 0, B: 255, A: 255})
}

func generarGrafico(datos plotter.XYs, titulo string, rutaArchivo string, colorLinea color.RGBA) {
	p := plot.New()

	p.Title.Text = titulo
	p.X.Label.Text = "Tamaño de Input (n)"
	p.Y.Label.Text = "Tiempo de Ejecución (ms)"

	linea, puntos, err := plotter.NewLinePoints(datos)
	if err != nil {
		log.Fatalf("Error creando línea: %v", err)
	}

	linea.Color = colorLinea
	puntos.Color = colorLinea

	p.Add(linea, puntos)

	// Guardar en formato PNG con dimensiones 800x600
	if err := p.Save(8*vg.Inch, 6*vg.Inch, rutaArchivo); err != nil {
		log.Fatalf("Error guardando gráfica: %v", err)
	}
	log.Printf("Gráfica generada exitosamente: %s\n", rutaArchivo)
}