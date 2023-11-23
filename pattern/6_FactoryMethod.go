package main

import (
	"fmt"
	"log"
)

/*
Порождающий шаблон проектирования, предоставляющий подклассам (дочерним классам) интерфейс
для создания экземпляров некоторого класса. В момент создания наследники могут определить,
какой класс создавать. Иными словами, данный шаблон делегирует создание объектов наследникам
родительского класса. Это позволяет использовать в коде программы не конкретные классы,
а манипулировать абстрактными объектами на более высоком уровне.

Структура:
- Product — продукт
	определяет интерфейс объектов, создаваемых абстрактным методом;
- ConcreteProduct — конкретный продукт
	реализует интерфейс Product;
- Creator — создатель
	объявляет фабричный метод, который возвращает объект типа Product. Может также содержать реализацию этого метода «по умолчанию»;
	может вызывать фабричный метод для создания объекта типа Product;
- ConcreteCreator — конкретный создатель
	переопределяет фабричный метод таким образом, чтобы он создавал и возвращал объект класса ConcreteProduct.

Плюсы:
- позволяет сделать код создания объектов более универсальным,
	не привязываясь к конкретным классам (ConcreteProduct), а оперируя лишь общим интерфейсом (Product);
	позволяет установить связь между параллельными иерархиями классов.
Минусы:
- Может привести к созданию больших параллельных иерархий классов,
	так как для каждого класса продукта надо создать свой подкласс создателя.
Примеры использования:
- Система должна оставаться расширяемой путем добавления объектов новых типов.
	Непосредственное использование выражения new является нежелательным, так как в этом случае код
	создания объектов с указанием конкретных типов может получиться разбросанным по всему приложению.
	Тогда такие операции как добавление в систему объектов новых типов или замена объектов одного
	типа на другой будут затруднительными (подробнее в разделе Порождающие паттерны).
	Паттерн Factory Method позволяет системе оставаться независимой как от самого процесса
	порождения объектов, так и от их типов.
- Заранее известно, когда нужно создавать объект, но неизвестен его тип.
*/

type Product interface {
	Name() string
}
type BMW struct {
}

func NewBMW() *BMW {
	return &BMW{}
}

func (b *BMW) Name() string {
	return "I am not pipez korobke"
}

type Mers struct {
}

func NewMers() *Mers {
	return &Mers{}
}
func (m *Mers) Name() string {
	return "The best or nothing"
}

type Audi struct {
}

func NewAudi() *Audi {
	return &Audi{}
}

func (a *Audi) Name() string {
	return "Br Br"
}

type Creator interface {
	CreateProduct(str string) Product
}

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(str string) Product {
	var product Product
	switch str {
	case "B":
		product = NewBMW()
	case "M":
		product = NewMers()
	case "A":
		product = NewAudi()
	default:
		log.Fatal("Unknown product")
	}
	return product
}

func main() {
	c := NewCreator()
	fmt.Println(c.CreateProduct("B").Name())
	fmt.Println(c.CreateProduct("M").Name())
	fmt.Println(c.CreateProduct("A").Name())
	fmt.Println(c.CreateProduct("C").Name())
}
