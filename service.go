package main

import (
	"log"
	"sync"
	"time"
)

type Service struct {
	mu     *sync.Mutex
	queue  map[string]*Queue
	waiter map[string]int
	c      chan string
}

func (s *Service) Add(name string, value string) {
	s.mu.Lock()

	// waiter
	_, ok := s.waiter[name]
	if ok {
		s.c <- value
		s.waiter[name]--
		return
	}

	s.mu.Unlock()

	s.mu.Lock()
	node, ok := s.queue[name]
	s.mu.Unlock()
	// если очередь пустая - просто добавляем
	if !ok {
		s.mu.Lock()
		s.queue[name] = &Queue{
			First: &Node{
				Current: value,
				Next:    nil,
			},
			Last: nil,
		}
		s.mu.Unlock()
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
			s.mu.Lock()
			s.queue[name] = &Queue{
				First: &Node{
					Current: node.First.Current,
					Next:    last,
				},
				Last: last,
			}
			s.mu.Unlock()
			//log.Println(s.queue[name].First, s.queue[name].First.Next, s.queue[name].Last)
			return
		}

		// если элементов больше
		// теперь последний указывает на новый элемент
		s.mu.Lock()
		s.queue[name].Last.Next = last
		s.queue[name] = &Queue{
			First: node.First,
			Last:  last,
		}
		s.mu.Unlock()
	}
}

func (s *Service) Get(name string) *string {
	s.mu.Lock()

	node, ok := s.queue[name]
	s.mu.Unlock()
	if !ok {
		return nil
	} else {
		cur := node.First.Current
		// если один элемент

		if node.First.Next == nil {
			s.mu.Lock()
			delete(s.queue, name)
			s.mu.Unlock()
			return &cur
		} else if node.First.Next == node.Last {
			s.mu.Lock()
			// если элементов два
			s.queue[name] = &Queue{
				First: node.Last,
				Last:  nil,
			}
			s.mu.Unlock()
			return &cur
		} else {
			// если элементов больше
			s.mu.Lock()
			next := node.First.Next
			s.queue[name] = &Queue{
				First: next,
				Last:  node.Last,
			}
			s.mu.Unlock()
		}
		return &cur
	}
}

func (s *Service) Wait(n int, name string) *string {
	timeout := time.After(time.Duration(n) * time.Second)
	value := s.Get(name)
	if value != nil {
		return value
	}

	s.mu.Lock()
	s.waiter[name]++
	s.mu.Unlock()
	select {
	case value := <-s.c:
		return &value
	case <-timeout:
		log.Println("end", n)
		return nil
	}
}
