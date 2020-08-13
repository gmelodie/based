package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/eknkc/basex"
)

func getAlphabet(base int) (string, error) {
	switch base {
	case 2:
		return "01", nil
	case 8:
		return "01234567", nil
	case 16:
		return "0123456789ABCDEF", nil
	// case 32:
	// 	return "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567", nil
	// case 64:
	// 	return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", nil
	default:
		return "", errors.New("Could not convert: unsuported base " + string(base))
	}
}

func ChangeBase(content []byte, base int) ([]byte, error) {
	alphabet, err := getAlphabet(base)
	if err != nil {
		return nil, err
	}

	encoding, err := basex.NewEncoding(alphabet)
	if err != nil {
		return nil, err
	}

	decoded, err := encoding.Decode(string(content))
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("usage: based <BASE> <INPUT_FILE> <OUTPUT_FILE>")
		os.Exit(1)
	}

	base, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	inputFile := os.Args[2]
	outputFile := os.Args[3]

	content, err := ioutil.ReadFile(inputFile)

	content = bytes.Trim(content, " \n\r\t=")
	content = bytes.ReplaceAll(content, []byte(" "), []byte(""))

	if err != nil {
		panic(err)
	}

	converted, err := ChangeBase(content, base)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outputFile, converted, 0644)
	if err != nil {
		panic(err)
	}
}
