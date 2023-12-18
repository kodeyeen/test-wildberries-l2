package factory_method

import "fmt"

/*
Фабричный метод — это порождающий паттерн проектирования,
который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов.

Применимость:
1. Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.
Благодаря этому, код производства можно расширять, не трогая основной.
Так, чтобы добавить поддержку нового продукта, нужно создать новый подкласс и определить в нём фабричный метод,
возвращая оттуда экземпляр нового продукта.

Плюсы:
1. Избавляет класс от привязки к конкретным классам продуктов.
2. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
3. Упрощает добавление новых продуктов в программу.
4. Реализует принцип открытости/закрытости.

Минусы:
Может привести к созданию больших параллельных иерархий классов,
так как для каждого класса продукта надо создать свой подкласс создателя.

Пример:
В Go невозможно реализовать классический вариант паттерна Фабричный метод,
поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность.
Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика (Symple Factory)

В этом примере мы будем создавать разные типы оружия при помощи структуры фабрики.
Сперва, мы создадим интерфейс Weapon, который определяет все методы будущих пушек.
Также имеем структуру gun (пушка), которая применяет интерфейс Weapon.
Две конкретных пушки — ak47 и musket — обе включают в себя структуру gun и не напрямую реализуют все методы от Weapon.
WeaponFactory служит фабрикой, которая создает пушку нужного типа в зависимости от аргумента на входе.
Вместо прямого взаимодействия с объектами ak47 или musket,
клиент создает экземпляры различного оружия при помощи WeaponFactory,
используя для контроля изготовления только параметры в виде строк.
*/

func main() {
	ak47, _ := NewWeapon("ak47")
	musket, _ := NewWeapon("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(w Weapon) {
	fmt.Printf("Gun: %s", w.getName())
	fmt.Println()
	fmt.Printf("Power: %d", w.getPower())
	fmt.Println()
}