package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Тип вектор
type Vector struct {
	v []float64 //Значение
	n int       //Длина
}

type Network struct {
	weights []Matrix
	layersN int
}

type Matrix struct {
	m int //Размеры
	n int
	v [][]float64 //значения
}

func newNetwork(sizes []int, r *rand.Rand) (Network, error) {
	//{2, 6, 2}
	err := error(nil)
	var net Network
	net.layersN = len(sizes) - 1
	net.weights = make([]Matrix, len(sizes))
	for i := 0; i < len(sizes); i++ {
		net.weights[i].v = make([][]float64, sizes[i])
	}

	for i := 0; i < len(sizes); i++ {
		for j := 1; j < len(sizes); j++ {
			net.weights[i].v[j-1] = make([]float64, sizes[j-1])
		}
	}

	for i := 0; i < len(sizes)-1; i++ {
		net.weights[i], err = newMatrix(sizes[i], sizes[i+1], r)
		net.weights[i].m = sizes[i]
		net.weights[i].n = sizes[i]
	}
	return net, err
}

func printMatrix(matrix Matrix) {
	println(matrix.m)
	println(matrix.n)
	for i := 0; i < matrix.m; i++ {
		for j := 0; j < matrix.n; j++ {
			fmt.Print(matrix.v[i][j])
		}
		fmt.Println()
	}
}

func newMatrix(m int, n int, r *rand.Rand) (Matrix, error) {
	err := error(nil)
	var rMatrix Matrix
	rMatrix.v = make([][]float64, m)

	for i := 0; i < n; i++ {
		rMatrix.v[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rMatrix.v[i][j] = r.Float64() - 0.5
		}
	}

	return rMatrix, err
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	layers := []int{2, 6, 2}

	network, err := newNetwork(layers, r1)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(network)
	printMatrix(network.weights[1])
}
