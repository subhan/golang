package main


import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "os"
    "io"
)


func main() {
    var s,t,a,b,m,n int
    arr := read_input()
    //arr := newDeckfromFile()
    s, t = arr[0][0], arr[0][1]
    a, b = arr[1][0], arr[1][1]
    m, n = arr[2][0], arr[2][1]
    apples := arr[3]
    oranges := arr[4]
    if a > s || s > t || t > b{
        fmt.Println(0)
        fmt.Println(0)
        os.Exit(1)
    }
    process(s,t,a,b,m,n,apples,oranges)
}


func process(s int, t int, a int, b int, m int, n int, apples []int, oranges []int) {
    total := 0
    for _, val := range apples {
        val = a+val
        if  val >= s && val <= t {
            total ++
        }
    }
    fmt.Println(total)
    total = 0
    for _, val := range oranges {
        val = b+val
        if  val >= s && val <= t {
            total ++
        }
    }
    fmt.Println(total)
}

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )
  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}


/*func newDeckfromFile() [5][]int{
    file, err := os.Open("inputs.txt")
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    defer file.Close()

    var arr [5][]int
    reader := bufio.NewReader(file)
    s, e := Readln(reader)
    i := 0
    for e == nil {
        var temp []int
        for _, val := range strings.Split(s, " ") {
            val, _ := strconv.Atoi(val)
            temp = append(temp, val)
        }
        arr[i] = temp
        i++
        s, e = Readln(reader)
    }
    return arr
}

*/
func read_input() [5][]int {
    var arr [5][]int
    i := 0
    snr := bufio.NewScanner(os.Stdin)
    for snr.Scan(){
        text := snr.Text()
        err := snr.Err()
        if len(text) == 0  || err == io.EOF{
            fmt.Println(text)
            break
        }
        text = strings.Replace(text, "\n", "", -1)
        var temp []int
        for _, val := range strings.Split(text, " ") {
            val, _ := strconv.Atoi(val)
            temp = append(temp, val)
        }
        arr[i] = temp
        i++
    }
    return arr
}
