package time

import (
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := NewTimer(100)
	tds := make([]*TimerData, 100)
	for i := 0; i < 100; i++ {
		tds[i] = timer.Start(time.Duration(i)*time.Second+5*time.Minute, nil)
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		log.Printf("td: %s\n", tds[i])
		tds[i].Stop()
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		tds[i] = timer.Start(time.Duration(i)*time.Second+5*time.Minute, nil)
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		tds[i].Stop()
	}
	printTimer(timer)
	timer.Start(time.Second, nil)
	time.Sleep(time.Second * 2)
	if len(timer.timers) != 0 {
		t.FailNow()
	}
}

func printTimer(timer *Timer) {
	log.Printf("----------timers: %d ----------\n", len(timer.timers))
	for i := 0; i < len(timer.timers); i++ {
		log.Printf("td: %s\n", timer.timers[i])
	}
	log.Printf("--------------------\n")
}

func TestAfter(t *testing.T) {
	now := time.Now().Unix()
	after := After(time.Second * 1)
	if after.Unix()-now != 1 {
		t.FailNow()
	}
}

func TestAfterFunc(t *testing.T) {
	i := 0
	td := AfterFunc(time.Second*1, func() {
		i++
	})
	time.Sleep(time.Second * 2)
	td.Stop()
	if i != 1 {
		t.FailNow()
	}
}
