package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Len", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInStrıng")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// s is a string assigned a literal value represeenting the word "hello" in the Thai language.
//go string literals are utf-8 encoded text.
//since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.

//indexing into a string produces the raw byte values at each index. this loop generates the hex values of all the bytes that constitute the code points in s.

// to count how many runes are in a string, wee can use the utf8 package.
//note that the run-time of RuneCountInString depends on thee size of the string,
//because it has to decode each utf-8 rune sequentially.
//some thai characters are represented by utf-8 rune sequeentially.
//some thai characters are represented by utf-8 code points that can span multiple bytes, so thee result-
//+of this count may be suprising.

//a range loop handles strings specially and decodees eeach rune along with its offset in the string.

//we can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.

//this deemonstrates passing a rune value to a function.

//values enclosed in single quotes are rune literals. we can compare a rune value to a rune literal direectly.
