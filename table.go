package mjlib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Table struct {
	tbl map[int32]struct{}
}

func (this *Table) init() {
	this.tbl = map[int32]struct{}{}
}

func (this *Table) check(key int32) bool {
	_, ok := this.tbl[key]
	return ok
}

func (this *Table) add(key int32) {
	this.tbl[key] = struct{}{}
}

func (this *Table) dump(name string) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	buf := bufio.NewWriter(file)
	for key := range this.tbl {
		n := int32(key)
		fmt.Fprintf(buf, "%d\n", n)
	}
	buf.Flush()
}

func (this *Table) load(name string) {
	file, _ := os.Open(name)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(buf)
		key, _ := strconv.Atoi(str)
		this.tbl[int32(key)] = struct{}{}
	}
}
