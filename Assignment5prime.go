package main

import "fmt"

func Numbers()chan int{
    x:=make(chan int)
    go func(){
    for i := 2; i <= 100; i++ {
			x <- i
}
    }()
return x
}

func prime(z<- chan int,n int)chan int{
    y:=make(chan int)
    go func(){
        for{
       if i := <-z; i % n !=0 {
            y <- i
        }
        }
    }()
    return y
}
func main(){
    x:= Numbers()
       for i := 2; i < 100; i++ {
        n:=<-x
      fmt.Printf("%v: %v\n", i+1, n)
        x=prime(x,n)
    }
}
