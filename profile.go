// Copyright 2024 Jelly Terra
// Use of this source code form is governed under the MIT license.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/user"
	"strings"
	"text/template"
	"time"
)

var predefinedData map[string]any

func init() {

	currentUser, _ := user.Current()

	predefinedData = map[string]any{
		"name":  currentUser.Name,
		"year":  time.Now().Year(),
		"month": time.Now().Month(),
		"day":   time.Now().Day(),
	}
}

func isAllWhitespace(s string) bool {
	for _, r := range s {
		switch r {
		case ' ', '\t', '\r':
		default:
			return false
		}
	}
	return true
}

type Profile struct {
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
	LinePrefix string `json:"line_prefix"`
	LineSuffix string `json:"line_suffix"`
	Template   string `json:"template"`
}

func (p *Profile) Generate(w io.Writer) error {

	t, err := template.New("profile").Parse(p.Template)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)

	if err = t.Execute(buf, predefinedData); err != nil {
		return err
	}

	if p.Prefix != "" {
		if _, err := fmt.Fprintln(w, p.Prefix); err != nil {
			return err
		}
	}

	for _, line := range strings.Split(buf.String(), "\n") {
		if _, err := fmt.Fprint(w, p.LinePrefix, line, p.LineSuffix, "\n"); err != nil {
			return err
		}
	}

	if p.Suffix != "" {
		if _, err := fmt.Fprintln(w, p.Suffix); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\n"); err != nil {
		return err
	}

	return nil
}

func (p *Profile) GetContent(lines []string) ([]string, error) {

	end := 0
	if end == len(lines) {
		return nil, errors.New("no valid content found")
	}

	for lines[end] == "" {
		end++
		if end == len(lines) {
			return nil, errors.New("no valid content found")
		}
	}

	if p.Prefix != "" && strings.HasPrefix(lines[end], p.Prefix) {
		for !strings.HasSuffix(lines[end], p.Suffix) {
			end++
			if end == len(lines) {
				return nil, errors.New("no valid content found")
			}
		}
		end++
		if end == len(lines) {
			return nil, errors.New("no valid content found")
		}
	} else {
		for strings.HasPrefix(lines[end], p.LinePrefix) {
			end++
			if end == len(lines) {
				return nil, errors.New("no valid content found")
			}
		}
	}

	for lines[end] == "" {
		end++
		if end == len(lines) {
			return nil, errors.New("no valid content found")
		}
	}

	return lines[end:], nil
}

func (p *Profile) Apply(r io.Reader, w io.Writer) error {

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// Fetch source code form content.
	content, err := p.GetContent(strings.Split(string(b), "\n"))
	if err != nil {
		return err
	}

	// Write copyright header.
	err = p.Generate(w)
	if err != nil {
		return err
	}

	// Clear empty lines with idents.
	for i, line := range content {
		if isAllWhitespace(line) {
			content[i] = ""
		}
	}

	// Append content.
	if _, err := fmt.Fprint(w, strings.Join(content, "\n")); err != nil {
		return err
	}

	return nil
}
