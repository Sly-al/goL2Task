package main

import "fmt"

/*
Состояние (англ. State) — поведенческий шаблон проектирования.
Используется в тех случаях, когда во время выполнения программы объект
должен менять своё поведение в зависимости от своего состояния.
Создается впечатление, что объект изменил свой класс.
Паттерн State является объектно-ориентированной реализацией конечного автомата.

- Вводит класс Context, в котором определяется интерфейс для внешнего мира.
- Вводит абстрактный класс State.
- Представляет различные "состояния" конечного автомата в виде подклассов State.
- В классе Context имеется указатель на текущее состояние, который изменяется
при изменении состояния конечного автомата.

Пример использования:
- Когда у вас есть объект, поведение которого кардинально меняется в зависимости
от внутреннего состояния, причём типов состояний много, и их код часто меняется.
*/

// MobileAlertStater provides a common interface for various states.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert implements an alert depending on its state.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert returns a alert string
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState changes state
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert is the MobileAlert constructor.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration implements vibration alert
type MobileAlertVibration struct {
}

// Alert returns a alert string
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong implements beep alert
type MobileAlertSong struct {
}

// Alert returns a alert string
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

func main() {
	n := NewMobileAlert()
	n.SetState(&MobileAlertVibration{})
	fmt.Println(n.Alert())
	n.SetState(&MobileAlertSong{})
	fmt.Println(n.Alert())
}
