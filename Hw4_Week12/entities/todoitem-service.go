package entities

//TodoItemAtomicService .
type TodoItemAtomicService struct{}

//TodoItemService .
var TodoItemService = TodoItemAtomicService{}

//RegistTodoItem .
func (*TodoItemAtomicService) RegistTodoItem(t *TodoItem) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := todoItemDao{tx}
	err = dao.AddtodoItem(t)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

//DeleteTodoItem .
func (*TodoItemAtomicService) DeleteTodoItem(t *TodoItem) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := todoItemDao{tx}
	err = dao.DeletetodoItem(t)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

//FindAll .
func (*TodoItemAtomicService) FindAll() []TodoItem {
	dao := todoItemDao{mydb}
	return dao.FindAlltodoItems()
}

//FindAllByUserName .
func (*TodoItemAtomicService) FindAllByUserName(username string) []TodoItem {
	dao := todoItemDao{mydb}
	return dao.FindAlltodoItemsByUserName(username)
}

//FindByID .
func (*TodoItemAtomicService) FindByID(id int) *TodoItem {
	dao := todoItemDao{mydb}
	return dao.FindtodoItemByID(id)
}

//FindByUserName .
func (*TodoItemAtomicService) FindByUserName(username string) *TodoItem {
	dao := todoItemDao{mydb}
	return dao.FindtodoItemByUserName(username)
}
