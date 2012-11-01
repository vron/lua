// Provides core functionality from Lua C Api 
package core

/*
#cgo linux pkg-config: luajit
#include <stdlib.h>
#include <lua.h>
#include <lauxlib.h>
*/
import "C"

import (
	"fmt"
)

func Test() {
	fmt.Println(int(C.random()))
}

func Do() {
	C.luaL_newstate()
}
