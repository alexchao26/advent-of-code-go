package fetchers

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func panicWrap(err error, str string) {
	panic(fmt.Sprintf("%s: %s", str, err))
}

func ParseFlags() (day, year int, cookie string) {
	flag.IntVar(&day, "day", time.Now().Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", 2020, "AOC year")
	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "AOC session cookie")
	flag.Parse()

	if day > 25 {
		panic("day out of range")
	}

	if cookie == "" {
		panic("No session cookie set on flag or env")
	}

	return day, year, cookie
}

func GetWithAOCCookie(url string, cookie string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panicWrap(err, "making request")
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(&sessionCookie)

	cli := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := cli.Do(req)
	if err != nil {
		panicWrap(err, "making request")
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("response length is", len(body))

	if strings.HasPrefix(string(body), "Please don't repeatedly") {
		panic("Repeated request github.com/alexchao26/advent-of-code-go error")
	}

	return body
}

func WriteToFile(filename string, contents []byte) {
	err := ioutil.WriteFile(filename, contents, os.ModePerm)
	if err != nil {
		panicWrap(err, "writing file")
	}
}
