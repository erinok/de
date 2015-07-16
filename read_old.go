package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type mode int

const (
	UNK mode = iota
	ENG
	GER
	WORDS
)

type st struct {
	mode     mode
	eng, ger string
	words    []string
}

var all []st

func (s *st) reset() {
	all = append(all, *s)
	*s = st{}
}

func main() {
	f, err := os.Open("flash.txt")
	if err != nil {
		fatal(err)
	}
	r := bufio.NewReader(f)
	var st st
	for lineno := 1; ; lineno++ {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		c := strings.Split(l, "")[0]
		if st.mode == WORDS && (c == ">" || c == "¶") {
			st.reset()
		}
		switch c {
		case ">":
			if st.mode != UNK {
				fatal("error: line ", lineno, ": expected UNK, got ", st.mode)
			}
			st.eng = l
			st.mode = GER
		case "¶":
			if st.mode != UNK && st.mode != GER {
				fatal("error: line ", lineno, ": expected UNK or GER, got ", st.mode)
			}
			st.ger = l
			st.mode = WORDS
		default:
			if st.mode != WORDS {
				fatal("error: line ", lineno, ": expected WORDS, got ", st.mode)
			}
			st.mode = WORDS
			st.words = append(st.words, l)
		}
	}
	for _, s := range all {
		for _, w := range s.words {
			fmt.Println(w)
		}
		if s.ger != "" {
			fmt.Println("\n" + s.ger)
		}
		if s.eng != "" {
			fmt.Println("\n" + s.eng)
		}
		fmt.Println()
	}
}

func fatal(args ...interface{}) {
	fmt.Fprint(os.Stderr, args...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}
