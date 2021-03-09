package egotimer

import (
	"fmt"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	count := 0
	timer := New(1*time.Second, func(t time.Time) bool {
		if count == 5 {
			return true
		}
		count++
		fmt.Printf("Попытка %d - Время срабатывания таймера: %s\n",
			count,
			t.Format("2006-01-02 15:04:05"))
		return false
	})
	go func() {
		timer.Start()
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Перезапуск")
	count = 0
	go func() {
		timer.Restart()
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Остановка")
	timer.Stop()
}
