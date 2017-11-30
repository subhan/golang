package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    var N int
    fmt.Scanf("%d",&N)
    input := make([]int, N)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    for i, val := range strings.Split(text, " "){
        val, _ := strconv.Atoi(val)
        input[i] = val
    }
    fmt.Println(sum(input))
}


func sum(arr []int) int{
    var total int
    for _, val := range arr {
        total += val
    }
    return total
}

