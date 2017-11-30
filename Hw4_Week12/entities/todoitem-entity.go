package entities

//TodoItem 用户账户类型
type TodoItem struct {
	TID      int
	UserName string
	Todo     string
}

/*
// TodoItem .
func NewTodoItem(u TodoItem) *TodoItem {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	return &u
}
*/
