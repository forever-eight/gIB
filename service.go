package main

import "log"

type Service struct {
	// todo потокозащищенная map
	queue map[string]*Queue
}

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
		// теперь последний указывает на новый элемент
		s.queue[name].Last.Next = last
		s.queue[name] = &Queue{
			First: node.First,
			Last:  last,
		}

	}
	log.Println(s.queue[name].First, s.queue[name].First.Next, s.queue[name].Last)
}

func (s *Service) Get(name string) *string {
	node, ok := s.queue[name]
	if !ok {
		return nil
	} else {
		cur := node.First.Current
		// если один элемент
		if node.First.Next == nil {
			delete(s.queue, name)
			return &cur
		} else if node.First.Next == node.Last {
			// если элементов два
			s.queue[name] = &Queue{
				First: node.Last,
				Last:  nil,
			}
			return &cur
		} else {
			// если элементов больше
			next := node.First.Next
			s.queue[name] = &Queue{
				First: next,
				Last:  node.Last,
			}
		}
		return &cur
	}
}
