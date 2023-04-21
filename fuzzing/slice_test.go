package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

type sliceTest struct {
	Id       int
	IntSlice []int
	Name     string
}

type MapEntity struct {
	FoodNum     int
	Probability int
}

func (m *MapEntity) Str2Init(str string) {
	// str -> FoodNum=xx,Probability=yy
	fieldsStr := strings.Split(str, ",")
	theMap := make(map[string]int, 2)
	for _, s := range fieldsStr {
		// fieldsStr[0] -> FoodNum=xx
		// fieldsStr[1] -> Probability=yy
		sList := strings.Split(s, "=")
		value, _ := strconv.Atoi(sList[1])
		theMap[sList[0]] = value
	}
	_ = mapstructure.Decode(theMap, m)
}

func TestAppendIntSlice(t *testing.T) {
	//mgr := GetExcelUtilMgr()
	//mgr.readExcelFile(excel.EXCEL_RHYTHM_MASTER_CONFIG)
	intSlice1 := make([]int, 5)
	var intSlice2 []int
	data := []int{1, 2, 3, 4, 5}

	for _, datum := range data {
		intSlice1 = append(intSlice1, datum)
		intSlice2 = append(intSlice2, datum)
	}

	fmt.Printf("the intSlict1 is : %v \n", intSlice1)
	fmt.Printf("the intSlict2 is : %v \n", intSlice2)
}

func TestReflectAppendSlice(t *testing.T) {
	dataIntSlice := []string{"1;3;5;7;9;11", "2;4;6;8;10"}
	ids := []int{1, 2}
	typeTCD := reflect.TypeOf(sliceTest{})                        // 获取struct的类型
	targetMapType := reflect.MapOf(reflect.TypeOf(1), typeTCD)    // 定义map构成类型  key是什么类型 value是什么类型
	targetMap := reflect.MakeMapWithSize(targetMapType, len(ids)) // 生成指定的map
	for index, id := range ids {
		singleTarget := reflect.New(typeTCD) // 指向该新对象的指针
		for i := 0; i < typeTCD.NumField(); i++ {
			fmt.Printf("the field.name is : %v \n", typeTCD.Field(i).Name)
			switch typeTCD.Field(i).Name {
			case "Id":
				reflect.Indirect(singleTarget).FieldByName("Id").SetInt(int64(id)) // 指针取对象向对象内的字段进行赋值操作
				break
			case "Name":
				reflect.Indirect(singleTarget).FieldByName("Name").SetString("testING")
				break
			case "IntSlice":
				stringSlice := strings.Split(dataIntSlice[index], ";")
				var intSlice []int // 仅声明slice 即空slice 若指定长度则会初始化各个元素为0
				for _, v := range stringSlice {
					vInt, e := strconv.Atoi(v)
					if e != nil {
						fmt.Printf("数组解析有误：%s\n", e)
					}
					intSlice = append(intSlice, vInt)
				}
				filedSlice := reflect.Indirect(singleTarget).FieldByName("IntSlice")   // 获取到对象内的原slice字段
				newSlice := reflect.AppendSlice(filedSlice, reflect.ValueOf(intSlice)) // 将slice 拼接slice 生成新的slice
				reflect.Indirect(singleTarget).FieldByName("IntSlice").Set(newSlice)
				break
			}
		}
		//targetMap.SetMapIndex(reflect.ValueOf(index), singleTarget)
		targetMap.SetMapIndex(reflect.ValueOf(index), reflect.Indirect(singleTarget)) // 将对象放置到map中 注意不是将指针放入（结合上文定义的map类型）
	}
	fmt.Printf("the result is %+v \n", targetMap)

}

func TestParseData2Map(t *testing.T) {
	dataIntSlice := []string{"1;3;5;7;9;11", "2;4;6;8;10"}
	ids := []int{1, 2}
	targetMap := make(map[int]*sliceTest)
	ParseData2Map(dataIntSlice, ids, targetMap)
	//fmt.Printf("the result is %+v \n", targetMap)
	for i, result := range targetMap {
		fmt.Printf("the key is : %v , and the value is : %+v; \n", i, result)
	}
}

func TestFloat32ToStr(t *testing.T) {
	a := float32(3.79)
	t.Logf("the duration result is %v", fmt.Sprint(a, "s"))
	aDuration, _ := time.ParseDuration(fmt.Sprint(a, "s"))
	t.Logf("the a result in millisecond is %v", aDuration.Milliseconds())
}

func TestStr2Init(t *testing.T) {
	tarString := "FoodNum=9,Probability=1"
	tarConfig := &MapEntity{}
	tarConfig.Str2Init(tarString)
	t.Logf("the config result is : %+v", tarConfig)
}

func TestWeightedChoiceMap(t *testing.T) {
	m := make(map[int]int, 3)
	m[7] = 1
	m[8] = 2
	m[9] = 3

	m2 := make(map[int]int, 4)
	m2[1] = 7
	m2[2] = 9
	m2[3] = 8
	m2[4] = 10

	for i := 0; i < 10; i++ {
		t.Logf("the result of m is : %v", WeightedChoiceMap(m))
	}
	for i := 0; i < 10; i++ {
		t.Logf("the result of m2 is : %v", WeightedChoiceMap(m2))
	}

}
