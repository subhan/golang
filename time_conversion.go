package main

import  (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    arr := strings.Split(text, ":")
    h := arr[0]
    m := arr[1]
    s := arr[2]
    ap := s[2:]
    s = s[:2]

    arr = make([]string, 0, 2)
    arr = append(arr, h)
    arr = append(arr, m)
    arr = append(arr, s)
    if "12" == h && "AM" ==  ap {
        arr[0] = "00"
    }else if "12" == h && "PM" == ap {
        arr[0] =  "12"
    }else if  "PM" == ap {
        temp, _ := strconv.Atoi(h)
        arr[0] = strconv.Itoa(temp+12)
    }

    fmt.Println(strings.Join(arr, ":"))
}
