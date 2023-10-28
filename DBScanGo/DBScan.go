/*
	DbScan -- кластеризация --

Данные:
-- bbox (xl,yl,xr,yr) 	-- строка
-- WIDTH, HEIGHT 		-- целые (пикселы = 256|512)
-- data 				-- срез объектов --
-- minC					-- минимальное количество соседей в окрестности --
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Object struct {
	_ string  //`gorm:"column:object_id; primary_key; not null" json:"object_id"`
	_ string  //`gorm:"column:__table_id" json:"table_id"`
	X float64 //`gorm:"column:x"`
	Y float64 //`gorm:"column:y"`
}

var data []Object // -- объекты из БД --
var bbox string
var WIDTH, HEIGHT int

var length int     // размер data
var minC int = 4   // -- минимальное количество соседей в окрестности --
var radius float64 // -- вычисляется функцией Radius --

var visited []byte   // -- 1-точка просмотрена, 0-точка не просмотрена --
var cluster []int    // -- номера кластеров для каждой точки
var neighbours []int // -- соседи в окрестности = очередь не просмотренных --
var neigh []int      // соседи одной точки

// Функция вычисляет расстояние между двумя точками
func Distance(l Object, r Object) float64 {
	kX := float64(l.X - r.X)
	kY := float64(l.Y - r.Y)
	return math.Hypot(kX, kY)
}

// Функция вычисляет resolution = радиус эпсилон-окрестности кластера
func Radius(bbox string, WIDTH int) float64 {
	// перевод строки в реальные координаты -- обязательная замена запятой на пробел --
	coord := strings.Fields(strings.ReplaceAll(bbox, ",", " ")) // и делим на поля
	var xl, xr float64
	var err error
	xl, err = strconv.ParseFloat(coord[0], 64) // x левой точки
	if err != nil {                            // если что-то не так - это авария полная
		// -- ЧТО ДЕЛАЕМ -- ??? ---
	}
	xr, err = strconv.ParseFloat(coord[2], 64) // x правой точки
	if err != nil {                            // если что-то не так - это авария полная
		// -- ЧТО ДЕЛАЕМ -- ??? ---
	}
	// -- xr > xl -- (bbox: xr - правая точка, xl - левая точка) --
	radius := math.Abs((xr - xl) / float64(WIDTH))
	return radius
}

// Функция считает количество соседей и помещает их в список соседей
// k - индекс точки в массиве объектов
func region(k int) int {
	countP := 0                       // просматриваемая точка НЕ учитывается
	neigh = make([]int, 0, len(data)) // соседи текущей точки
	for i := 0; i < len(data); i++ {
		if i != k { // текущую (k-ю) точку не смотрим
			d := Distance(data[i], data[k]) // расстояние до всех
			if d < radius+0.00001 {         // потом посмотреть погрешность - 6 знаков -- ??? --
				if cluster[i] == 0 { // если точка не посещалась
					neigh = append(neigh, i) // -- поместить в соседи --
				}
				countP++ // считаются все - см. expandCluster() --
			}
		}
	}
	//fmt.Println("region", k, countP, len(neigh), neigh)
	return countP // количество соседей
}

// Функция расширяет список соседей и нумерует точки кластером
/* параметры:
- номер кластера
- количество соседей
- индекс точки в data
*/
func expandCluster(cluster_no int, num_point int, next_point int) {
	cluster[next_point] = cluster_no // для текущей точки номер кластера присвоили
	countP := 0
	//fmt.Println("expan1", next_point, num_point, cluster_no, neighbours)
	for i := 0; i < num_point; i++ { // для всех соседей текущей точки
		if visited[neighbours[i]] == 0 { // если соседку не смотрели ранее
			visited[neighbours[i]] = 1
			// надо посчитать соседей не просмотренной точки
			// соседи текущей точки создаются внутри --
			countP = region(neighbours[i]) // в окрестность могут входить и просмотренные ранее точки
			// -- в neigh попадают только не кластеризованные точки --
			if (countP + 1) >= minC { // учли текущую рассматриваемую точку
				neighbours = append(neighbours, neigh...) // добавили соседей
				num_point += len(neigh)                   // увеличили количество посещаемых соседей точек - это размер цикла по i = размер neighbours
			}
		}
		// -- если точке ранее не был назначен кластер, то назначается --
		// -- и для просмотренной точки (ее соседи уже вычислялись) --
		if cluster[neighbours[i]] == 0 {
			cluster[neighbours[i]] = cluster_no
			//fmt.Println(neighbours[i], cluster_no)
		}
		//fmt.Println("expan2", next_point, num_point, cluster_no, neighbours)
	}
}
func Begin(bbox string, WIDTH int) {
	radius = Radius(bbox, WIDTH) // resolution = радиус эпсилон-окрестности точки
	minC = 4                     // вместе с точкой
	// создаем необходимые срезы
	visited = make([]byte, len(data)) // признак просмотра точки
	//t := len(data)/minC + 1
	//fmt.Println(t)
	cluster = make([]int, len(data)) // номера кластеров для точек
	//neighbours = make([]int, 0, len(data)*len(data)) // соседи - общий список (по алгоритму размер = len(data)^2 )
}

// Функция DBScan - главная функция кластеризации
func DBScan(radius float64, minC int) {

	next_cluster := 0  // текущий номер кластера
	var num_points int // количество соседей
	for next_point := 0; next_point < len(data); next_point++ {
		// -- для очередной не просмотренной точки --
		if visited[next_point] == 0 {
			visited[next_point] = 1 // точка просмотрена
			// создаем соседей для новой точки
			neighbours = make([]int, 0, len(data)*len(data)) // соседи текущей точки
			num_points = region(next_point)                  // количество соседей i-й точки в окрестности

			if (num_points + 1) < minC { // учитываем текущую точку
				// -- некластеризованная точка: cluster[i] = 0 --
				// посчитаны расстояния от текущей точки до всех других - у нее нет minC соседей
			} else {
				next_cluster++                            // новый номер кластера
				neighbours = append(neighbours, neigh...) // прицепили соседей текущей точки
				fmt.Println("DBScan", next_point, num_points, next_cluster, neighbours)
				expandCluster(next_cluster, num_points, next_point) // посмотреть всех соседей точки i
			}
		}
		// если точка ранее просмотрена, то она пропускается
	}
	// -- разобраться с некластеризованными точками --
	fmt.Println("cluster = ", next_cluster)
}

func main() {
	fmt.Println("Кластеризация точек DBScanGo")
	data = make([]Object, 20) // -- cоздать срез data --

	//RandomPoints(500) // максимальная координата по обеим осям
	data20() // конкретные кластеры

	// начальные значения
	bbox = "0.0,0.0,1024.0,1024.0" // resolution = radius = 4.0
	WIDTH = 256

	Begin(bbox, WIDTH) // minC присваивает
	fmt.Println(len(data), len(cluster))

	start := time.Now()
	DBScan(radius, minC)
	finish := time.Now()
	t := finish.Sub(start)
	fmt.Println(t)
	printResult()
}

// -- просто вывод точек --
func printPoint(t Object) {
	fmt.Printf("<%4.1f,%4.1f>", t.X, t.Y)
}
func printAllPoints() {
	for _, t := range data {
		printPoint(t)
	}
	fmt.Println()
}

// -- заполнение конкретным набором точек --
func data20() {
	// -- первый кластер --
	data[0].X = 1
	data[0].Y = 1
	data[1].X = 3
	data[1].Y = 3
	data[2].X = 0
	data[2].Y = 4
	data[3].X = 4
	data[3].Y = 2
	data[4].X = 3
	data[4].Y = 0

	data[5].X = 10
	data[5].Y = 10
	data[6].X = 12
	data[6].Y = 12
	data[7].X = 14
	data[7].Y = 14
	data[8].X = 16
	data[8].Y = 16
	data[9].X = 10
	data[9].Y = 15
	data[10].X = 14
	data[10].Y = 12

	data[11].X = 5
	data[11].Y = 15
	data[12].X = 3
	data[12].Y = 15
	data[13].X = 3
	data[13].Y = 17
	data[14].X = 6
	data[14].Y = 18
	data[15].X = 2
	data[15].Y = 12

	data[16].X = 18
	data[16].Y = 3

	data[17].X = 1
	data[17].Y = 9

	data[18].X = 10
	data[18].Y = 5

	data[19].X = 6
	data[19].Y = 9
}

// -- заполнение случайными точками --
func RandomPoints(MaxR int) {
	seed := rand.NewSource(time.Now().UnixNano()) // создание объекта Source для инициализации датчика случайных чисел временем
	kk := rand.New(seed)                          // создание нового объекта Rand с инициализацией
	for i := range data {
		data[i].X = float64(kk.Intn(MaxR))
		data[i].Y = float64(kk.Intn(MaxR))
	}
}

// -- окончательный вывод всего --
func printResult() {

	for k, t := range data {
		fmt.Printf("%2d <%4.1f,%4.1f> ", k, t.X, t.Y)
		fmt.Printf("%2d %d\n", cluster[k], visited[k])
	}

	maximum := cluster[0]
	for _, t := range cluster {
		if t > maximum {
			maximum = t
		}
	}
	fmt.Println("max = ", maximum)

	countC := make([]int, maximum+1)
	for i := range cluster {
		countC[cluster[i]]++
	}
	summa := 0
	for i := range countC {
		summa += countC[i]
	}
	fmt.Println(summa, countC)
}
