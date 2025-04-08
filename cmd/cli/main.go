package main

import (
	"fmt"
	"os"

	"github.com/spacebarlabs/lean_vault/pkg/commands"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "init":
		if len(args) != 0 {
			fmt.Fprintln(os.Stderr, "Error: init command takes no arguments")
			fmt.Fprintf(os.Stderr, "\nUsage: %s init\n", os.Args[0])
			os.Exit(1)
		}
		err = commands.Init()
	case "add":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Error: add command requires a key name")
			fmt.Fprintf(os.Stderr, "\nUsage: %s add <key-name>\n", os.Args[0])
			fmt.Fprintln(os.Stderr, "\nExample:")
			fmt.Fprintf(os.Stderr, "  %s add my-api-key\n", os.Args[0])
			os.Exit(1)
		}
		err = commands.Add(args[0])
	case "get":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Error: get command requires a key name")
			fmt.Fprintf(os.Stderr, "\nUsage: %s get <key-name>\n", os.Args[0])
			os.Exit(1)
		}
		// TODO: Implement get command
		fmt.Println("Get command not implemented yet")
	case "list":
		if len(args) != 0 {
			fmt.Fprintln(os.Stderr, "Error: list command takes no arguments")
			fmt.Fprintf(os.Stderr, "\nUsage: %s list\n", os.Args[0])
			os.Exit(1)
		}
		err = commands.List()
	case "remove":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Error: remove command requires a key name")
			fmt.Fprintf(os.Stderr, "\nUsage: %s remove <key-name>\n", os.Args[0])
			os.Exit(1)
		}
		// TODO: Implement remove command
		fmt.Println("Remove command not implemented yet")
	case "rotate":
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Error: rotate command requires a key name")
			fmt.Fprintf(os.Stderr, "\nUsage: %s rotate <key-name>\n", os.Args[0])
			os.Exit(1)
		}
		// TODO: Implement rotate command
		fmt.Println("Rotate command not implemented yet")
	case "usage":
		if len(args) != 0 {
			fmt.Fprintln(os.Stderr, "Error: usage command takes no arguments")
			fmt.Fprintf(os.Stderr, "\nUsage: %s usage\n", os.Args[0])
			os.Exit(1)
		}
		// TODO: Implement usage command
		fmt.Println("Usage command not implemented yet")
	case "version":
		fmt.Printf("lean_vault version %s\n", version)
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: %s <command> [arguments]

Commands:
  init                Initialize the vault
  add <key-name>      Add a new OpenRouter API key
  get <key-name>      Retrieve a stored key
  list               List all stored keys
  remove <key-name>   Remove and revoke a key
  rotate <key-name>   Rotate a key (create new + revoke old)
  usage              Display usage information for all keys
  version            Show version information

For detailed usage instructions, see: https://github.com/spacebarlabs/lean_vault
`, os.Args[0])
}
