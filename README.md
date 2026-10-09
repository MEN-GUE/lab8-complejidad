# Laboratorio 8 - Análisis de Complejidad

Este repositorio contiene la implementación en Go para el perfilamiento (profiling) y medición de tiempos de ejecución de los algoritmos analizados matemáticamente en el Laboratorio 8.

## Autor
* Juan Menéndez

## Video Demostrativo
La ejecución, medición de tiempos y explicación de las gráficas puede visualizarse en el siguiente enlace:

[Ver demostración en YouTube](https://youtu.be/499CGrmeyFk)

## Estructura del Proyecto
* `cmd/ejercicio1/`: Algoritmo de complejidad superlineal.
* `cmd/ejercicio2/`: Algoritmo de complejidad lineal.
* `cmd/ejercicio3/`: Algoritmo de complejidad cuadrática.
* `cmd/graficador/`: Script para la generación de gráficas en formato PNG.
* `docs/`: Contiene el documento PDF con los análisis asintóticos (Big-Oh) y las gráficas resultantes.

## Instrucciones de Ejecución

Tener Go instalado. Abra una terminal en la raíz de este repositorio.

Para ejecutar el analizador de tiempos de cualquier ejercicio y ver la tabla de resultados, utilice el comando `go run` especificando la ruta del archivo principal. Por ejemplo, para el Ejercicio 1:

```bash
go run ./cmd/ejercicio1/main.go
```