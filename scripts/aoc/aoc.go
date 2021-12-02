// Package aoc gets inputs and prompts from adventofcode.com
package aoc

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ParseFlags() (day, year int, cookie string) {
	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", today.Year(), "AOC year")
	// defaults to env variable
	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "AOC session cookie")
	flag.Parse()

	if day > 25 || day < 1 {
		log.Fatalf("day out of range: %d", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	if cookie == "" {
		log.Fatalf("no session cookie set on flag or env var (AOC_SESSION_COOKIE)")
	}

	return day, year, cookie
}

func GetWithAOCCookie(url string, cookie string) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(&sessionCookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("making request: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("reading response body: %s", err)
	}
	fmt.Println("response length is", len(body))

	// specific error message from AOC site
	if strings.HasPrefix(string(body), "Please don't repeatedly") {
		log.Fatalf("Repeated request github.com/alexchao26/advent-of-code-go error")
	}

	return body
}

func WriteToFile(filename string, contents []byte) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}
	err = os.WriteFile(filename, contents, os.FileMode(0644))
	if err != nil {
		log.Fatalf("writing file: %s", err)
	}
}
