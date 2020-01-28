package main
import (
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int 
	array [4]int
	head int // 指向队列队首
	tail int // 指向队尾
}

// 如队列 AddQueue(push) GetQueue(pop)
func (this *CircleQueue) push(val int) (err error) {

}

func (this *CircleQueue) Pop() (val int, err error) {
	
}

func main() {
	// 先创建一个队列, 并给其赋值
	// queue := &CircleQueue{
	// 	maxSize : 5,
	// 	front : -1,
	// 	rear : -1,
	// }

	var key string 
	var val int 
	for {
		fmt.Println("1. 输入add 添加数据到队列")
		fmt.Println("2. 输入get 从队列中获取数据")
		fmt.Println("3. 输入show 获取队列的信息")
		fmt.Println("4. 输入exit 退出程序")

		fmt.Scanln(&key)
		switch key {
			case "add":
				fmt.Println("输入需要入队列的数据")
				fmt.Scanln(&val)
				err := queue.AddQueue(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("加入队列成功")
				}
			case "get": 
				val, err := queue.GetQueue()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("取出的值为：", val)
				}
			case "show": 
				queue.ShowQueue()
			case "exit" :
				os.Exit(0)
		}
	}
}