package meme

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func GetJpegImage() ([]string, error) {
	fc, err := os.Open("jpg.txt")
	if err != nil {
		return []string{}, nil
	}
	scanner := bufio.NewScanner(fc)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fc.Close()
	return lines, nil
}

func GetRandomImage() (string, error) {
	i, _ := GetJpegImage()
	//fix -- deprecated
	rand.Seed(time.Now().Unix())
	random := i[rand.Intn(len(i))]
	return random, nil
}
