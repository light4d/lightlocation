package tests

import (
	"github.com/light4d/lightlocation/location"
	"testing"
)

func Test_location(T *testing.T) {
	//testMap := make(map[string]string)
	////fmt.Printf("map长度:%d\n", len(testMap))
	//testMap["ip"] = "116.30.218.255"
	//testMap["ak"] = "EcaqNxUoy0LHtsVKXIshPOqHZZAHN7sj"
	//
	//paramsStr := location.toQueryString(testMap)
	//
	//fmt.Println(paramsStr)
	//
	//fmt.Println("sn: " + location.createBaiduLbsSn(paramsStr))
	//
	//fmt.Println("url: " + location.createBaiduReqUrl("116.30.218.255"))
	//
	location.GetLocation(nil)
}
