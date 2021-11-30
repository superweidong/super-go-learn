package main

type WeekDay int
//go 枚举
//定义const  iota从0开始 递增赋值
const (
	ONE  WeekDay = iota
	TWO
	THREE
)

func main() {
	println(ONE, TWO, THREE)
	println(switchEnumStr(TWO))

}

//类枚举转换
func switchEnumStr(enum WeekDay) string {
	switch enum {
	case ONE:
		return "ONE"
	case TWO:
		return "TWO"
	case THREE:
		return "THREE"
	default:
		return ""
	}
}
