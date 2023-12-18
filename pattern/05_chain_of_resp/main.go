package chain_of_resp

/*
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

Применимость:
1. Когда программа должна обрабатывать разнообразные запросы несколькими способами,
но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
С помощью Цепочки обязанностей вы можете связать потенциальных обработчиков в одну цепь
и при получении запроса поочерёдно спрашивать каждого из них, не хочет ли он обработать запрос.
2. Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
Цепочка обязанностей позволяет запускать обработчиков последовательно один за другим в том порядке,
в котором они находятся в цепочке.
3. Когда набор объектов, способных обработать запрос, должен задаваться динамически.
В любой момент мы можем вмешаться в существующую цепочку и переназначить связи так,
чтобы убрать или добавить новое звено.

Плюсы:
1. Уменьшает зависимость между клиентом и обработчиками.
2. Реализует принцип единственной обязанности.
3. Реализует принцип открытости/закрытости.

Минусы:
1. Запрос может остаться никем не обработанным.

Пример:
Рассмотрим паттерн Цепочка обязанностей на примере приложения больницы.
Госпиталь может иметь разные помещения, например:

    Приемное отделение
    Доктор
    Комната медикаментов
    Кассир

Когда пациент прибывает в больницу, первым делом он попадает в Приемное отделение,
оттуда – к Доктору, затем в Комнату медикаментов, после этого – к Кассиру, и так далее.
Пациент проходит по цепочке помещений, в которой каждое отправляет его по ней дальше сразу после выполнения своей функции.
Этот паттерн можно применять в случаях, когда для выполнения одного запроса есть несколько кандидатов, и когда мы не хотим, чтобы клиент сам выбирал исполнителя.
Важно знать, что клиента необходимо оградить от исполнителей, ему необходимо знать лишь о существовании первого звена цепи.
Используя пример больницы, пациент сперва попадает в Приемное отделение.
Затем, зависимо от его состояния, Приемное отделение отправляет его к следующему исполнителю в цепи.
*/

func main() {
	cashier := &Cashier{}

	// устанавливаем next для комнаты медикаментов
	medical := &Medical{}
	medical.setNext(cashier)

	// устанавливаем next для доктора
	doctor := &Doctor{}
	doctor.setNext(medical)

	// устанавливаем next для приемного отделения
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	// пришел пациент
	reception.execute(patient)
}