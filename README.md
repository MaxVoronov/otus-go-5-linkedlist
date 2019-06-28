# Otus Homework #5
## Двусвязный список
Ожидаемые типы (псевдокод):
```
List        // тип контейнер
    Len()   // длинна списка
    First() // первый Item
    Last()  // последний Item
    PushFront(v interface{}) // добавить значение в начало
    PushBack(v interface{})  // добавить значение в конец
    
Item        // элемент списка
    Value() interface{}  // возвращает значение
    Next() *Item         // следующий Item
    Prev() *Item         // предыдущий
    Remove()             // удалить Item из списка
```

### Пример использования

```go
package main

import (
   "fmt"
   "github.com/maxvoronov/otus-go-5-linkedlist"
)

func main() {
    list := doublylinkedlist.NewList()
    
    firstNode := list.PushFront(72)
    lastNode := list.PushBack(98)
    list.PushAfter(firstNode, 55)
    list.PushBefore(lastNode, 48)
    
    tmpNode := list.PushBefore(firstNode, 0)
    list.Remove(tmpNode)
    
    fmt.Printf("Doubly Linked List: %s\n", list)
}

// Output:
// Doubly Linked List: [H] -> [72] -> [55] -> [48] -> [98] -> [T]
```

### Полезные ссылки
- [Wiki: Doubly Linked List](https://en.wikipedia.org/wiki/Doubly_linked_list)
