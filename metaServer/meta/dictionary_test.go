package meta

import "testing"

func TestDirectory_AddDir(t *testing.T) {
	root := NewDirectory("root")

	// 添加目录结构
	root.AddDir("home")
	root.AddDir("etc")
	root.AddDir("var")

	// 在 home 目录下添加用户目录
	home, _ := root.FindDir("home")
	home.AddDir("user")
	home.AddDir("admin")

	// 在 user 目录下添加文件
	user, _ := home.FindDir("user")
	user.AddFile("example.txt")

	// 打印整个文件系统结构
	root.Print("")
}
