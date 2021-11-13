package golang_offer_test

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"testing"
)

func TestGOGC(t *testing.T) {
	fmt.Println(os.Getenv("GOGC"))
	os.Setenv("GOGC", "200")

	v := debug.SetGCPercent(100)
	fmt.Println(v)
	defer debug.SetGCPercent(v)
	fmt.Println(os.Getenv("GOGC"))
	runtime.GC()
}
