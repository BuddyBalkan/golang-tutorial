package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	// 对输入的字符串进行Unicode的检查
	if !utf8.ValidString(s) {
		return s, errors.New("the input value is not Unicode for not valide UTF-8")
	}
	// fmt.Printf("the input: %q \n", s)
	// b := []byte(s) // bug fixed 事实上有用不止一个byte来表示字符的unicode
	b := []rune(s) // bug go在将string转化成rune时 可能会将单字节但是非有效的unicode的string转化成unicode中代表未知字符的字符集 该操作就改变了双反转与原字符串相等的结果
	// fmt.Printf("the rune: %q \n", b)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}

func main() {
	input := "the quick brown fox jumped over the lazy dog"

	rev, revErr := Reverse(input)

	doubleRev, doubleRevErr := Reverse(rev)

	fmt.Printf("original: %q \n", input)
	fmt.Printf("reverse: %q, and error: %v\n", rev, revErr)
	fmt.Printf("reverse again: %q, and error: %v\n", doubleRev, doubleRevErr)
}

// 使用反射来解决构造指定map并填充entity （已有更好的解决） 测试方法：TestParseData2Map
func ParseData2Map(dataSlice []string, ids []int, targetMap interface{}) {
	dataMapValue := reflect.Indirect(reflect.ValueOf(targetMap)) // 这里变量命名是已知targetMap为map类型 但是该做法兼容了slice类型
	mapEntryType := reflect.TypeOf(targetMap).Elem().Elem()      // 获取map中entry的类型  （若只使用了一次Elem()方法 则获取的是该结构体的指针 则下文中的reflect.New()创建的是？）
	//fmt.Printf("the mapEntryType is : %+v \n", mapEntryType)
	//fmt.Printf("the mapEntryType type is : %T \n", mapEntryType)
	//fmt.Printf("the dataMapValue is : %+v \n", dataMapValue)
	//fmt.Printf("the dataMapValue type is : %T \n", dataMapValue)
	for index, id := range ids {
		singleTarget := reflect.New(mapEntryType) // 创建了指定结构体的新零值指针 （空对象）
		for i := 0; i < mapEntryType.NumField(); i++ {
			fmt.Printf("the field.name is : %v \n", mapEntryType.Field(i).Name)
			switch mapEntryType.Field(i).Name {
			case "Id":
				reflect.Indirect(singleTarget).FieldByName("Id").SetInt(int64(id)) // 指针取对象 向对象内的字段进行赋值操作
				break
			case "Name":
				reflect.Indirect(singleTarget).FieldByName("Name").SetString("testING")
				break
			case "IntSlice":
				stringSlice := strings.Split(dataSlice[index], ";")
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
		dataMapValueKind := reflect.TypeOf(dataMapValue.Interface()).Kind()
		if dataMapValueKind == reflect.Map {
			dataMapValue.SetMapIndex(reflect.ValueOf(index), singleTarget) // 将对象的指针放置到map中（结合上文定义的map类型）
		} else if dataMapValueKind == reflect.Slice { //做法兼容了slice类型的情况 但当前不会进入该判断中
			dataMapValue.Set(reflect.AppendSlice(dataMapValue, singleTarget))
		}
	}
}
