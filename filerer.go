package main

import (
  "bufio"
  "fmt"
  "regexp"
  "strings"
  "sync"
)

func telephoneNumbersInFile(path string) int {
  file := strings.NewReader(path)

  var telephone = regexp.MustCompile(`\(\d+\)\s\d+-\d+`)

  // do I need buffered channels here?
  jobs := make(chan string)
  results := make(chan int)

  // I think we need a wait group, not sure.
  wg := new(sync.WaitGroup)

  // start up some workers that will block and wait?
  for w := 1; w <= 3; w++ {
    wg.Add(1)
    go matchTelephoneNumbers(jobs, results, wg, telephone)
  }

  // go over a file line by line and queue up a ton of work
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    // Later I want to create a buffer of lines, not just line-by-line here ...
    jobs <- scanner.Text()
  }

  close(jobs)
  wg.Wait()

  // Add up the results from the results channel.
  // The rest of this isn't even working so ignore for now.
  counts := 0
  // for v := range results {
  //   counts += v
  // }

  return counts
}

func matchTelephoneNumbers(jobs <-chan string, results chan<- int, wg *sync.WaitGroup, telephone *regexp.Regexp) {
  // Decreasing internal counter for wait-group as soon as goroutine finishes
  defer wg.Done()

  // eventually I want to have a []string channel to work on a chunk of lines not just one line of text
  for j := range jobs {
    if telephone.MatchString(j) {
      results <- 1
    }
  }
}

func main() {
  // An artificial input source.  Normally this is a file passed on the command line.
  const input = "Foo\n(555) 123-3456\nBar\nBaz"
  numberOfTelephoneNumbers := telephoneNumbersInFile(input)
  fmt.Println(numberOfTelephoneNumbers)
}
