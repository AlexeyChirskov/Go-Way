package main

import (
	"fmt"
	"io"
	"strings"

	// "io/ioutil"
	"net/http"
	"os"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		url := url
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
		fmt.Println("\n Status code = ", resp.StatusCode)
	}

	/*
		Упражнение 1.7. Вызов функции io.Copy(dst,src) выполняет чтение src и запись в dst.
		Воспользуйтесь ею вместо ioutil.ReadAll для копирования тела ответа в поток os.Stdout без необходимости
		выделения достаточно большого для хранения всего ответа буфера. Не забудьте проверить, не произошла ли ошибка при вызове io.Сору.

		_, err = io.Copy(os.Stdout, resp.Body)
	*/
	/*
		Упражнение 1.8. Измените программу fetch так, чтобы к каждому аргументу URL автоматически добавлялся
		префикс http:// в случае отсутствия в нем таково­ го. Можете воспользоваться функцией strings.HasPrefix.

		url := url
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
	*/
	/*
		Упражнение 1.9. Измените программу fetch так, чтобы она выводила код состо­яния HTTP, содержащийся в resp.Status.
		fmt.Println("\n Status code = ", resp.StatusCode)
	*/

}
