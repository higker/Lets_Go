package lets_36_testing

import (
	"fmt"
	"testing"
)

func TestSqlitStr(t *testing.T) {
	sqlit := Sqlit("222222", "2‍")
	fmt.Println(sqlit)
}
