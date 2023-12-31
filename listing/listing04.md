Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Выведутся числа от 0 до 9, а затем произойдет deadlock

Сначало инициализируется канал из интов
Далее отправляем функцию в отдельную горутину
Затем доходим до цикла, который читает значения из канала и горутина main блокируется, т.к. в канал пока ничего не записано
Далее в какой-то момент времени планировщик запускает созданную нами горутину
В ней мы итерируемся 10 раз до 0 до 9 и записываем числа в канал
Когда происходит первая запись в канал (число 0), происходит переход на следующую итерацию и попытка записать следующее значение (число 1),
но т.к. предыдущее значение еще не было считано, горутина блокируется.
У горутины main появляется шанс проснутся и планировщик в какой-то момент запускает ее
Она продолжает выполнение с места, где остановилась: считывает число из канала (0), выводит его,
переходит на следующую итерацию и пытается снова прочесть из канала. Не может, т.к. канал пуст, снова блокировка.
Планировщик запускает другую горутину, она записывает следующее число (1) и опять блокировка.
Снова запускается main... и так до числа 9.
Т.о. две горутины работают попеременно.
```
