/*  Вычисление расстояний между точками на плоскости
    -- параллельный вариант --
	Для параллельности нужно добавлять к-ю строку расстояний в срез
	Иначе получается гонка
	-- или сделать go-рутины по строке
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Point struct {
	X, Y int
}

// -- просто вывод точек --
func printPoint(t Point) {
	fmt.Printf("<%d,%d> ", t.X, t.Y)
}
func printAllPoints() {
	for i := range points {
		printPoint(points[i])
	}
	fmt.Println()
}

const length = 15000 // количество точек на карте

// Генерация правдоподобных данных по России для экспериментов
// Y = rand(север - юг) + юг
// X = rand(восток - запад) + запад
//const MaxR = 10 // -- диапазон чисел --

var points [length]Point              // -- множество точек --
var distances [length][length]float64 // -- все расстояния от одной точки до всех других --

// -- заполнение случайными точками --
func RandomPoints(MaxR int) {
	seed := rand.NewSource(time.Now().UnixNano()) // создание объекта Source для инициализации датчика случайных чисел временем
	kk := rand.New(seed)                          // создание нового объекта Rand с инициализацией
	for i := range points {
		points[i].X = kk.Intn(MaxR)
		points[i].Y = kk.Intn(MaxR)
	}
	//printAllPoints()
}

// Функция вычисляет расстояние между двумя точками
func Distance(l Point, r Point) float64 {
	kX := float64(l.X - r.X)
	kY := float64(l.Y - r.Y)
	return math.Hypot(kX, kY)
}

// -- считаем расстояния от одной точки до всех других --
// -- расчет от точки k вправо; влево вычислено в предыдущих строках --
func Do(k int) {
	distances[k][k] = 0
	current := points[k]
	for i := k + 1; i < length; i++ {
		distances[k][i] = Distance(current, points[i])
	}
}

// Вычисляет все расстояния между всеми точками -- последовательная схема --
// Развертывание не помогает
func DoAll() {

	k := 0
	for k < length {
		Do(k)
		Do(k + 1)
		Do(k + 2)
		Do(k + 3)
		k += 4
	}
	/*
		for k := range distances {
			Do(k)
		}
	*/
}

// -- вывод одной строки расстояний --
func printDistances() {
	for n := range distances {
		fmt.Printf("%f ", distances[n])
	}
	fmt.Println()
}

// -- параллельные вычисления --
func DoPar() {

	var wg sync.WaitGroup
	var k int = 0
	for k = 0; k < length; k++ {
		wg.Add(1)
		go func(i int) {
			Do(i) // вычисление расстояния по строке матрицы
			//fmt.Println(i)
			wg.Done()
		}(k)
	}
	wg.Wait()

}
func main() {
	fmt.Println("Вычисление расстояний между точками")
	RandomPoints(1000)

	start := time.Now()
	DoPar()
	finish := time.Now()
	t := finish.Sub(start)
	fmt.Println(t)

}
