package main

import "fmt"

/*
Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда
- это объект. Такие запросы, например, можно ставить в очередь, отменять или возобновлять.

Требуется для реализации:
- Базовый абстрактный класс Command описывающий интерфейс команды;
- Класс ConcreteCommand, реализующий команду;
- Класс Receiver, реализующий получателя и имеющий набор действий,
которые команда можем запрашивать;

Сначала клиент создает объект ConcreteCommand, конфигурируя его получателем запроса.
Этот объект также доступен инициатору. Инициатор использует его при отправке запроса,
вызывая метод execute(). Этот алгоритм напоминает работу функции обратного вызова в
процедурном программировании – функция регистрируется, чтобы быть вызванной позднее.

Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает,
как ее выполнить. Единственное, что должен знать инициатор, это как отправить команду.
Это придает системе гибкость: позволяет осуществлять динамическую замену команд, использовать
сложные составные команды, осуществлять отмену операций.

*/

type Command interface {
	Execute() string
}

type ConcreteCommand1 struct {
	r *Reciever
}

func (c *ConcreteCommand1) Execute() string {
	return c.r.Action1()
}

type ConcreteCommand2 struct {
	r *Reciever
}

func (c *ConcreteCommand2) Execute() string {
	return c.r.Action2()
}

type Reciever struct {
}

func (r *Reciever) Action1() string {
	return "Action1"
}

func (r *Reciever) Action2() string {
	return "Action2"
}
func main() {
	commands := []Command{&ConcreteCommand1{&Reciever{}}, &ConcreteCommand2{&Reciever{}}}
	for _, com := range commands {
		fmt.Println(com.Execute())
	}

}
