package main
import (
	"fmt"
	"os"
	"errors"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int 
	array [5]int
	head int // 指向队列队首
	tail int // 指向队尾
}

// 如队列 AddQueue(push) GetQueue(pop)
// 入队列
func (this *CircleQueue) push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	// this.tail 在队列的尾部，但是不包含最后的元素
	this.array[this.tail] = val // 把值放到后面
	this.tail = (this.tail + 1) % this.maxSize // 当前指向的位置没有值
	return 
}

// 出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	// 取数据, head是指向队首的，并且含有队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

// 显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列显示如下：")
	// 取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}
	// 声明一个辅助变量， 指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d] = %d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 判断环形队列为空
func (this *CircleQueue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head
}

// 判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	// 这是个主要的算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	// 先初始化一个队列, 并给其赋值
	queue := &CircleQueue{
		maxSize : 5,
		head : 0,
		tail : 0,
	}

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
				err := queue.push(val)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("加入队列成功")
				}
			case "get": 
				val, err := queue.Pop()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("取出的值为：", val)
				}
			case "show": 
				queue.ListQueue()
			case "exit" :
				os.Exit(0)
		}
	}
}