package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"./uniq"
)

func handleError(err error) {
	if err != nil && err.Error() != "" {
		panic(err)
	}
}

func getOptions() (uniq.RunOptions, string, string) {
	c := flag.Bool("c", false, "prefix lines by the number of occurrences")
	d := flag.Bool("d", false, "only print duplicate lines, one for each group")
	u := flag.Bool("u", false, "only print unique lines")
	f := flag.Int("f", 0, "avoid comparing the first N fields")
	s := flag.Int("s", 0, "avoid comparing the first N characters")
	i := flag.Bool("i", false, "ignore differences in case when comparing")
	flag.Parse()

	options := uniq.RunOptions{
		Count:      *c,
		Duplicates: *d,
		Unique:     *u,
		SkipFields: *f,
		SkipChars:  *s,
		IgnoreCase: *i,
	}

	flag.Parse()

	if options.Count && (options.Duplicates || options.Unique) {
		fmt.Println("You should use only one of following options: -c -d -u\nAvaliable options:")
		flag.PrintDefaults()
		handleError(errors.New("error while parsing agrs"))
	}

	return options, flag.Arg(0), flag.Arg(1)
}

func readFile(fname string) string {
	var file *os.File
	var err error

	if fname == os.Stdin.Name() {
		file = os.Stdin
	} else {
		file, err = os.Open(fname)
	}

	handleError(err)

	reader := io.Reader(file)
	data, err := ioutil.ReadAll(reader)

	handleError(err)

	if file != os.Stdin {
		file.Close()
	}

	return string(data)
}

func writeFile(fname, stringData string) {
	var file *os.File
	var err error

	if fname == os.Stdout.Name() {
		file = os.Stdout
	} else {
		file, err = os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0755)
	}

	handleError(err)

	writer := io.Writer(file)
	_, err = writer.Write([]byte(stringData))

	handleError(err)

	if file != os.Stdout {
		file.Close()
	}
}

func main() {
	options, fin, fout := getOptions()

	if fin == "" {
		fin = os.Stdin.Name()
	}

	if fout == "" {
		fout = os.Stdout.Name()
	}

	stringData := readFile(fin)

	resultLines := uniq.Uniq(strings.Split(stringData, "\n"), options)

	writeFile(fout, strings.Join(resultLines, "\n"))
}
