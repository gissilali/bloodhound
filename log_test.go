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
		ClientIp          string
		DateTime          string
		UtcOffset         string
		HttpRequestType   string
		HttpRequest       string
		UserAgent         string
		HttpRequestStatus string
		HttpReferer       string
	}{
		ClientIp:          "127.0.0.1",
		DateTime:          "2022-09-28T22:15:40",
		UtcOffset:         "+0000",
		HttpRequestType:   "GET",
		HttpRequest:       "/nginx_status",
		UserAgent:         "nginx-amplify-agent/1.7.0-1",
		HttpRequestStatus: "200",
		HttpReferer:       "-",
	}

	if parsedLine.RemoteAddress != expectedParsedLine.ClientIp {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.ClientIp, parsedLine.RemoteAddress)
	}

	if parsedLine.DateTime != expectedParsedLine.DateTime {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.DateTime, parsedLine.DateTime)
	}

	if parsedLine.UtcOffset != expectedParsedLine.UtcOffset {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.UtcOffset, parsedLine.UtcOffset)
	}

	if parsedLine.HttpRequestType != expectedParsedLine.HttpRequestType {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.HttpRequestType, parsedLine.HttpRequestType)
	}

	if parsedLine.HttpRequest != expectedParsedLine.HttpRequest {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.HttpRequest, parsedLine.HttpRequest)
	}

	if parsedLine.UserAgent != expectedParsedLine.UserAgent {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.UserAgent, parsedLine.UserAgent)
	}

	if parsedLine.HttpRequestStatus != expectedParsedLine.HttpRequestStatus {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.HttpRequestStatus, parsedLine.HttpRequestStatus)
	}

	if parsedLine.HttpReferer != expectedParsedLine.HttpReferer {
		t.Errorf("expected: %s but got: %s", expectedParsedLine.HttpReferer, parsedLine.HttpReferer)
	}

}
