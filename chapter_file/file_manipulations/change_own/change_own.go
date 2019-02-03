package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {

	file := "/home/etgo/test.txt"

	user, err := user.Lookup("newuser")
	if err != nil {
		panic(err)
	}

	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		panic(err)
	}

	gid, err := strconv.Atoi(user.Gid)
	if err != nil {
		panic(err)
	}
	// 修改文件所有者，参数是uid，用户ID，以及gid，组ID
	// 注意这个函数只有在UNIX系统才有实际作用，Windows下返回错误
	err = os.Chown(file, uid, gid)
	if err != nil {
		fmt.Println(err)
	}
}
