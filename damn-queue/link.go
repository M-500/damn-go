package damn_queue

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-21 22:11

type LinkNode struct {
	prev, next *LinkNode
	data       any // 这里我是存值 还是存指针？
}

func NewLinkNode() *LinkNode {
	return &LinkNode{}
}
