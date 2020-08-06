package static

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadHtml(w io.Writer) {
	pwd, err := os.Getwd()
	fileName := strings.Join([]string{pwd, "index.html"}, "/base/tcp_server/static/")
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return
	}

	w.Write(file)
}
