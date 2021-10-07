package genshin

import (
	"fmt"
	"math"
	"testing"
)

func TestDs(t *testing.T) {
	//1633416571,3e72ln,abc4a6b2658d79c4909d675399d85f63
	// version := "2.13.1"
	var salt = genSalt() //6zT9berkIjLBimVKLeQiyYCN0tatGDpP
	const time = "1633416571"
	const random = "3e72ln"

	md5 := md5Str(fmt.Sprintf("salt=%s&t=%s&r=%s", salt, time, random))
	if md5 != "abc4a6b2658d79c4909d675399d85f63" {
		t.Log(md5)
		t.Fail()
	}
}

func genSalt() string {
	var arrayOfInt = make([]int, 32)
	arrayOfInt[0] = -531441
	arrayOfInt[1] = 222
	arrayOfInt[2] = 108
	arrayOfInt[3] = -14348907
	arrayOfInt[4] = 150
	arrayOfInt[5] = -102
	arrayOfInt[6] = 198
	arrayOfInt[7] = -108
	arrayOfInt[8] = -74
	arrayOfInt[9] = 174
	arrayOfInt[10] = 84
	arrayOfInt[11] = 54
	arrayOfInt[12] = -106
	arrayOfInt[13] = -110
	arrayOfInt[14] = 114
	arrayOfInt[15] = -76
	arrayOfInt[16] = 84
	arrayOfInt[17] = -102
	arrayOfInt[18] = -82
	arrayOfInt[19] = -106
	arrayOfInt[20] = -122
	arrayOfInt[21] = -90
	arrayOfInt[22] = -68
	arrayOfInt[23] = 90
	arrayOfInt[24] = -729
	arrayOfInt[25] = 204
	arrayOfInt[26] = -98
	arrayOfInt[27] = 204
	arrayOfInt[28] = -72
	arrayOfInt[29] = 60
	arrayOfInt[30] = 192
	arrayOfInt[31] = 96

	var b = 0
	res := make([]byte, 0, 32)
	for b < 32 {
		j := arrayOfInt[b]
		if j < 0 {
			if -j >= 729 {
				tmp := math.Log(float64(-j)) / math.Log(3.0)
				j = int(tmp - 6 + 48)
			} else {
				j ^= 0xFFFFFFFF
			}
		} else {
			j = j/3 + 48
		}
		res = append(res, byte(j))
		b++
	}
	return string(res)
}

func TestGenSalt(t *testing.T) {
	res := genSalt()
	t.Log(res)
}
