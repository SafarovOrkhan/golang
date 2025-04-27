package main

import (
	"fmt"
	"strings"
)

type LogEntry struct {
	Level   string
	Message string
}

type Logs struct {
	Lines []LogEntry
}

func logParser(line string) (level string, message string) {
	var sliceOfSubStrings []string
	sliceOfSubStrings = strings.Split(line, " ")
	level = sliceOfSubStrings[0]
	for i := 1; i < len(sliceOfSubStrings); i++ {
		message += sliceOfSubStrings[i]
		message += " "
	}
	return level, message
}

func (self *Logs) addLog(level string, message string) {
	if level != "" && message != "" {
		item := LogEntry{Level: level, Message: message}
		self.Lines = append(self.Lines, item)
	} else {
		fmt.Println("level or message cannot be empty")
	}
}

func (self *Logs) createMap() map[string][]string {
	contentMap := make(map[string][]string)
	for i, item := range self.Lines {
		contentMap[self.Lines[i].Level] = append(contentMap[item.Level], item.Message)
	}
	return contentMap
}

func countLogs(contentMap map[string][]string) map[string]int {
	countLogs := make(map[string]int)
	for key, value := range contentMap {
		countLogs[key] = len(value)
	}
	return countLogs
}

func main() {
	inputs := []string{
		"INFO Server started on port 8080",
		"DEBUG Initializing database connection",
		"INFO New user registration completed",
		"ERROR Failed to connect to database",
		"WARNING Low memory detected",
		"INFO Request received for /api/data",
		"DEBUG Response serialized to JSON",
		"ERROR Timeout while contacting external API",
		"WARNING High CPU usage detected",
		"INFO Server shutdown initiated",
		"DEBUG Cleanup temporary files",
		"INFO Server shutdown completed successfully",
	}

	var logs Logs
	for _, item := range inputs {
		level, message := logParser(item)
		logs.addLog(level, message)
	}

	contentMap := logs.createMap()
	countLogs := countLogs(contentMap)

	for level, message := range contentMap {
		fmt.Printf("Level: %v\n", level)
		for _, item := range message {
			fmt.Printf("- %v\n", item)
		}
		fmt.Println()
	}

	fmt.Println("log counts :")
	fmt.Println(countLogs)

}
