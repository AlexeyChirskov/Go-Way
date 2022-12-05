package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {

	var duration1 time.Duration
	var duration2 time.Duration
	//Task 1.1
	//Измените программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды.
	{
		var s, sep string
		start := time.Now()
		for i := 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		duration1 = time.Until(start)
		fmt.Println(s)
	}

	//Task 1.2
	//Измените программу echo так, чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке
	{
		s, sep := " ", " "
		for i, arg := range os.Args {
			s += sep + arg
			sep = " "
			fmt.Println("i = ", i, " arg = ", arg)
		}
		fmt.Println(s)
	}

	/*Task 1.3
	Поэкспериментируйте с измерением разницы времени выполне­ ния потенциально неэффективных версий
	и версии с применением s t r i n g s . Doin. (В разделе 1.6 демонстрируется часть пакета time,
	а в разделе 11.4 — как написать тест производительности для ее систематической оценки.
	*/
	{
		start := time.Now()
		s := strings.Join(os.Args[0:], " ")
		duration2 = time.Until(start)
		fmt.Println(s)
		fmt.Println(duration1, duration2)
	}

	//dup1
	{
		counts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			counts[input.Text()]++
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}

	//dup2
	{
		counts := make(map[string]int)
		names := make(map[string]string)
		files := os.Args[1:]
		if len(files) == 0 {
			countLines(os.Stdin, counts, names)
		} else {
			for _, arg := range files {
				f, err := os.Open(arg)
				defer f.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
					continue
				}
				countLines(f, counts, names)
			}
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, names[line])
			}
		}
	}

	//dup3
	{
		counts := make(map[string]int)
		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}

	//Task 1.4
	//Измените программу dup2 так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.
	//см dup2
}
func countLines(f *os.File, counts map[string]int, names map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		names[input.Text()] = f.Name()
	}
}
