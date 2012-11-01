package main

import (
	"fmt"
	"github.com/vron/lua/core"
)

var code = `
io.write("sdfsadf\n");
x = 0
for i = 1, #foo do
	print(i, ,foo[i])
	x = x + foo[i]
end
io.write("kjkjkjkjkjjk\n")
return x
`

func main() {
	core.Test()
	a := core.Newstate()
	core.Openlibs(a)

	status := core.Loadstring(a, code)
	if status != 0 {
		fmt.Println("Error: ", core.Tostring(a,-1))
		return
	}
	core.Createtable(a,10,10)

	for i:=1.0; i<=5; i++ {
		core.Pushnumber(a,i)
		core.Pushnumber(a,i*2)
		core.Rawset(a,-3)
	}

	core.Setglobal(a, "foo")

	result := core.Pcall(a,0, core.MULTRET, 0)
	if (result != 0) {
		fmt.Println("Error 2")
		return
	}

	sum := core.Tonumber(a,-1)

	fmt.Println(sum)

	core.Close(a)

	fmt.Println(a)
}
