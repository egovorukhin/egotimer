package egotimer

import (
	"sync"
	"time"
)

type Timer struct {
	duration time.Duration
	f        handler
	mu       sync.Mutex
	ticker   *time.Ticker
}

//Функция которая должна выполниться
//по истечении указанного времени
type handler func(t time.Time) bool

//Инициализируем таймер, передаем длительность
func New(d time.Duration, f handler) *Timer {
	return &Timer{
		duration: d,
		f:        f,
	}
}

//Запускаем таймер. Если по истечении времени функция,
//после выполнения, вернула true, то останавливаем таймер,
//иначе продолжаем движение таймера.
func (timer *Timer) Start() {
	timer.ticker = time.NewTicker(timer.duration)
	for t := range timer.ticker.C {
		if timer.f(t) {
			return
		}
	}
}

//Передаем новую переменную для времени срабатывания
func (timer *Timer) Reset(d time.Duration) {
	timer.ticker.Reset(d)
}

//Перезапускаем таймер
func (timer *Timer) Restart() {
	timer.Stop()
	timer.Start()
}

//Останавливаем таймер.
//Mutex позволяет остановить
//таймер во всех горутинах
func (timer *Timer) Stop() {
	timer.mu.Lock()
	timer.ticker.Stop()
	timer.mu.Unlock()
}
