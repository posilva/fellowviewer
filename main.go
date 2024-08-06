package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//go:embed template.tpl
var templateTpl embed.FS

type User struct {
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Manager    string `json:"manager"`
	Title      string `json:"title"`
	Department string `json:"department"`
	DateJoined string `json:"date_joined"`
	IP         string `json:"ip"`
	UserAgent  string `json:"user_agent"`
	LastSeen   string `json:"last_seen"`
}

type FeedbackValue struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}
type Feedback struct {
	ID    string        `json:"id"`
	Label string        `json:"label"`
	Value FeedbackValue `json:"value"`
}

type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

type Calendar struct {
	Name   string  `json:"name"`
	Events []Event `json:"events"`
}

type Note struct {
	Title   string   `json:"title"`
	Content []string `json:"content"`
	Start   string   `json:"start"`
	End     string   `json:"end"`
}

type Data struct {
	UserInformation User         `json:"user_information"`
	Feedback        [][]Feedback `json:"feedback"`
	Calendars       []Calendar   `json:"calendars"`
	Notes           []Note       `json:"notes"`
	Attachments     []string     `json:"attachments"`
}

func main() {
	app := &cli.App{
		Name:  "fellowviewer",
		Usage: "A tool to generate HTML from Fellow JSON exported notes",
		Authors: []*cli.Author{
			{
				Name:  "Pedro Marques da Silva",
				Email: "posilva@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Value:    "",
				Required: true,
				Usage:    "Input file path (json exported from Fellow)",
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Value:    "",
				Required: true,
				Usage:    "Output file path (html generated file)",
			},
		},
		Action: func(c *cli.Context) error {
			inputFile := c.String("input")
			outputFile := c.String("output")

			run(inputFile, outputFile)
			// Your logic here to process the input file and write to the output file

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func run(jsonF string, htmlF string) {
	file, err := os.Open(jsonF)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	data, err := io.ReadAll(file) // Replace "user_data.json" with your actual file path
	if err != nil {
		log.Panic(err)
	}

	// Parse JSON data
	var userInformation Data
	err = json.Unmarshal(data, &userInformation)
	if err != nil {
		log.Panic(err)
	}

	// Generate HTML
	tmpl, err := template.New("template.tpl").Funcs(
		template.FuncMap{
			"unscapeHtml": unscapeHtml,
		}).ParseFS(templateTpl, "template.tpl")
	if err != nil {
		log.Panic(err)
	}

	f, err := os.Create(htmlF)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, userInformation)
	if err != nil {
		panic(err)
	}
}

func unscapeHtml(str string) template.HTML {
	return template.HTML(html.UnescapeString(str))
}
