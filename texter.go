package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func subCommandHelpMenu() {
	fmt.Println(" concat    combines 2 text files and stores it on first argument")
	fmt.Println(" replace   replaces string defined in --from with the one defined in --to option")
}

func main() {

	replaceCmd := flag.NewFlagSet("replace", flag.ExitOnError)
	replace_from_ptr := replaceCmd.String("from", "", "define which string going to be replaced")
	replace_to_ptr := replaceCmd.String("to", "", "define which string going to be replaced with the one defined in --from option")
	replace_file_ptr := replaceCmd.String("file", "", "define file that process will happen")

	concatCmd := flag.NewFlagSet("concat", flag.ExitOnError)
	concat_file_ptr := concatCmd.String("file", "", "File that you want to work with")
	concat_with_ptr := concatCmd.String("with", "", "File that you want to add to file defined in --file option")

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	switch os.Args[1] {
	case "replace":
		replaceCmd.Parse(os.Args[2:])
		// fmt.Println(*replace_file_ptr)
		fileMode, err := os.Stat(*replace_file_ptr)

		if err != nil {
			log.Fatal(err)
		}

		openMode := fileMode.Mode()

		// fmt.Println(fileMode, openMode)

		file_data, err := os.OpenFile(*replace_file_ptr, os.O_RDWR, openMode)
		// file_data, err := os.Open(*replace_file_ptr)

		if err != nil {
			log.Fatal(err)
		}

		defer file_data.Close()

		scanner := bufio.NewScanner(file_data)
		var lines []string
		ifFound := 0
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			textLine := scanner.Text()
			// print(textLine + "\n")
			if strings.Contains(textLine, *replace_from_ptr) {
				// fmt.Println("found match")
				textLine = strings.ReplaceAll(textLine, *replace_from_ptr, *replace_to_ptr)
				// fmt.Println("After Change ", textLine+"\n")
				lines = append(lines, textLine)
				ifFound += 1
			} else {
				// fmt.Println("Not found")

				lines = append(lines, textLine)
			}

		}
		if ifFound == 0 {
			log.Fatal("does not found \"", *replace_from_ptr, "\" in ", *replace_file_ptr)
		}

		fullNewText := strings.Join(lines, "\n")
		fmt.Println(fullNewText)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		file_data.Truncate(0)
		file_data.Seek(0, 0)
		file_data.WriteString(fullNewText)
		file_data.Close()

		// fmt.Println(file_data)

	case "concat":
		concatCmd.Parse(os.Args[2:])
		fileMode, err := os.Stat(*concat_file_ptr)
		if err != nil {
			log.Fatal(err)
		}
		chmodMode := fileMode.Mode()
		mainFileData, err := os.OpenFile(*concat_file_ptr, os.O_RDWR, chmodMode)
		if err != nil {
			log.Fatal(err)
		}
		defer mainFileData.Close()

		fileLines := bufio.NewScanner(mainFileData)
		var lines []string
		for fileLines.Scan() {
			textLine := fileLines.Text()
			lines = append(lines, textLine)
		}

		with_fileMode, err := os.Stat(*concat_with_ptr)
		if err != nil {
			log.Fatal(err)
		}
		with_chmodMode := with_fileMode.Mode()
		withFileData, err := os.OpenFile(*concat_with_ptr, os.O_RDWR, with_chmodMode)
		if err != nil {
			log.Fatal(err)
		}
		defer withFileData.Close()

		withFileLines := bufio.NewScanner(withFileData)

		for withFileLines.Scan() {
			newLines := withFileLines.Text()
			fmt.Println(newLines)
			lines = append(lines, newLines)
		}

		withFileData.Close()

		mainFileData.Truncate(0)
		mainFileData.Seek(0, 0)
		fullText := strings.Join(lines, "\n")
		mainFileData.WriteString(fullText)
		mainFileData.Close()

	case "help":
		subCommandHelpMenu()
	case "--help":
		subCommandHelpMenu()
	case "-h":
		subCommandHelpMenu()
	default:
		log.Fatal("wrong subcommand")
	}
}
