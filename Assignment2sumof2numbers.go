package main
import "fmt"

func sum(x, y int) int {
    for i := 1; i <= y; i++ {
        x++
    }
    return x
}


func main() {

  x := sum(20, 30)
    fmt.Println(x)

}
