package moon

import (
	"fmt"
	"testing"
	"time"
)

func Test_Rise(t *testing.T) {
	fmt.Println(RiseTime(time.Now(), 115, 32, true))
}
