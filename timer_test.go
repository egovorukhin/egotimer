package egotimer

import (
	"fmt"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	//Объявляем счетчик
	count := 0
	//Инициализация таймера
	timer := New(1*time.Second, func(t time.Time) bool {
		if count == 15 {
			return true
		}
		count++
		fmt.Printf("Попытка %d - Время срабатывания таймера: %s\n",
			count,
			t.Format("2006-01-02 15:04:05"))
		return false
	})
	fmt.Printf("Количество итераций: %d\n", timer.Count)
	//Запуск в отдельной горутине
	go timer.Start()
	//Ждем некоторое время для того чтобы функция Start отработала
	time.Sleep(3 * time.Second)
	fmt.Printf("Состояние таймера: %t\n", timer.IsStarted)
	fmt.Printf("Количество итераций: %d\n", timer.Count)

	fmt.Println("Перезапуск")
	//Скидываем на ноль переменную счетчика
	count = 0
	//Перезапуск таймера в отдельной горутине
	go timer.Restart()
	//Ждем некоторое время для того чтобы функция Restart отработала
	time.Sleep(5 * time.Second)
	fmt.Printf("Количество итераций: %d\n", timer.Count)
	fmt.Println("Остановка")
	//Остановка таймера
	timer.Stop()
}
