// +build linux, darwin, !windows

// Package terminal provides methods for working with user input
package terminal

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"pkg.re/essentialkaos/go-linenoise.v3"

	"pkg.re/essentialkaos/ek.v7/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrKillSignal is error type when user cancel input
var ErrKillSignal = linenoise.ErrKillSignal

// Prompt is prompt string
var Prompt = "> "

// MaskSymbol is symbol used for masking passwords
var MaskSymbol = "*"

// MaskSymbolColorTag is fmtc color tag used for MaskSymbol output
var MaskSymbolColorTag = ""

// ////////////////////////////////////////////////////////////////////////////////// //

// ReadUI read user input
func ReadUI(title string, nonEmpty bool) (string, error) {
	return readUserInput(
		title, nonEmpty, false,
	)
}

// ReadAnswer read user answer for Y/n question
func ReadAnswer(title, defaultAnswer string) (bool, error) {
	for {
		answer, err := readUserInput(
			getAnswerTitle(title, defaultAnswer), false, false,
		)

		if err != nil {
			return false, err
		}

		if answer == "" {
			answer = defaultAnswer
		}

		switch strings.ToUpper(answer) {
		case "Y":
			return true, nil
		case "N":
			return false, nil
		default:
			PrintWarnMessage("\nPlease enter Y or N\n")
		}
	}
}

// ReadPassword read password or some private input which will be hidden
// after pressing Enter
func ReadPassword(title string, nonEmpty bool) (string, error) {
	return readUserInput(title, nonEmpty, true)
}

// PrintErrorMessage print error message
func PrintErrorMessage(message string, args ...interface{}) {
	if len(args) == 0 {
		fmtc.Fprintf(os.Stderr, "{r}%s{!}\n", message)
	} else {
		fmtc.Fprintf(os.Stderr, "{r}%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// PrintWarnMessage print warning message
func PrintWarnMessage(message string, args ...interface{}) {
	if len(args) == 0 {
		fmtc.Fprintf(os.Stderr, "{y}%s{!}\n", message)
	} else {
		fmtc.Fprintf(os.Stderr, "{y}%s{!}\n", fmt.Sprintf(message, args...))
	}
}

// PrintActionMessage print message about action currently in progress
func PrintActionMessage(message string) {
	fmtc.Printf("{*}%s:{!} ", message)
}

// PrintActionStatus print message with action execution status
func PrintActionStatus(status int) {
	switch status {
	case 0:
		fmtc.Println("{g}OK{!}")
	case 1:
		fmtc.Println("{r}ERROR{!}")
	}
}

// AddHistory add line to input history
func AddHistory(data string) {
	linenoise.AddHistory(data)
}

// SetCompletionHandler add function for autocompletion
func SetCompletionHandler(h func(input string) []string) {
	linenoise.SetCompletionHandler(h)
}

// SetHintHandler add function for input hints
func SetHintHandler(h func(input string) string) {
	linenoise.SetHintHandler(h)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getPrivateHider(message string) string {
	prefix := strings.Repeat(" ", utf8.RuneCountInString(Prompt))
	masking := strings.Repeat(MaskSymbol, utf8.RuneCountInString(message))

	return fmt.Sprintf("%s\033[1A%s", prefix, masking)
}

func getAnswerTitle(title, defaultAnswer string) string {
	if title == "" {
		return ""
	}

	switch strings.ToUpper(defaultAnswer) {
	case "Y":
		return fmt.Sprintf("%s (Y/n)", title)
	case "N":
		return fmt.Sprintf("%s (y/N)", title)
	default:
		return fmt.Sprintf("%s (y/n)", title)
	}
}

func readUserInput(title string, nonEmpty bool, private bool) (string, error) {
	if title != "" {
		fmtc.Printf("{c}%s{!}\n", title)
	}

	var (
		input string
		err   error
	)

	for {
		input, err = linenoise.Line(Prompt)

		if err != nil {
			return "", err
		}

		if nonEmpty && strings.TrimSpace(input) == "" {
			PrintWarnMessage("\nYou must enter non empty value\n")
			continue
		}

		if private && input != "" {
			if MaskSymbolColorTag == "" {
				fmt.Println(getPrivateHider(input))
			} else {
				fmtc.Println(MaskSymbolColorTag + getPrivateHider(input) + "{!}")
			}
		}

		break
	}

	return input, err
}
