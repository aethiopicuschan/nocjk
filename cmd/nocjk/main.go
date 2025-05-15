package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/aethiopicuschan/nocjk/pkg/nocjk"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing/format/gitignore"
)

var rootCmd = &cobra.Command{
	Use:          "nocjk",
	Long:         "nocjk is a command-line tool for detecting CJK text",
	RunE:         run,
	SilenceUsage: true,
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok && bi.Main.Version != "" {
		rootCmd.Version = bi.Main.Version
	} else {
		rootCmd.Version = "unknown"
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	var matcher gitignore.Matcher
	ignorePath := filepath.Join(dir, ".nocjkignore")
	if f, err := os.Open(ignorePath); err == nil {
		defer f.Close()

		ps := []gitignore.Pattern{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			ps = append(ps, gitignore.ParsePattern(line, nil))
		}
		matcher = gitignore.NewMatcher(ps)
	}

	detected := false

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}

		if matcher != nil {
			rel, err := filepath.Rel(dir, path)
			if err == nil && matcher.Match(strings.Split(rel, string(filepath.Separator)), d.IsDir()) {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if d.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Failed to read %s: %v", path, err)
			return nil
		}

		text := string(data)
		detectedLinesMap := nocjk.FindCJKLines(text)

		var lines []int
		lines = append(lines, detectedLinesMap["chinese"]...)
		lines = append(lines, detectedLinesMap["japanese"]...)
		lines = append(lines, detectedLinesMap["korean"]...)
		slices.Sort(lines)
		detectedLines := slices.Compact(lines)

		if len(detectedLines) > 0 {
			detected = true
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				relPath = path
			}
			for _, line := range detectedLines {
				fmt.Printf("%s:%d\n", relPath, line)
			}
		}

		return nil
	})

	if detected {
		err = fmt.Errorf("cjk text detected")
	}
	return
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
