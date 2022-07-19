package main

import "log"

type Service struct {
	queue map[string]*Queue
}

// берем по нашему значению, проверяем ok или нет, если не ok - удаляем. Далее будем потокозащищенную мапку использовать.
// храним самый старый узел, если next == nil-закрываем, если нет, то он становится главным node

func (s *Service) Add(name string, value string) {
	node, ok := s.queue[name]
	// если очередь пустая - просто добавляем
	if !ok {
		s.queue[name] = &Queue{
			First: &Node{
				Current: value,
				Next:    nil,
			},
			Last: nil,
		}
		log.Println(s.queue[name].First, s.queue[name].Last)
		return
		// если не пустая - добавляем к последнему
	}
	if ok {
		// если всего один элемент
		last := &Node{
			Current: value,
			Next:    nil,
		}
		if node.First.Next == nil {
			s.queue[name] = &Queue{
				First: &Node{
					Current: node.First.Current,
					Next:    last,
				},
				Last: last,
			}
			log.Println(s.queue[name].First, s.queue[name].First.Next, s.queue[name].Last)
			return
		}
		// если элементов больше
		next := &Node{
			Current: value,
			Next:    nil,
		}
		// теперь последний указывает на новый элемент
		s.queue[name].Last.Next = next
		s.queue[name] = &Queue{
			First: node.First,
			Last:  next,
		}

	}
	log.Println(s.queue[name].First, s.queue[name].First.Next, s.queue[name].Last)
}
