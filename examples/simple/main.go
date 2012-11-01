package main

import (
	"fmt"
	"github.com/vron/lua/lua"
)

var code = `
io.write("Hi from Lua\n");
x = 0
for i = 1, #foo do
	print(i, ,foo[i])
	x = x + foo[i]
end
io.write("By from Lua\n")
return x
`

func main() {
	a := lua.Newstate()
	lua.Openlibs(a)
	status := lua.Loadstring(a, code)
	if status != 0 {
		fmt.Println("Error: ", lua.Tostring(a,-1))
		return
	}
	lua.Createtable(a,10,10)
	for i:=1.0; i<=5; i++ {
		lua.Pushnumber(a,i)
		lua.Pushnumber(a,i*2)
		lua.Rawset(a,-3)
	}

	lua.Setglobal(a, "foo")

	result := lua.Pcall(a,0, lua.MULTRET, 0)
	if (result != 0) {
		fmt.Println("Error 2")
		return
	}

	sum := lua.Tonumber(a,-1)

	fmt.Println(sum)

	lua.Close(a)

	fmt.Println(a)
}
