package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const filename = "text.txt"

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	for range os.Args[1:] {
		// fmt.Println(<-ch)
		_, err := f.Write([]byte(<-ch))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())

}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

/*
Упражнение 1.10. Найдите веб-сайт, который содержит большое количество дан­ных. Исследуйте работу кеширования
путем двукратного запуска fetchall и срав нения времени запросов. Получаете ли вы каждый раз одно и тоже содержимое?
Измените fetchall так, чтобы вывод осуществлялся в файл и чтобы затем можно было его изучить.

_, err := f.Write([]byte(<-ch))

*/
/*
Упражнение 1.11. Выполните fetchall с длинным списком аргументов, таким как образцы, доступные на сайте alexa.com.
Как ведет себя программа, когда веб­ сайт просто не отвечает? (В разделе 8.9 описан механизм отслеживания таких си­туаций.)
*/
