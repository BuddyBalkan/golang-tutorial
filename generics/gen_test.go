package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"kk.com/generics/msggame"
	"kk.com/generics/msggameNoZero"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGeneration(t *testing.T) {
	cells := make([]Cell, 0, 4)
	for i := 0; i < cap(cells); i++ {
		if i%2 == 0 {
			n := &NormalCell{
				Cell: Cell{
					X:     i,
					Y:     int(math.Pow(float64(i), 2.0)),
					Total: 0,
				},
			}
			cells = append(cells, n.Cell)
		} else {
			s := &SpecialCell{
				Cell: Cell{
					X:     i,
					Y:     int(math.Pow(float64(i), 2.0)),
					Total: 0,
				},
				Xi: i - 1,
				Yi: int(math.Pow(float64(i), 2.0)) + 2,
			}
			cells = append(cells, s.Cell)
		}
	}

	for _, cell := range cells {
		p := &Param{Z: 9}
		cell.Summarize(p)
		t.Logf("this is the result summarize: %v", cell.Total)
	}
}

func TestPathLink(t *testing.T) {
	pathList := make([]*Point, 0, 6)
	for i := 0; i < cap(pathList); i++ {
		p := &Point{
			X: i,
			Y: i + 2,
		}
		pathList = append(pathList, p)
	}
	l := list.New()                 // 链表初始化
	prt := l.PushFront(pathList[0]) // 定义链表至第一个元素
	for i := 0; i < len(pathList)-1; i++ {
		prt = l.InsertAfter(pathList[i+1], prt) // 修正指针到下一个
	}
	for e := l.Front(); e != nil; e = e.Next() { // 链表遍历迭代方法
		t.Logf("this is element in list %v", e.Value.(*Point).ToString())
	}
}

// 验证指针强转的第二个返回值效果
func TestListConversion(t *testing.T) {
	path := make([]*Param, 0, 3)
	for i := 0; i < cap(path); i++ {
		e := &Param{Z: int(math.Pow(float64(i), 2.0))}
		path = append(path, e)
	}
	l := list.New()
	prt := l.PushFront(path[0])
	for i := 0; i < len(path)-1; i++ {
		prt = l.InsertAfter(path[i+1], prt)
	}
	for e := l.Front(); e != nil; e = e.Next() { // 链表遍历迭代方法
		point, ok := e.Value.(*Point)
		if !ok {
			t.Error("解析成point出错")
		} else {
			t.Logf("this is point element in list %v", point.ToString())
		}
		param, ok2 := e.Value.(*Param)
		if !ok2 {
			t.Error("解析成param出错")
		} else {
			t.Logf("this is param element in list %v", param.Z)
		}
	}
}

func TestConnectionPathToJsonWary0(t *testing.T) {
	//var p0, p1, p2, p3 *int32
	//var n0, n1, n2, n3 int32 = 0, 1, 2, 3
	//p0, p1, p2, p3 = &n0, &n1, &n2, &n3
	pA := &msggame.Point{
		XLocation: 0,
		YLocation: 1,
	}
	pB := &msggame.Point{
		XLocation: 1,
		YLocation: 2,
	}
	pC := &msggame.Point{
		XLocation: 1,
		YLocation: 3,
	}
	pD := &msggame.Point{
		XLocation: 2,
		YLocation: 3,
	}
	pPath := []*msggame.Point{pA, pB, pC, pD}
	conPath := &msggame.ConnectionPath{
		Pair: &msggame.TerminalPair{
			PointA: &msggame.Point{
				XLocation: 1,
				YLocation: 0,
			},
			PointB: &msggame.Point{
				XLocation: 3,
				YLocation: 3,
			},
		},
		Path: pPath,
	}
	jsonStr, err := json.Marshal(conPath)
	if err != nil {
		t.Error("解析成json出错")
	}
	t.Logf("the json of connectionPath is : %v", string(jsonStr))
}

func TestConnectionPathToJson(t *testing.T) {
	var p0, p1, p2, p3 *int32
	var n0, n1, n2, n3 int32 = 0, 1, 2, 3
	p0, p1, p2, p3 = &n0, &n1, &n2, &n3
	pA := &msggameNoZero.Point{
		XLocation: p0,
		YLocation: p1,
	}
	pB := &msggameNoZero.Point{
		XLocation: p1,
		YLocation: p2,
	}
	pC := &msggameNoZero.Point{
		XLocation: p1,
		YLocation: p3,
	}
	pD := &msggameNoZero.Point{
		XLocation: p2,
		YLocation: p3,
	}
	pPath := []*msggameNoZero.Point{pA, pB, pC, pD}
	conPath := &msggameNoZero.ConnectionPath{
		Pair: &msggameNoZero.TerminalPair{
			PointA: &msggameNoZero.Point{
				XLocation: p1,
				YLocation: p0,
			},
			PointB: &msggameNoZero.Point{
				XLocation: p3,
				YLocation: p3,
			},
		},
		Path: pPath,
	}
	jsonStr, err := json.Marshal(conPath)
	if err != nil {
		t.Error("解析成json出错")
	}
	t.Logf("the json of connectionPath is : %v", string(jsonStr))
}

// 应该注意proto中使用数据类型之后 在转化成json数据格式时 0值会被忽略的问题
func TestJsonToConnectionPath(t *testing.T) {
	jsonConnection := "{\"Pair\":{\"PointA\":{\"XLocation\":1,\"YLocation\":0},\"PointB\":{\"XLocation\":3,\"YLocation\":3}},\"Path\":[{\"XLocation\":0,\"YLocation\":1},{\"XLocation\":1,\"YLocation\":2},{\"XLocation\":1,\"YLocation\":3},{\"XLocation\":2,\"YLocation\":3}]}"
	connectionPath := &msggameNoZero.ConnectionPath{}
	err := json.Unmarshal([]byte(jsonConnection), connectionPath)
	if err != nil {
		t.Fatalf("json解析错误：%v", err.Error())
	}
	//if connectionPath.Pair.PointA.YLocation == 0 {
	//	t.Logf("connectionPath got 0 for its pair pointA locationY")
	//}
	if connectionPath.Pair.PointA.YLocation == nil {
		t.Logf("connectionPath got nil for its pair pointA locationY")
	}
	t.Logf("after unmarshal value is : %+v", connectionPath)
}

// 测试数组的json解析，数组容量用关键字计数
func TestJsonArrayToConnectionPathType(t *testing.T) {
	jsonString := "[{\"Pair\":{\"PointA\":{\"XLocation\":1,\"YLocation\":0},\"PointB\":{\"XLocation\":3,\"YLocation\":3}},\"Path\":[{\"XLocation\":0,\"YLocation\":1},{\"XLocation\":1,\"YLocation\":2},{\"XLocation\":1,\"YLocation\":3},{\"XLocation\":2,\"YLocation\":3}]}]"
	paths := make([]*msggame.ConnectionPath, 0, strings.Count(jsonString, "Pair"))
	_ = json.Unmarshal([]byte(jsonString), &paths)
	t.Logf("after unmarshal value is : %+v", paths)
}

// 测试用链表装填对象，保证随机的对象会因为链表中存储的关键字而不同上一个相同
func TestPCBRandWithAlreadyFrontElement(t *testing.T) {
	l := list.New()
	num := 2
	result := make(map[int]*PCB, 5)
	//frontValue := &PCB{}
	//frontValue.BeforeInit(num)
	//frontValue.Init()
	//l.PushFront(frontValue.Mark)
	for i := 0; i < 20; i++ {
		backElement := l.Back()
		p := &PCB{}
		if backElement == nil {
			p.BeforeInit(num)
			l.PushFront(p.Mark)
		} else {
		again:
			p.BeforeInit(num)
			if backElement.Value.(string) == p.Mark {
				t.Logf("the mark in pcb is : %v", p.Mark)
				goto again
			}
			l.InsertAfter(p.Mark, backElement)
		}
		p.Init()
		result[i] = p
	}
	for e := l.Front(); e != nil; e = e.Next() {
		t.Logf("the mark in list is : %v", e.Value.(string))
	}
	for _, v := range result {
		t.Logf("the pcb in result is : %+v", v)
	}
}

type PCBTest struct {
	Cells        []Cell
	RotationType msggame.PCBRotation
	FlipType     msggame.PCBFlip
}

func TestGameMsgEnum(t *testing.T) {
	rand.Seed(time.Now().UnixNano() / int64(time.Millisecond))
	t.Logf(fmt.Sprint(rand.Intn(12)))

	randRotation := rand.Intn(4) //[0,4)
	flipRotation := rand.Intn(3) //[0,3)

	p := &PCBTest{}
	p.RotationType = msggame.PCBRotation(randRotation)
	p.FlipType = msggame.PCBFlip(flipRotation)
	p.Cells = []Cell{{1, 2, 3}, {2, 4, 6}}

	t.Logf("here got the PCBTest is : %+v", p)
}

// 测试proto结构体中显式字段 即0值也必须展示 （常量取指针操作）
func TestConnectionPathToJsonZeroPrt(t *testing.T) {
	//var p0, p1, p2, p3 *int32
	//var n0, n1, n2, n3 int32 = 0, 1, 2, 3
	//p0, p1, p2, p3 = &n0, &n1, &n2, &n3
	//pA := &msggame.Point{
	//	XLocation: p0,
	//	YLocation: p1,
	//}
	//pB := &msggame.Point{
	//	XLocation: p1,
	//	YLocation: p2,
	//}
	//pC := &msggame.Point{
	//	XLocation: p1,
	//	YLocation: p3,
	//}
	//pD := &msggame.Point{
	//	XLocation: p2,
	//	YLocation: p3,
	//}
	//pPath := []*msggame.Point{pA, pB, pC, pD}
	//conPath := &msggame.ConnectionPath{
	//	Pair: &msggame.TerminalPair{
	//		PointA: &msggame.Point{
	//			XLocation: p1,
	//			YLocation: p0,
	//		},
	//		PointB: &msggame.Point{
	//			XLocation: p3,
	//			YLocation: p3,
	//		},
	//	},
	//	Path: pPath,
	//}
	//jsonStr, err := json.Marshal(conPath)
	//if err != nil {
	//	t.Error("解析成json出错")
	//}
	//t.Logf("the json of connectionPath is : %v", string(jsonStr))
}
