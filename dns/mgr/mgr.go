package main

import(
    "os"
    "os/signal"
    "time"
    "fmt"
    "syscall"
)

func main () {
    ch := make(chan os.Signal)
    signal.Notify(ch, os.Interrupt, syscall.SIGKILL)
    go ins.Run()
    for {
        s := <- ch
        fmt.Println("Go Signal s : ", s)
        ins.Stop()
        break
    }
    time.Sleep(5 * time.Second)
}
