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

func getOptions() (uniq.RunOptions, string, string, error) {
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
		return uniq.RunOptions{}, "", "", errors.New("error while parsing agrs")
	}

	return options, flag.Arg(0), flag.Arg(1), nil
}

func readFile(fname string) (string, error) {
	var file *os.File
	var err error

	if fname == os.Stdin.Name() {
		file = os.Stdin
	} else {
		file, err = os.Open(fname)
	}

	if err != nil {
		return "", err
	}

	reader := io.Reader(file)
	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	if file != os.Stdin {
		file.Close()
	}

	return string(data), nil
}

func writeFile(fname, stringData string) error {
	var file *os.File
	var err error

	if fname == os.Stdout.Name() {
		file = os.Stdout
	} else {
		file, err = os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0755)
	}

	if err != nil {
		return err
	}

	writer := io.Writer(file)
	_, err = writer.Write([]byte(stringData))

	if err != nil {
		return err
	}

	if file != os.Stdout {
		file.Close()
	}

	return nil
}

func main() {
	options, fin, fout, err := getOptions()

	if err != nil {
		fmt.Println(err)
		return
	}

	if fin == "" {
		fin = os.Stdin.Name()
	}

	if fout == "" {
		fout = os.Stdout.Name()
	}

	stringData, err := readFile(fin)

	if err != nil {
		fmt.Println(err)
		return
	}

	resultLines := uniq.Uniq(strings.Split(stringData, "\n"), options)

	err = writeFile(fout, strings.Join(resultLines, "\n"))

	if err != nil {
		fmt.Println(err)
		return
	}
}
