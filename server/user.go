package server

import (
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"strconv"
	"time"
)

type Gender hessian.JavaEnum

const (
	MAN hessian.JavaEnum = iota
	WOMAN
)

var genderName = map[hessian.JavaEnum]string{
	MAN:   "MAN",
	WOMAN: "WOMAN",
}

var genderValue = map[string]hessian.JavaEnum{
	"MAN":   MAN,
	"WOMAN": WOMAN,
}

func (g Gender) JavaClassName() string {
	return "com.ikurento.user.Gender"
}

func (g Gender) String() string {
	s, ok := genderName[hessian.JavaEnum(g)]
	if ok {
		return s
	}

	return strconv.Itoa(int(g))
}

func (g Gender) EnumValue(s string) hessian.JavaEnum {
	v, ok := genderValue[s]
	if ok {
		return v
	}

	return hessian.InvalidJavaEnum
}

type (
	User struct {
		// !!! Cannot define lowercase names of variable
		Id   string
		Name string
		Age  int32
		Time time.Time
		Sex  Gender // notice: java enum Object <--> go string
	}
)

var (
	DefaultUser = User{
		Id: "0", Name: "Alex Stocks", Age: 31,
		Sex: Gender(MAN),
	}

	userMap = make(map[string]User)
)

func init() {
	userMap["A000"] = DefaultUser
	userMap["A001"] = User{Id: "001", Name: "ZhangSheng", Age: 18, Sex: Gender(MAN)}
	userMap["A002"] = User{Id: "002", Name: "Lily", Age: 20, Sex: Gender(WOMAN)}
	userMap["A003"] = User{Id: "113", Name: "Moorse", Age: 30, Sex: Gender(WOMAN)}
	for k, v := range userMap {
		v.Time = time.Now()
		userMap[k] = v
	}
}

func (u User) String() string {
	return fmt.Sprintf(
		"User{Id:%s, Name:%s, Age:%d, Time:%s, Sex:%s}",
		u.Id, u.Name, u.Age, u.Time, u.Sex,
	)
}

func (u User) JavaClassName() string {
	return "com.ikurento.user.User"
}

func println(format string, args ...interface{}) {
	fmt.Printf("\033[32;40m"+format+"\033[0m\n", args...)
}