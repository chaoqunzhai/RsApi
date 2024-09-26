/*
*
@Author: chaoqun
* @Date: 2023/5/28 23:46
*/
package utils

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"go-admin/global"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func DirNotCreate(dir string) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 目录不存在，创建它
		err = os.Mkdir(dir, 0755) // 0755是权限设置，你可以根据需要修改
		if err != nil {
			fmt.Printf("创建目录 %s 失败: %s\n", dir, err)
			return
		}
		fmt.Printf("目录 %s 已创建\n", dir)
	} else if err != nil {
		// 出现了其他错误
		fmt.Printf("获取目录 %s 信息时出错: %s\n", dir, err)
		return
	} else {
		// 目录已经存在
		//fmt.Printf("目录 %s 已存在\n", dir)
	}
	return

}

// GetWeekdayTimestamps 获取指定星期几的开始和结束时间戳
func GetWeekdayTimestamps(weekdayNumber int) (weekTime time.Time, err error) {
	// 获取当前时间

	now := time.Now()
	var weekday time.Weekday
	switch weekdayNumber {
	case 0:
		weekday = time.Sunday

		return now.AddDate(0, 0, int(now.Weekday())+1), nil
	case 1:
		weekday = time.Monday
	case 2:
		weekday = time.Tuesday
	case 3:
		weekday = time.Wednesday
	case 4:
		weekday = time.Thursday
	case 5:
		weekday = time.Friday
	case 6:
		weekday = time.Saturday

	default:

		return time.Time{}, errors.New("非法日期")
	}
	// 获取当前周的第一天的时间
	startOfWeek := now.AddDate(0, 0, int(-now.Weekday()+weekday))

	return startOfWeek, nil
}
func StructToMap(obj interface{}) map[string]interface{} {

	value := reflect.ValueOf(obj)
	structType := value.Type()
	structMap := make(map[string]interface{})

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		valueField := value.Field(i)
		jsonTag := field.Tag.Get("json")
		key := ""
		if jsonTag != "" {
			key = jsonTag
		} else {
			key = field.Name
		}
		structMap[key] = valueField.Interface()
	}

	return structMap
}

// 对数值进行补0 或者 带有小数点的数字 只保留2位
func StringDecimal(value interface{}) string {
	amount, err := decimal.NewFromString(fmt.Sprintf("%v", value))
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	return amount.StringFixed(2)

}
func StringToInt(v interface{}) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%v", v))
	return n
}
func StringToFloat64(v interface{}) float64 {
	n, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
	return n
}

func RoundDecimalFlot64(value interface{}) float64 {
	toStr := fmt.Sprintf("%v", value)
	amount3, _ := decimal.NewFromString(toStr)
	f, _ := amount3.Round(2).Float64()

	return f
}
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
func MinAndMax(values []float64) (float64, float64) {
	min1 := values[0] //assign the first element equal to min
	max1 := values[0] //assign the first element equal to max
	for _, number := range values {
		if number < min1 {
			min1 = number
		}
		if number > max1 {
			max1 = number
		}
	}
	return min1, max1
}

// 获取当前周几
func HasWeekNumber() int {
	n := time.Now()
	week := 0
	switch n.Weekday().String() {
	case "Sunday":
		week = 0
	case "Monday":
		week = 1
	case "Tuesday":
		week = 2
	case "Wednesday":
		week = 3
	case "Thursday":
		week = 4
	case "Friday":
		week = 5
	case "Saturday":
		week = 6
	}
	return week
}
func IsArray(key string, array []string) bool {
	set := make(map[string]struct{})
	for _, value := range array {
		set[value] = struct{}{}
	}
	_, ok := set[key]
	return ok
}
func IsArrayInt(key int, array []int) bool {
	set := make(map[int]struct{})
	for _, value := range array {
		set[value] = struct{}{}
	}
	_, ok := set[key]
	return ok
}

// 判断当前时间 是否在开始和结束时间区间
// TimeCheckRange("09:00","16:00")
func TimeCheckRange(start, end string) bool {
	now := time.Now()
	yearMD := now.Format("2006-01-02")
	//转换开始时间
	dbStartStr := fmt.Sprintf("%v %v", yearMD, start)
	dbStartTimer, _ := time.Parse("2006-01-02 15:04", dbStartStr)

	//转换结束时间
	dbEndStr := fmt.Sprintf("%v %v", yearMD, end)
	dbEndTimer, _ := time.Parse("2006-01-02 15:04", dbEndStr)

	//转换当前时间
	nowParse, _ := time.Parse("2006-01-02 15:04", time.Now().Format("2006-01-02 15:04"))

	return dbStartTimer.Before(nowParse) && nowParse.Before(dbEndTimer)
}
func ParseStrTime(timeStr string) (t time.Time, err error) {

	layout := fmt.Sprintf("2006-01-02 15:04")
	return time.ParseInLocation(layout, timeStr, global.LOC)
}
func ParInt(n float64) float64 {

	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	if err != nil {
		return n
	}
	return value
}

// 数组去重
func RemoveRepeatStr(list []string) (result []string) {
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]bool)
	for _, v := range list {
		// 遍历数组元素，判断此元素是否已经存在map中
		_, ok := temp[v]
		if !ok {
			temp[v] = true
			result = append(result, v)
		}
	}
	return result
}

// 数值去重
func RemoveRepeatInt(list []int) (result []int) {
	// 创建一个临时map用来存储数组元素
	temp := make(map[int]bool)
	for _, v := range list {
		// 遍历数组元素，判断此元素是否已经存在map中
		_, ok := temp[v]
		if !ok {
			temp[v] = true
			result = append(result, v)
		}
	}
	return result
}

func Avg(a []float64) float64 {
	sum := 0.0

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return ParInt(sum / float64(len(a)))
}
func Min(a []float64) float64 {
	minV := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			continue
		}
		if a[i] < minV {
			minV = a[i]
		}
	}
	return ParInt(minV)
}
func Max(a []float64) float64 {
	maxV := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxV {
			maxV = a[i]
		}
	}
	return ParInt(maxV)
}
func Percentile(N []float64, P float64) float64 {
	i := int(P * float64(len(N)))

	return N[i-1]
}

func RoundDecimal(value interface{}) decimal.Decimal {
	toStr := fmt.Sprintf("%v", value)
	amount3, _ := decimal.NewFromString(toStr) //0.8
	return amount3.Round(2)                    //0.80
}

// 处理精度问题
func DecimalMul(n int, k float32) float32 {
	a := decimal.NewFromFloat32(k)
	b := a.Mul(decimal.NewFromInt(int64(n)))

	c, _ := b.Float64()
	return float32(c)

}
func DecimalAdd(n1, n2 float32) float32 {
	a := decimal.NewFromFloat32(n1)
	b := a.Add(decimal.NewFromFloat32(n2))

	c, _ := b.Float64()
	return float32(c)

}

func ReplacePhone(phone string) (err error, phoneText string) {
	//str := `13734351278`ReplacePhone
	pattern := `^(\d{3})(\d{4})(\d{4})$`
	re := regexp.MustCompile(pattern) //确保正则表达式的正确 遇到错误会直接panic
	match := re.MatchString(phone)
	if !match {
		fmt.Println("phone number error")

		return errors.New("非法手机号"), ""
	}
	repStr := re.ReplaceAllString(phone, "$1****$3")

	return nil, repStr
}
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)

	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func CreateCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)) //这里面前面的04v是和后面的1000相对应的
}

// 求并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}
func IntersectInt(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}
func CheckStringSize(str string) bool {
	sizeInBytes := len(str)
	maxSizeInBytes := 3 * 1024 * 1024 // 3MB in bytes

	if sizeInBytes > maxSizeInBytes {
		fmt.Println("Error: String size exceeds 3MB.")
		return false
	}
	return true
}

// 求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 求差集 slice1-并集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
func DifferenceInt(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := IntersectInt(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// 随机字符串
func GetRandStr(n int) string {
	rand.Seed(time.Now().UnixNano())

	// 生成4位随机字母
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func RemoveDirectory(dir string) error {
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			err := os.RemoveAll(path) // 递归删除目录及其内容
			fmt.Println("删除path", path)
			if err != nil {
				return err
			}
		} else {
			err := os.Remove(path) // 删除文件
			fmt.Println("删除文件", path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return os.Remove(dir) // 删除空目录本身
}
func GetLayoutUnix(day, layout string) (start int64, end int64) {
	date, err := time.Parse(layout, day)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// 计算这一天的开始时间戳（0点）
	thisNow := time.Date(2024, date.Month(), date.Day(), 0, 0, 0, 0, global.LOC)
	startOfDayTimestamp := thisNow.Unix()
	fmt.Printf("Start of Day (2024-08-09 00:00:00) Timestamp: %d\n", startOfDayTimestamp)

	// 计算这一天的结束时间戳（23:59:59）
	// 通过将日期设置为下一天的0点，然后减去1秒来实现
	endOfDay := thisNow.Add(24 * time.Hour).Add(-time.Second)
	endOfDayTimestamp := endOfDay.Unix()
	fmt.Printf("End of Day (2024-08-09 23:59:59) Timestamp: %d\n", endOfDayTimestamp)

	return startOfDayTimestamp, endOfDayTimestamp
}

// Contains 检查slice中是否包含某个元素
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// FindDifferences 找出两个slice之间的差异
func FindDifferences(A, B []int) (added []int, removed []int) {
	for _, elem := range B {
		if !Contains(A, elem) {
			added = append(added, elem)
		}
	}
	for _, elem := range A {
		if !Contains(B, elem) {
			removed = append(removed, elem)
		}
	}
	return
}
