/*
Copyright © 2024 Mikkel Ricky <mikkel@mikkelricky.dk>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	godiffpatch "github.com/sourcegraph/go-diff-patch"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Args: func(cmd *cobra.Command, args []string) error {
		// Require at least one domain
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		for _, arg := range args {
			err := validateDomain(arg)
			if err == nil {
				domains = append(domains, arg)
				continue
			}
			err = validateIpAddress(arg)
			if err == nil {
				ip = arg
				continue
			}

			return fmt.Errorf("invalid argument: %s", arg)
		}

		if len(domains) == 0 {
			return fmt.Errorf("at least one domain is required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile(hostsPath)
		if err != nil {
			log.Fatal(err)
		}

		input := string(b)
		output := removeHostsEntries(ip, domains, input)

		if diff {
			patch := godiffpatch.GeneratePatch(hostsPath, input, output)
			fmt.Println(patch)
		}

		if !isRootUser() {
			fmt.Printf("Running as non-root user. No writes possible.\n")
			return
		}

		writeFile := noInteraction || confirm(fmt.Sprintf("Remove\n\n%s\n\nfrom %s", formatIpAndDomains(ip, domains), hostsPath))
		if writeFile {
			requireRoot()

			tempFile, err := os.CreateTemp("", "hosts-harker")
			if err != nil {
				log.Fatal(err)
			}
			tempName := tempFile.Name()
			os.WriteFile(tempName, []byte(output), 0644)
			os.Rename(tempName, hostsPath)
			os.Chmod(hostsPath, 0644)
			fmt.Printf("Successfully removed\n\n%s\n\nfrom %s\n", formatIpAndDomains(ip, domains), hostsPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
