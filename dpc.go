package main

import (
	util "Datapack-Compressor/util"
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	flags "github.com/jessevdk/go-flags"
)

// Options コマンドラインオプション
type Options struct {
	OutputPath     string   `short:"o" long:"output-path" description:"The output file path of the datapack destination."`
	DoNotRmCmt     bool     `short:"d" long:"do-not-remove-cmt" description:"Do not delete comment lines."`
	ExcludeRmFiles []string `short:"f" long:"exclude-remove-file" description:"Files not to be deleted."`
	ShowLog        bool     `short:"s" long:"show-log" description:"View detailed logs."`
}

var opts Options

var inexts []string = []string{".mcmeta", ".png", ".nbt", ".json", ".mcfunction"}

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "dpc"
	parser.Usage = "[PATH] [OPTIONS]"
	args, err := parser.Parse()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	// フォルダ見つからない
	if !util.Exists(os.Args[1]) {
		fmt.Println("[ERROR] The specified folder cannot be found.")
		os.Exit(1)
	}

	// 正規表現ミス
	for _, v := range opts.ExcludeRmFiles {
		_, err := regexp.Compile(v)
		if err != nil {
			fmt.Println(`[ERROR] "` + v + `" is an invalid character string.`)
			os.Exit(1)
		}
	}

	soursePath := os.Args[1]

	outputDirPath, outputFileName := "", ""

	if opts.OutputPath == "" {
		outputDirPath, _ = os.Getwd()
		outputFileName = filepath.Base(soursePath) + ".zip"
	} else {
		ex := filepath.Ext(opts.OutputPath)
		// フォルダ指定された
		if ex == "" {
			outputDirPath = opts.OutputPath
			outputFileName = filepath.Base(soursePath) + ".zip"
		} else {
			outputDirPath = filepath.Dir(opts.OutputPath)
			outputFileName = filepath.Base(opts.OutputPath)
		}
	}

	file, err := os.Create(outputDirPath + "/" + outputFileName)
	defer file.Close()
	w := zip.NewWriter(file)
	defer w.Close()

	fmt.Println("Archive:  " + outputDirPath + "/" + outputFileName)
	start := time.Now()
	count := 0
	// walk
	err = filepath.Walk(soursePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 隠しファイルの除去
			if string(getRuneAt(info.Name(), 0)) == "." {
				return nil
			}
			if !contains(filepath.Ext(path), inexts) {
				// 正規表現
				if len(opts.ExcludeRmFiles) == 0 {
					return nil
				}
				for _, v := range opts.ExcludeRmFiles {
					r := regexp.MustCompile(v)
					if !r.Match([]byte(path)) {
						return nil
					}
				}
			}
			s, ierr := filepath.Rel(soursePath, path)
			f, ierr := w.Create(s)
			if ierr != nil {
				return ierr
			}
			bytes, _ := ioutil.ReadFile(path)
			if !opts.DoNotRmCmt && filepath.Ext(path) == ".mcfunction" {
				newst := ""
				for _, line := range strings.Split(string(bytes), "\n") {
					line = strings.TrimSpace(line)
					if utf8.RuneCountInString(line) < 2 || line[0] == '\n' || line[0] == '#' {
						continue
					}
					newst += line + "\n"
				}
				f.Write([]byte(newst))
			} else {
				f.Write(bytes)
			}
			if opts.ShowLog {
				fmt.Println("  adding:  " + s)
			}
			count++
		}
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	info, _ := file.Stat()
	fmt.Printf("  Total %v files(%v bytes)\n", count, info.Size())
	end := time.Now()
	fmt.Printf("Done!(%vms)", (end.Sub(start)).Milliseconds())
}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func contains(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
