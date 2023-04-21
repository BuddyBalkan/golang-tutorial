package main

import (
	"encoding/json"
	"fmt"
	"kk.com/generics/msggame"
	"math/rand"
	"strings"
)

type Number interface {
	int64 | float64
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"1st": 34,
		"2nd": 12,
		"3rd": 58,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"1st": 35.98,
		"2nd": 26.99,
		"3rd": 47.03,
	}

	// 不使用泛型 为每一种类型添加一种同一过程的方法
	fmt.Printf("Non-Generic Sums: %v and %v \n",
		SumInts(ints),
		SumFloats(floats))

	// 使用泛型 并在调用方法的时候 声明符合方法中定义的泛型里具体的一种
	fmt.Printf("Generic Sums: %v and %v \n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// 使用泛型 并在调用方法的时候不进行声明泛型的具体类型 由编译器自行推断
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v \n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	// 使用定义的Number的泛型
	fmt.Printf("Genric Sums with Constraint: %v and %v \n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts adds together the values of m.
// 整数累加 map的集合
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// Sumloats adds together the values of m.
// 小数累加 map的集合
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOfFloats sums the values of map m.
// It supports both int64 and float64 as types for map values.
// 泛型的数据累加 map的集合
// comparable -- map的key值必须为comparable的；任何可以用运算符“==”和“！=”操作的类型
// 由于已经定义了K和V的泛型 故 m变量是可以使用的
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m, It supports both integers and floats as map values.
// 支持类型Number（上文定义的）泛型的数据累加 map的集合
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// 以下内容希望测试实现多态

// 要是数组行不通则使用以下泛型试试？
//type Cell interface {
//	NormalCell | CrossCell
//	hasFinishTarget()                  // 获取方格达成目标没有
//	ChangeFillStatusByPath(path *Path) // 通过答案中的Path来改变本cell的填充状态
//}

type Param struct {
	Z int
}

type Cell struct {
	X     int
	Y     int
	Total int
	//Summarize func(param *Param)
}

type NormalCell struct {
	Cell
}

type SpecialCell struct {
	Cell
	Xi int
	Yi int
}

func (nC *NormalCell) Summarize(param *Param) {
	//nC.Total = nC.X + nC.Y + param.Z
	nC.Cell.Summarize(param)
	nC.Total += param.Z
}

func (sC *SpecialCell) Summarize(param *Param) {
	//sC.Total = sC.X + sC.Y + sC.Xi + sC.Yi + param.Z
	sC.Cell.Summarize(param)
	sC.Total += sC.Xi + sC.Yi + param.Z
}

func (c *Cell) Summarize(param *Param) {
	c.Total = c.X + c.Y
}

type Point struct {
	X int
	Y int
}

type Path struct {
	PrevPath  *Path
	HerePoint *Point
	NextPath  *Path
}

func (p *Path) addPrev(path *Path) {
	p.PrevPath = path
}

func (p *Path) addNext(path *Path) {
	p.NextPath = path
}

func (p *Point) ToString() string {
	return fmt.Sprintf("this Point is (%v, %v).", p.X, p.Y)
}

func JsonStrToConnectionPathSlice(jsonStr string) []*msggame.ConnectionPath {
	countCPath := strings.Count(jsonStr, "Pair")
	paths := make([]*msggame.ConnectionPath, 0, countCPath)
	_ = json.Unmarshal([]byte(jsonStr), &paths)
	return paths
}

type PCB struct {
	Size  int
	Mark  string
	X     int
	Y     int
	Total int
}

func (p *PCB) BeforeInit(size int) {
	p.Size = size
	p.X = rand.Intn(p.Size)
	p.Y = rand.Intn(p.Size)
	p.Mark = fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func (p *PCB) Init() {
	p.Total = p.X + p.Y
}
