package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
)

var (
	colorWhite   *color.Color = color.New(color.FgWhite)
	colorMagenta *color.Color = color.New(color.FgMagenta)
	colorCyan    *color.Color = color.New(color.FgCyan)
	colorGreen   *color.Color = color.New(color.FgGreen)
	colorYellow  *color.Color = color.New(color.FgYellow)
	colorRed     *color.Color = color.New(color.FgRed)
	colorBlue    *color.Color = color.New(color.FgBlue)
)

func ColorfulRequest(str string) {
	lines := strings.Split(str, "\n")
	if printOption&printReqHeader == printReqHeader {
		strs := strings.Split(lines[0], " ")
		colorMagenta.Print(strs[0], " ")
		colorCyan.Print(strs[1], " ")
		colorMagenta.Println(strs[2])
	}
	for _, line := range lines[1:] {
		substr := strings.SplitN(line, ":", 2)
		if len(substr) < 2 {
			colorWhite.Println(line)
		} else {
			colorWhite.Print(substr[0] + ":")
			colorCyan.Println(substr[1])
		}
	}
}

func ColorfulResponse(str, contenttype string, pretty bool) {
	match, err := regexp.MatchString(contentJsonRegex, contenttype)
	if err != nil {
		log.Fatalln("failed to compile regex", err)
	}
	if match {
		ColorfulJson(str, pretty)
	} else {
		ColorfulHTML(str)
	}
}

func ColorfulJson(str string, pretty bool) {
	formatter := jsoncolor.NewFormatter()
	formatter.SpaceColor = colorWhite
	formatter.CommaColor = colorWhite
	formatter.ColonColor = colorWhite
	formatter.ObjectColor = colorWhite
	formatter.ArrayColor = colorWhite
	formatter.FieldQuoteColor = colorWhite
	formatter.FieldColor = colorBlue
	formatter.StringQuoteColor = colorWhite
	formatter.StringColor = colorCyan
	formatter.TrueColor = colorYellow
	formatter.FalseColor = colorYellow
	formatter.NumberColor = colorGreen
	formatter.NullColor = colorRed
	if !pretty {
		formatter.Prefix = ""
		formatter.Indent = ""
	} else {
		formatter.Prefix = ""
		formatter.Indent = "  "
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(str)))
	if err := formatter.Format(buf, []byte(str)); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())
}

func ColorfulHTML(str string) {
	colorGreen.Println(str)
}
