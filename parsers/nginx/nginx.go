package nginx

import (
	"fmt"
	"regexp"
	"time"
)

type ParsedLine struct {
	ClientIp string
	DateTime string
	line     string
}

func ParseLine(line string) ParsedLine {
	parsedLine := newParsedLine(line)
	parsedLine.setIp().setDateTime()
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
	parsedLine.ClientIp = matches[0]
	return parsedLine
}

func (parsedLine *ParsedLine) setDateTime() *ParsedLine {
	lineRegex := regexp.MustCompile(`\d{2}/[A-Za-z]{3}/\d{4}:\d{2}:\d{2}:\d{2}`)
	matches := lineRegex.FindStringSubmatch(parsedLine.line)
	format := "2006-01-02T15:04:05"
	parsedDateTime, err := time.Parse("02/Jan/2006:15:04:05", matches[0])
	if err != nil {
		fmt.Println(err)
		return parsedLine
	}
	parsedLine.DateTime = parsedDateTime.Format(format)
	return parsedLine
}

func (parsedLine *ParsedLine) build() ParsedLine {
	return *parsedLine
}
