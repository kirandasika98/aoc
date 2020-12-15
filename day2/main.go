package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type policy struct {
	min, max int
	letter   byte
	password []byte
}

func (p *policy) String() string {
	return fmt.Sprintf("min: %d, max: %d, letter: %c", p.min, p.max, p.letter)
}

type password struct {
	raw string
}

type policyDecoder struct {
	buf string
}

func (d *policyDecoder) Decode(v *policy) error {
	if v == nil {
		return fmt.Errorf("there was an error")
	}
	end := strings.Index(d.buf, "-")
	v.min, _ = strconv.Atoi(d.buf[:end])
	d.buf = d.buf[end+1:]
	end = strings.Index(d.buf, " ")
	v.max, _ = strconv.Atoi(d.buf[:end])
	d.buf = d.buf[end+1:]

	v.letter = byte(d.buf[0])
	d.buf = d.buf[end+1:]

	v.password = []byte(strings.TrimSpace(d.buf[:]))
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	part1 := 0
	part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		pd := &policyDecoder{line}
		p := new(policy)
		pd.Decode(p)
		ok1 := validate(p, p.password)
		ok2 := validate2(p, p.password)
		if ok1 {
			part1++
		}
		if ok2 {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func validate(p *policy, pass []byte) bool {
	counts := make([]int, 26)
	for _, c := range pass {
		counts[c-'a'] += 1
	}

	found := counts[p.letter-'a']
	if found >= p.min && found <= p.max {
		return true
	}

	return false
}

func validate2(p *policy, pass []byte) bool {
	counts := make([]int, 26)
	counts[pass[p.min-1]-'a'] += 8
	counts[pass[p.max-1]-'a'] += 8
	D := counts[p.letter-'a']
	if 8 == D {
		return true
	}
	return false
}
