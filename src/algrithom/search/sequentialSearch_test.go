package search

import (
	"bufio"
	// "fmt"
	"io/ioutil"
	// "log"
	"os"
	"strings"
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	sequence := Sequence{}

	read := func(path string) []byte {
		fi, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer fi.Close()
		fd, err := ioutil.ReadAll(fi)
		if err != nil {
			panic(err)
		}
		return fd
	}

	scannerWords := func(input string) {
		scanner := bufio.NewScanner(strings.NewReader(input))
		// Set the split function for the scanning operation.
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			word := Key(strings.Trim(scanner.Text(), ",.'!?"))
			if v, ok := sequence.get(word); ok {
				sequence.put(word, v+1)
			} else {
				sequence.put(word, 1)
			}
		}
		if err := scanner.Err(); err != nil {
			t.Errorf("[ERROR]reading input:", err)
		}
	}

	scannerWords(string(read("SequentialSearch.txt")))

	if sequence.Len() != 1 {
		t.Errorf("[EXCEPT]:1\t[FACTOR]=%v", sequence.Len())
	}
	if v, ok := sequence.get("is"); ok {
		if v != 4 {
			t.Errorf("[EXCEPT]:4\t[FACTOR]=%v", v)
		}
	} else {
		t.Errorf("[EXCEPT]:found 'is'\t[FACTOR]=not found 'is'")
	}
}

func BenchmarkSequentialSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sequence := Sequence{}

		read := func(path string) []byte {
			fi, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer fi.Close()
			fd, err := ioutil.ReadAll(fi)
			if err != nil {
				panic(err)
			}
			return fd
		}

		scannerWords := func(input string) {
			scanner := bufio.NewScanner(strings.NewReader(input))
			// Set the split function for the scanning operation.
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				word := Key(strings.Trim(scanner.Text(), ",.'!?"))
				if v, ok := sequence.get(word); ok {
					sequence.put(word, v+1)
				} else {
					sequence.put(word, 1)
				}
			}
			if err := scanner.Err(); err != nil {
				b.Errorf("[ERROR]reading input:", err)
			}
		}

		scannerWords(string(read("Twilight.txt")))
	}
}
