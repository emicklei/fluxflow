package nodes

import (
	"fmt"
	"reflect"
)

var stdpkg = map[string]map[string]reflect.Value{}

func init() {
	stdpkg["fmt"]["Println"] = reflect.ValueOf(fmt.Println)
}
