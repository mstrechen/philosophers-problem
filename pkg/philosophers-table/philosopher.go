package philosophers_table

import (
	"fmt"
	"io"
	"time"
)

type Philosopher struct {
	name string
	eatTimeMs, thinkTimeMs time.Duration

	forkGreaterNum, forkLessNum *Fork

	totalEatTimes int
}

const AwaitingForForkPhrase = "Philosopher %s is awaiting to fork %s\n"
const TookForkPhrase = "Philosopher %s just took fork %s\n"
const StartedToEatPhrase = "Philosopher %s just started to eat\n"
const FinishedToEatPhrase = "Philosopher %s just finished to eat\n"
const PutForkDownPhrase = "Philosopher %s just put fork %s down\n"


func (philosopher *Philosopher) thinkAndEatForever(w io.Writer){
	for {
		for _, f := range []*Fork{philosopher.forkLessNum, philosopher.forkGreaterNum} {
			_, _ = fmt.Fprintf(w, AwaitingForForkPhrase, philosopher.name, f.name)
			_ = <- f.isFree
			_, _ = fmt.Fprintf(w, TookForkPhrase, philosopher.name, f.name)
		}

		_, _ = fmt.Fprintf(w, StartedToEatPhrase, philosopher.name)
		time.Sleep(philosopher.eatTimeMs * time.Millisecond)
		_, _ = fmt.Fprintf(w, FinishedToEatPhrase, philosopher.name)
		philosopher.totalEatTimes++

		for _, f := range []*Fork{philosopher.forkGreaterNum, philosopher.forkLessNum} {
			f.isFree <- true
			_, _ = fmt.Fprintf(w, PutForkDownPhrase, philosopher.name, f.name)
		}
		time.Sleep(philosopher.thinkTimeMs)
	}
}

func NewPhilosopher(name string, eatTimeMs, thinkTimeMs int) *Philosopher{
	p := new(Philosopher)
	p.name = name
	p.eatTimeMs = time.Duration(eatTimeMs)
	p.thinkTimeMs = time.Duration(thinkTimeMs)
	p.totalEatTimes = 0
	return p
}

func (philosopher *Philosopher) takePlace(forkLessNum, forkGreaterNum *Fork) {
	philosopher.forkGreaterNum = forkGreaterNum
	philosopher.forkLessNum = forkLessNum
}