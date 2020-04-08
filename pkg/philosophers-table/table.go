package philosophers_table

import (
	"fmt"
	"io"
	"time"
)

type Table struct {
	forks []*Fork
}

func NewTable(placesCount int) * Table {
	t := new(Table)
	t.forks = make([]*Fork, placesCount)
	for i, _ := range t.forks {
		t.forks[i] = new(Fork)
		t.forks[i].name = fmt.Sprintf("Fork #%d", i + 1)
		t.forks[i].isFree = make(chan bool, 1)
		t.forks[i].isFree <- true
	}
	return t
}

func outPhilosopherStats(philosophers []*Philosopher, w io.Writer){
	for {
		_, _ = fmt.Fprint(w, "============\n")
		for _, philosopher := range philosophers {
			fmt.Fprintf(w, "%s: %d\n", philosopher.name, philosopher.totalEatTimes)
		}
		_, _ = fmt.Fprint(w, "============\n")
		time.Sleep(10 * time.Second)
	}
}

func (table *Table) StartDinner(philosophers []*Philosopher, w io.Writer) {
	for i, philosopher := range philosophers {
		lessForkI, greaterForkI := i, (i + 1) % len(table.forks)
		if lessForkI > greaterForkI {
			lessForkI, greaterForkI = greaterForkI, lessForkI
		}
		philosopher.takePlace(table.forks[lessForkI], table.forks[greaterForkI])
	}
	for _, philosopher := range philosophers {
		go philosopher.thinkAndEatForever(w)
	}

	go outPhilosopherStats(philosophers, w)
}
