package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type selpgArgs struct {
	s           int
	e           int
	page_len    int
	page_type   int
	in_filename string
	print_dest  string
}

func handleArgs() selpgArgs {

	sa := selpgArgs{
		s:           -1,
		e:           -1,
		page_len:    72,
		page_type:   1,
		in_filename: "",
		print_dest:  "",
	}

	snum := flag.Int("s", -1, "start page to get(mandatory)")
	enum := flag.Int("e", -1, "end page to get(mandatory)")
	lnum := flag.Int("l", 72, "lengths of per page")
	pageSign := flag.Bool("f", false, "use change page sign or not")
	printPos := flag.String("d", "", "specify the position to print")
	flag.Parse()

	if *snum == -1 {
		fmt.Fprintln(os.Stderr, "you must give start page by '-s num'")
		return sa
	}
	if *enum == -1 {
		fmt.Fprintln(os.Stderr, "you must give end page by '-e num'")
		return sa
	}
	if *snum < 1 || *enum < 1 || *lnum < 1 {
		fmt.Fprintln(os.Stderr, "parameter of -s -e and -l should be bigger than 0")
		return sa
	}
	if *enum < *snum {
		fmt.Fprintln(os.Stderr, "end page should be bigger than start page")
		return sa
	}
	if *lnum != 72 && *pageSign {
		fmt.Fprintln(os.Stderr, "can't specify -l and -f at the same time")
		return sa
	}
	if flag.NArg() > 1 {
		fmt.Fprintln(os.Stderr, "you should only give one non-flag arguments")
		return sa
	}
	sa.s = *snum
	sa.e = *enum
	sa.page_len = *lnum

	if *pageSign {
		sa.page_type = 2
	}
	if flag.NArg() > 0 {
		sa.in_filename = flag.Args()[0]
	}
	if *printPos != "" {
		sa.print_dest = *printPos
	}

	return sa
}

func type_1_output(sa selpgArgs) {
	var rd *bufio.Reader
	var iscat int
	cmd := exec.Command("cat", "-n")
	stdin, caterr := cmd.StdinPipe()
	if caterr != nil {
		fmt.Fprintln(os.Stderr, "error happened about standard input of command cat")
	}
	if sa.print_dest != "" {
		iscat = 1
	}
	cur_page := 1
	cur_line := 0
	if sa.in_filename == "" {
		rd = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(sa.in_filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		rd = bufio.NewReader(file)
	}

	for {
		line, err := rd.ReadString('\n')
		cur_line++
		if cur_line > sa.page_len {
			cur_page++
			cur_line = 1
		}
		if cur_page >= sa.s && cur_page <= sa.e {
			//fmt.Fprintf(line)
			if iscat == 1 {
				stdin.Write([]byte(line))
			} else {
				os.Stdout.Write([]byte(line))
			}
		}
		if err != nil || io.EOF == err {
			break
		}
	}
	if cur_page < sa.s {
		fmt.Fprintln(os.Stderr, "start page is bigger than total pages of the file")
	} else {
		if cur_page < sa.e {
			fmt.Fprintln(os.Stderr, "end page is bigger than total pages of the file")
		}
	}
	if iscat == 1 {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}

func type_2_output(sa selpgArgs) {
	var rd *bufio.Reader
	var iscat int
	cmd := exec.Command("cat", "-n")
	stdin, caterr := cmd.StdinPipe()
	if caterr != nil {
		fmt.Fprintln(os.Stderr, "error happened about standard input of command cat")
	}
	if sa.print_dest != "" {
		iscat = 1
	}
	cur_page := 0
	if sa.in_filename == "" {
		rd = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(sa.in_filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		rd = bufio.NewReader(file)
	}

	for {
		line, err := rd.ReadString('\f')
		line = strings.Trim(line, "\f")
		cur_page++
		if cur_page >= sa.s && cur_page <= sa.e {
			//fmt.Fprintf(line)
			if iscat == 1 {
				stdin.Write([]byte(line + "\n"))
			} else {
				os.Stdout.Write([]byte(line + "\n"))
			}
		}
		if err != nil || io.EOF == err {
			break
		}
	}
	if cur_page < sa.s {
		fmt.Fprintln(os.Stderr, "start page is bigger than total pages of the file")
	} else {
		if cur_page < sa.e {
			fmt.Fprintln(os.Stderr, "end page is bigger than total pages of the file")
		}
	}
	if iscat == 1 {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}

func main() {

	sa := handleArgs()
	if sa.s == -1 {
		return
	}
	if sa.page_type == 1 {
		type_1_output(sa)
	} else {
		type_2_output(sa)
	}
}
