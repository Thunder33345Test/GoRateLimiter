package RateLimiter

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)
//Lol more like testing then actual test
func init() {
	gl := NewLimiterGroupAnd(NewLimiter(time.Second, 3), NewLimiterGroupOr(NewLimiter(time.Second, 2), NewLimiter(time.Second*5, 4)))
	for i := 0; i < 5; i++ {
		fmt.Printf("rem:%d,%t\n", gl.RemainingTally(), gl.TryAddTally())
	}
	sgl, _ := gl.(Storable)
	gsd := sgl.GetStorableData()
	fmt.Printf("raw data: %v\n", gsd)

	jsonDat, _ := json.MarshalIndent(gsd, "", " ")

	fmt.Printf("json: %s\n", string(jsonDat))

	newI := &[]interface{}{}
	_ = json.Unmarshal(jsonDat, newI)
	fmt.Printf("NewI: %s\n", strings.Join(strings.Split(fmt.Sprintf("%#v", newI), "},"), "},\n"))
	ncl := NewLimiterGroupAnd(NewLimiter(time.Second, 5), NewLimiterGroupOr(NewLimiter(time.Second, 2), NewLimiter(time.Second, 3)))
	rs := ncl.(Storable)

	res := rs.SetStorableData(*newI)
	fmt.Printf("result: %t\n", res)
	for i := 0; i < 4; i++ {
		fmt.Printf("rem:%d,%t\n", gl.RemainingTally(), gl.TryAddTally())
	}
	db := ncl.(Debuggable)
	fmt.Printf("\n%#v\n", db.GetDebug())
	jsonDat, err := json.MarshalIndent(db.GetDebug(), "", " ")
	if err != nil {
		fmt.Printf("\n%s\n", err)
	}
	fmt.Printf("\n%s\n", jsonDat)
}

func TestAdd(t *testing.T) {
	t2 := 2 * time.Second
	unit := 5
	l := NewLimiter(t2, unit)
	c := 0
	for i := 0; i < 10; i++ {
		if l.TryAddTally() && c > unit {
			t.Fail()
		}
		c++
	}
}

func TestGet(t *testing.T) {

}
