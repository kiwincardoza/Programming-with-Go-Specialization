package main

import (
    "fmt"
    "time"
    "sync"
    "strconv"
    //"runtime"
)


type ChopS struct{ sync.Mutex }

type Philo struct {
    number int
    leftCS, rightCS *ChopS
}

func (p Philo) eat(ch chan *Philo, ch_status *chan int, wg *sync.WaitGroup) {
       for i:=0;i<3;i++ {

           ch <- &p
           p.leftCS.Lock()
           p.rightCS.Lock()
                
           fmt.Println("starting to eat " + strconv.Itoa(p.number))
           fmt.Println("finished eating " + strconv.Itoa(p.number))
                          
           p.rightCS.Unlock()
           p.leftCS.Unlock()
           *ch_status <- p.number
           wg.Done()
       }
}


func host(ch chan *Philo, ch_status *chan int, wg sync.WaitGroup, wg1 *sync.WaitGroup) {
    for {
        if len(ch) == 2 {

            time.Sleep(1*time.Nanosecond)
            <-ch
            <-ch
            
        }

        //sli := <-ch_staus
        if len(*ch_status) == 15 {

            (*wg1).Done()
            return
            //<-(*ch_status)
        }
        //fmt.Println(len(*ch_status))

        //if runtime.NumGoroutine()==2 {
        //    return 
        //}
    }
}


func main() {
    var wg sync.WaitGroup
    var wg1 sync.WaitGroup
    wg.Add(15)
    wg1.Add(1)
    ch := make(chan *Philo, 2)
    ch_status := make(chan int, 15)

    CSticks := make([]*ChopS, 5)
    for i := 0; i < 5; i++ {
        CSticks[i] = new(ChopS)
    }
    philos := make([]*Philo, 5)
    for i := 0; i < 5; i++ {
        philos[i] = &Philo{i+1, CSticks[i], CSticks[(i+1)%5]}
    }
    go host(ch, &ch_status, wg, &wg1)
    for i := 0; i < 5; i++ {
        go philos[i].eat(ch, &ch_status, &wg)
    }
    wg.Wait()
    wg1.Wait()
           
}
