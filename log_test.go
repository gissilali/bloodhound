package main

import (
	"bloodhound/parsers/nginx"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	logFileLines, _ := GetFileContent("./sample_log.log")
	expectedLogFileLines := []string{
		"127.0.0.1 - - [28/Sep/2022:22:15:20 +0000] \"GET /nginx_status HTTP/1.1\" 200 103 \"-\" \"nginx-amplify-agent/1.7.0-1\"",
		"127.0.0.1 - - [28/Sep/2022:22:15:40 +0000] \"GET /nginx_status HTTP/1.1\" 200 103 \"-\" \"nginx-amplify-agent/1.7.0-1\"",
	}
	got := len(logFileLines)
	expected := len(expectedLogFileLines)

	if !StringSlicesEqual(logFileLines, expectedLogFileLines) {
		t.Errorf("Expected log file lines to be equal.")
	}

	if got != expected {
		t.Errorf("Expected: %d but got: %d", expected, got)
	}
}

func TestParseLine(t *testing.T) {
	parsedLine := nginx.ParseLine("127.0.0.1 - - [28/Sep/2022:22:15:40 +0000] \"GET /nginx_status HTTP/1.1\" 200 103 \"-\" \"nginx-amplify-agent/1.7.0-1\"")
	expectedParsedLine := struct {
		ClientIp string
		DateTime string
	}{
		ClientIp: "127.0.0.1",
		DateTime: "2022-09-28T22:15:40",
	}

	if parsedLine.ClientIp != expectedParsedLine.ClientIp {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.ClientIp, parsedLine.ClientIp)
	}

	if parsedLine.DateTime != expectedParsedLine.DateTime {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.DateTime, parsedLine.DateTime)
	}
}
