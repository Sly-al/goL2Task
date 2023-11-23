package main

/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и
стоит ли передавать запрос дальше по цепи.

Плюсы
- уменьшает зависимость между клиентом и обработчиками. Каждый обработчик
сам выполняет свою логику независимо.
- реализует принцип единственной обязанности.
- реализует принцип открытости и закрытости.

Минусы
- запрос может остаться необработанным.

Примеры использования:
- в разрабатываемой системе имеется группа объектов, которые могут обрабатывать
сообщения определенного типа;
- все сообщения должны быть обработаны хотя бы одним объектом системы;
- сообщения в системе обрабатываются по схеме «обработай сам либо перешли другому»,
то есть одни сообщения обрабатываются на том уровне, где они получены,
а другие пересылаются объектам иного уровня.
*/

import "fmt"

type Handler interface {
	Handle(i int) int
	SetNext(h Handler)
}

type Little struct {
	next Handler
}

func (l *Little) SetNext(h Handler) {
	l.next = h
}

func (l *Little) Handle(i int) int {
	if i < 100 {
		return 100 - i
	} else if l.next != nil {
		return l.next.Handle(i)
	}
	return -1
}

type Small struct {
	next Handler
}

func (s *Small) SetNext(h Handler) {
	s.next = h
}

func (s *Small) Handle(i int) int {
	if i < 1000 {
		return 1000 - i
	} else if s.next != nil {
		return s.next.Handle(i)
	}
	return -1
}

type Big struct {
	next Handler
}

func (b *Big) SetNext(h Handler) {
	b.next = h
}

func (b *Big) Handle(i int) int {
	if i < 10000 {
		return 10000 - i
	} else if b.next != nil {
		return b.next.Handle(i)
	}
	return -1
}

func main() {
	s := &Small{}
	l := &Little{}
	b := &Big{}
	s.SetNext(b)
	l.SetNext(s)
	for _, now := range []int{70, 700, 7000, 70000} {
		fmt.Println(l.Handle(now))
	}
}
