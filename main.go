package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/tabwriter"
)

func exitUsage(err error) {
	flag.Usage()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func usage() {
	w := flag.CommandLine.Output()
	biname := os.Args[0]
	flags := []string{"-hn"}
	fmt.Fprintln(w, "Usage:", biname, flags, "s/pattern/replace/ path")
	flag.PrintDefaults()
}

func replaceFiles(path, pattern, replace string, dry bool) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(os.Stdout, 2, 1, 2, ' ', tabwriter.TabIndent)

	walk := func(fpath string, info os.FileInfo, err error) error {
		if !re.MatchString(fpath) {
			return nil
		}

		npath := re.ReplaceAllString(fpath, replace)
		if dry {
			fmt.Fprintln(tw, fpath, "\t-->\t", npath)
		} else {
			os.Rename(fpath, npath)
		}
		return nil
	}

	if err := filepath.Walk(path, walk); err != nil {
		return err
	}
	tw.Flush()
	return nil
}

func main() {
	flag.Usage = usage

	help := flag.Bool("h", false, "Print help")
	dry := flag.Bool("n", false, "Dry run, will not actually rename any files")
	flag.Parse()

	if *help {
		exitUsage(nil)
	}

	narg := flag.NArg()
	if narg <= 1 || narg > 2 {
		exitUsage(fmt.Errorf("Expect exactly 2 args, %d were provided\n", narg))
	}

	expression := flag.Arg(0)
	if expression[0] != 's' || len(expression) < 2 {
		exitUsage(fmt.Errorf("Expression should start with 's' followed by the separator"))
		flag.Usage()
	}
	sep := string(expression[1])
	expression = strings.TrimRight(expression, sep)
	split := strings.Split(expression[2:], sep)
	if len(split) != 2 {
		fmt.Println(split)
		exitUsage(fmt.Errorf("Bad expression format %s", expression))
	}

	if err := replaceFiles(flag.Arg(1), split[0], split[1], *dry); err != nil {
		exitUsage(err)
	}
	os.Exit(0)
}
