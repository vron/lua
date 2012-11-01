// Provides core functionality from Lua C Api 
package lua

/*
#cgo linux pkg-config: luajit
#include <stdlib.h>
#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>
*/
import "C"

import (
	"unsafe"
)

const (
	GLOBALSINDEX = C.LUA_GLOBALSINDEX

	MULTRET = C.LUA_MULTRET
)

type State C.lua_State

func Newstate() *State {
	return (*State)(C.luaL_newstate())
}

func Openlibs(s *State) {
	C.luaL_openlibs((*C.lua_State)(s))
}

func Loadstring(s *State, str string) int {
	// Create a c string
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.luaL_loadstring((*C.lua_State)(s),cstr))
}

func Createtable(s *State, narr, nrec int) {
	C.lua_createtable((*C.lua_State)(s), C.int(narr), C.int(nrec))
}

func Pushnumber(s *State, num float64) {
	C.lua_pushnumber((*C.lua_State)(s), C.lua_Number(num))
}

func Rawset(s *State, index int) {
	C.lua_rawset((*C.lua_State)(s), C.int(index))
}

func Setglobal(s *State, name string) {
	// Defined as a macro
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.lua_setfield((*C.lua_State)(s), GLOBALSINDEX, cs)
}

func Pcall(s *State, nargs, nres, errf int) int {
	return int(C.lua_pcall((*C.lua_State)(s),
		C.int(nargs), C.int(nres), C.int(errf)))
}

func Tonumber(s *State, index int) float64 {
	return float64(
		C.lua_tonumber((*C.lua_State)(s),
			C.int(index)))
}

func Pop(s *State, n int) {
	// Is defined as a macro, so we define it:
	C.lua_settop((*C.lua_State)(s), C.int(-n-1))
}

func Close(s *State) {
	C.lua_close((*C.lua_State)(s))
}

func Tostring(s *State, index int) string {
	var l C.size_t
	cs := C.lua_tolstring((*C.lua_State)(s),
		C.int(index), &l)
	// Remake this into a string!
	return C.GoStringN(cs, C.int(l))
}
