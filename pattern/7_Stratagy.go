package main

import "fmt"

/*
Паттерн Strategy переносит в отдельную иерархию классов все детали, связанные с реализацией алгоритмов.

Плюсы:
- Систему проще поддерживать и модифицировать, так как семейство алгоритмов перенесено
в отдельную иерархию классов.
- Паттерн Strategy предоставляет возможность замены одного алгоритма другим в процессе
выполнения программы.
- Паттерн Strategy позволяет скрыть детали реализации алгоритмов от клиента.

Минусы:
- Для правильной настройки системы пользователь должен знать об особенностях всех алгоритмов.
- Число классов в системе, построенной с применением паттерна Strategy, возрастает.

Пример использования:
- Отделение процедуры выбора алгоритма от его реализации.
Это позволяет сделать выбор на основании контекста.
*/

type Stratagy interface {
	Convert(i float64) float64
}

type Client struct {
	Stratagy
}

func (c *Client) SetStratagy(str string) {
	switch str {
	case "tenge":
		c.Stratagy = &RubToTenge{}
	case "sum":
		c.Stratagy = &RubToSum{}
	}
}

type RubToTenge struct {
}

func (r *RubToTenge) Convert(i float64) float64 {
	return i * 5.21
}

type RubToSum struct {
}

func (sum *RubToSum) Convert(i float64) float64 {
	return i * 139.17
}

func main() {
	client := Client{}
	client.SetStratagy("sum")
	fmt.Println(client.Convert(100))
}
