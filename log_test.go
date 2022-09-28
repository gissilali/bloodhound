package main

import (
	"testing"
)

func TestLog(t *testing.T) {
	logFileLines, _ := GetFileContent("./sample_log.log")
	expectedLogFileLines := []string{
		"127.0.0.1 - - [28/Sep/2022:22:15:20 +0000] \"GET /nginx_status HTTP/1.1\" 200 103 \"-\" \"nginx-amplify-agent/1.7.0-1\"",
		"127.0.0.1 - - [28/Sep/2022:22:15:40 +0000] \"GET /nginx_status HTTP/1.1\" 200 103 \"-\" \"nginx-amplify-agent/1.7.0-1\"",
	}
	got := len(logFileLines)
	expected := len(expectedLogFileLines)

	if got != expected {
		t.Errorf("Expected: %d but got: %d", expected, got)
	}
}
