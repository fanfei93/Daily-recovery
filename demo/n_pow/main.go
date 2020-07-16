package n_pow

import (
	"log"
	"time"
)

func GetNPow(a int, n int) int {
	if n == 1 {
		return a
	}
	if n % 2 == 0 {
		tmp := GetNPow(a, n / 2)
		return tmp * tmp
	} else {
		tmp := GetNPow(a, (n - 1) / 2)
		return  tmp * tmp * a
	}
}

func GetNPow2(a int, n int) int {
	var res = a
	for i := 2; i <= n; i++ {
		res = res * a
	}
	return res
}

func main() {
	start := time.Now().Nanosecond()
	var res1 int
	for i := 1; i < 10000; i++ {
		res1 = GetNPow(2,60)
	}
	end := time.Now().Nanosecond()
	log.Println("run time1:", end-start, "res:", res1)
	time.Sleep(time.Second)
	start2 := time.Now().Nanosecond()
	var res2 int
	for i := 1; i < 10000; i++ {
		res2 = GetNPow2(2,60)
	}
	end2 := time.Now().Nanosecond()
	log.Println("run time2:", end2 - start2, "res:", res2)
}
