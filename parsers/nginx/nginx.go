package nginx

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type ParsedLine struct {
	RemoteAddress     string
	DateTime          string
	line              string
	UtcOffset         string
	HttpRequestType   string
	HttpRequest       string
	HttpRequestStatus string
	UserAgent         string
	HttpReferer       string
}

func ParseLine(line string) ParsedLine {
	parsedLine := newParsedLine(line)
	parsedLine.setIp()
	parsedLine.setDateTime()
	parsedLine.setUtcOffset()
	parsedLine.setHttpRequest()
	parsedLine.setUserAgent()
	parsedLine.setHttpRequestStatus()
	parsedLine.setHttpReferer()
	return parsedLine.build()
}

func newParsedLine(line string) *ParsedLine {
	return &ParsedLine{
		line: line,
	}
}

func (parsedLine *ParsedLine) setIp() *ParsedLine {
	lineRegex := regexp.MustCompile(`(((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4})`)
	matches := lineRegex.FindStringSubmatch(parsedLine.line)
	parsedLine.RemoteAddress = matches[0]
	return parsedLine
}

func (parsedLine *ParsedLine) setDateTime() *ParsedLine {
	lineRegex := regexp.MustCompile(`\d{2}/[A-Za-z]{3}/\d{4}:\d{2}:\d{2}:\d{2}`)
	matches := lineRegex.FindStringSubmatch(parsedLine.line)
	parsedDateTime, err := time.Parse("02/Jan/2006:15:04:05", matches[0])
	if err != nil {
		fmt.Println(err)
		return parsedLine
	}
	parsedLine.DateTime = parsedDateTime.Format("2006-01-02T15:04:05")
	return parsedLine
}

func (parsedLine *ParsedLine) setUtcOffset() *ParsedLine {
	lineRegex := regexp.MustCompile(`\+\d{4}`)
	matches := lineRegex.FindStringSubmatch(parsedLine.line)
	parsedLine.UtcOffset = matches[0]
	return parsedLine
}

func (parsedLine *ParsedLine) build() ParsedLine {
	return *parsedLine
}

func (parsedLine *ParsedLine) setHttpRequest() *ParsedLine {
	lineRegex := regexp.MustCompile(`"(.*?)"`)
	matches := lineRegex.FindStringSubmatch(parsedLine.line)
	request := strings.Split(strings.Replace(matches[0], "\"", "", -1), " ")
	parsedLine.HttpRequestType = request[0]
	parsedLine.HttpRequest = request[1]
	return parsedLine
}

func (parsedLine *ParsedLine) setUserAgent() *ParsedLine {
	lineRegex := regexp.MustCompile(`"(.*?)"`)
	matches := lineRegex.FindAllStringSubmatch(parsedLine.line, 3)
	parsedLine.UserAgent = matches[2][1]
	return parsedLine
}

func (parsedLine *ParsedLine) setHttpReferer() *ParsedLine {
	lineRegex := regexp.MustCompile(`"(.*?)"`)
	matches := lineRegex.FindAllStringSubmatch(parsedLine.line, 3)
	parsedLine.HttpReferer = matches[1][1]
	return parsedLine
}
func (parsedLine *ParsedLine) setHttpRequestStatus() *ParsedLine {
	lineRegex := regexp.MustCompile(`\s\d{3}\s`)
	matches := lineRegex.FindStringSubmatch(parsedLine.line)
	parsedLine.HttpRequestStatus = strings.Replace(matches[0], " ", "", -1)
	return parsedLine
}
