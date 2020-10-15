package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"stashbox/pkg/archive"
	"stashbox/pkg/crawler"
)

func usage() {
	fmt.Println("Usage: stashbox <command> <options>")
	fmt.Println("")
	fmt.Println("  Where command is one of:")
	fmt.Println("    add   --  add a url to the archive")
	fmt.Println("    list  --  list all archives")
	fmt.Println("    open  --  open an archive")
	fmt.Println("")
	fmt.Println("  To see help text, you can run:")
	fmt.Println("    stashbox <command> -h")
	os.Exit(1)
}

func main() {
	// add subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addBase := addCmd.String("b", "./stashDb", "stashbox archive directory (defaults to ./stashDb)")
	url := addCmd.String("u", "", "url to download")

	//list subcommand
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listBase := listCmd.String("b", "./stashDb", "stashbox archive directory (defaults to ./stashDb)")

	//open subcommand
	openCmd := flag.NewFlagSet("open", flag.ExitOnError)
	openBase := openCmd.String("b", "./stashDb", "stashbox archive directory (defaults to ./stashDb)")
	n := openCmd.Int("n", 0, "archive number to open (from list command)")

	if len(os.Args) < 2 {
		usage()
	}

	if os.Args[1] == "-h" {
		usage()
	}

	switch os.Args[1] {
	case "add":
		if addCmd.Parse(os.Args[2:]) != nil {
			addCmd.Usage()
			os.Exit(2)
		}
		if *url == "" {
			fmt.Println("ERROR: -url is required")
			addCmd.Usage()
			os.Exit(1)
		}
		c, err := crawler.NewCrawler(*addBase)
		if err != nil {
			panic(err)
		}

		err = c.AddURL(*url)
		if err != nil {
			panic(err)
		}

		err = c.Crawl()
		if err != nil {
			panic(err)
		}

		err = c.Save()
		if err != nil {
			panic(err)
		}
	case "list":
		if listCmd.Parse(os.Args[2:]) != nil {
			listCmd.Usage()
			os.Exit(2)
		}
		archives, err := archive.GetArchives(*listBase)
		if err != nil {
			fmt.Println("Error listing archives", err)
			os.Exit(1)
		}
		fmt.Println("Archive listing...")
		for i, n := range archives {
			fmt.Printf("%d. %s [%d image(s)]\n", i+1, n.URL, len(n.Dates))
		}
	case "open":
		if openCmd.Parse(os.Args[2:]) != nil {
			openCmd.Usage()
			os.Exit(2)
		}
		archives, err := archive.GetArchives(*openBase)
		if err != nil {
			panic(err)
		}

		if *n < 1 {
			fmt.Println("Choose an archive to open:")
			for i, n := range archives {
				fmt.Printf("%d. %s [%d image(s)]\n", i+1, n.URL, len(n.Dates))
			}
			fmt.Print("\n> ")
			_, err := fmt.Scanf("%d", n)

			if err != nil {
				fmt.Printf("ERROR: reading input: %v", err)
				os.Exit(1)
			}
		}

		a := archives[*n-1]

		file := fmt.Sprintf("%s/%s/%s.pdf", *openBase, a.URL, a.Dates[len(a.Dates)-1])
		fmt.Printf("Opening: %s\n", file)
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command(file) // #nosec G204
		case "darwin":
			cmd = exec.Command("open", file) // #nosec G204
		default:
			cmd = exec.Command("xdg-open", file) // #nosec G204
		}
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("ERROR: unknown command (%s) specified\n", os.Args[1])
		usage()
	}
}
