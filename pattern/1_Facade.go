/*
Паттерн Facade относится к структурным паттернам уровня объекта.

Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных классов
или объектов некоторой подсистемы, что облегчает ее использование.

Плюсы:
- Изоляция клиентов от поведения сложной системы

Минусы:
- Фасад рискует стать суперклассом, привязанным ко всем классам программы.
*/

package main

import "strings"

type Swan struct {
}

func (s *Swan) Fly() string {
	return "flying"
}

type Cancer struct {
}

func (c *Cancer) Crawl() string {
	return "crawling"
}

type Pike struct {
}

func (p *Pike) Swim() string {
	return "swimming"
}

func NewAnimals() *Animals {
	return &Animals{
		swan:   &Swan{},
		canser: &Cancer{},
		pike:   &Pike{},
	}
}

type Animals struct {
	swan   *Swan
	canser *Cancer
	pike   *Pike
}

func (a *Animals) Action() string {
	result := []string{
		a.pike.Swim(),
		a.swan.Fly(),
		a.canser.Crawl(),
	}
	return strings.Join(result, " ")
}
