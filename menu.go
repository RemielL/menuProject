package main
 
import "fmt"
 
func helpPrint(){
	fmt.Println("可使用以下命令：")
	fmt.Println("  help：帮助信息")
	fmt.Println("  list：列表信息")
	fmt.Println("  quit：退出")
}

func main() {
	var str string
	fmt.Println("欢迎使用menu，可输入“help”获取帮助")
	for {
		fmt.Print("请输入命令：")
		fmt.Scan(&str)
		if str == "quit" {
			break
		}
		switch str {
		case "help":
			helpPrint()
		case "list":
			fmt.Println("列表信息")
		default:
			fmt.Println("错误指令，请重新输入")
		}
	}
}
