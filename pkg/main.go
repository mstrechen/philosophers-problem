package main

import (
	"github.com/mstrechen/philosophers-problem/pkg/philosophers-table"
	"os"
	"sync"
)

func main()  {
	philosophers := []*philosophers_table.Philosopher{
		philosophers_table.NewPhilosopher("Socrates", 1000, 2000),
		philosophers_table.NewPhilosopher("Plato", 1000, 2000),
		philosophers_table.NewPhilosopher("Nietzsche", 1000, 2000),
		philosophers_table.NewPhilosopher("Schopenhauer", 1000, 2000),
		philosophers_table.NewPhilosopher("Sartre", 1000, 2000),
	}
	table := philosophers_table.NewTable(5)

	table.StartDinner(philosophers, os.Stdout)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}