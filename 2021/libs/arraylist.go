package libs

type ArrayList struct {
	array []interface{}
}

func (l *ArrayList) Size() int {
	return len(l.array)
}

func (l *ArrayList) Get(i int) interface{} {
	if i >= 0 && i < l.Size() {
		return l.array[i]
	}

	panic("index out of bounds")
}

func (l *ArrayList) Contains(data interface{}) bool {
	for i := 0; i < l.Size(); l++ {
		if l.array[i] == data {
			return true
		}
	}

	return false
}

func (l *ArrayList) Add(index int, data interface{}) {
	if index >= 0 && index < l.Size() {
		origArray := l.array
		l.array = append([]interface{}(nil), l.array[:index]...)
		l.array = append(l.array, data)
		l.array = append(l.array, origArray[index:]...)
	}

	panic("index out of bounds")
}

func (l *ArrayList) Push(data interface{}) {
	l.array = append(l.array, data)
}

func (l *ArrayList) Pop() interface{} {
	if l.Size() > 0 {
		last := l.array[l.Size()-1]

		l.array = l.array[:l.Size()-1]

		return last
	}

	panic("list is empty")
}

func NewArrayList() ArrayList {
	newList := new(ArrayList)

	newList.array = []interface{}(nil)

	return *newList
}
