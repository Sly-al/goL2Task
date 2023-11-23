/*
Шаблон «Строитель» предназначен для поиска решения проблемы антипаттерна Telescoping constructor.
Плюсы:
- Инкапсуляция процесса создания сложного объекта.
- Возможность поэтапного конструирования объекта с переменным набором этапов (в отличие	от «одноэтапных» фабрик).
- Сокрытие внутреннего представления продукта от клиента.
- Реализации продуктов могут свободно изменяться, потому что клиент имеет дело только с абстрактным интерфейсом

Минусы:
-ConcreteBuilder и создаваемый им продукт жестко связаны между собой, поэтому при внесеннии изменений в класс продукта скорее всего придется
	соотвествующим образом изменять и класс ConcreteBuilder.
Пример использования:
-В системе могут существовать сложные объекты, создание которых за одну
	операцию затруднительно или невозможно. Требуется поэтапное построение объектов
	с контролем результатов выполнения каждого этапа.
*/

package main

import "fmt"

type Director struct {
	b Builder
}

func NewDirector() *Director {
	return &Director{
		b: &ConcreteBuilder{
			a: new(Action),
		},
	}
}

func (d *Director) Build(id, cost int, task, dif string) *Action {
	d.b.SetCost(cost)
	d.b.SetId(id)
	d.b.SetDifficulty(dif)
	d.b.SetTask(task)
	return d.b.GetAction()
}

type Builder interface {
	SetId(id int)
	SetCost(cost int)
	SetTask(task string)
	SetDifficulty(d string)
	GetAction() *Action
}

type ConcreteBuilder struct {
	a *Action
}

func (c *ConcreteBuilder) SetId(id int) {
	c.a.Id = id
}

func (c *ConcreteBuilder) SetCost(cost int) {
	c.a.Cost = cost
}

func (c *ConcreteBuilder) SetTask(task string) {
	c.a.Task = task
}

func (c *ConcreteBuilder) SetDifficulty(d string) {
	c.a.Difficulty = d
}

func (c *ConcreteBuilder) GetAction() *Action {
	return c.a
}

type Action struct {
	Id, Cost         int
	Task, Difficulty string
}

func main() {
	d := NewDirector()
	fmt.Println(d.Build(100, 150, "work", "hard"))
}
