// Copyright 2024 Jelly Terra
// Use of this source code form is governed under the MIT license.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:", os.Args[0], "profile", "file...")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	err := _main()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func _main() error {

	args := flag.Args()

	var (
		profilePath = args[0]
	)

	b, err := os.ReadFile(profilePath)
	if err != nil {
		return err
	}

	var profile Profile

	err = json.Unmarshal(b, &profile)
	if err != nil {
		return err
	}

	for _, path := range flag.Args()[1:] {
		err := func() error {
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			buf := bytes.NewBuffer(nil)

			err = profile.Apply(bytes.NewReader(b), buf)
			if err != nil {
				return err
			}

			_, err = buf.WriteTo(os.Stdout)
			if err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			fmt.Fprintln(os.Stderr, "[", path, "]", err)
		}
	}

	return nil
}
