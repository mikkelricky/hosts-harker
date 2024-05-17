/*
Copyright © 2024 Mikkel Ricky <mikkel@mikkelricky.dk>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/dchest/validator"
)

func readLines(name string) []string {
	file, err := os.Open(hostsPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func isComment(s string) bool {
	return regexp.MustCompile(`^\s*#`).MatchString(s)
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func removeHostsEntries(ip string, domains []string, input string) string {
	output := input
	for _, domain := range domains {
		output = removeHostsEntry(ip, domain, output)
	}

	return output
}

func removeHostsEntry(ip string, domain string, input string) string {
	lines := strings.Split(input, "\n")
	newLines := []string{}
	previousLine := "(not empty)"
	for _, line := range lines {
		// Keep empty lines and comments.
		if strings.TrimSpace(line) == "" || isComment(line) {
			// Keep blank line only if previous line it not blank.
			if !(isBlank(line) && !isBlank(previousLine)) {
				newLines = append(newLines, line)
			}
		} else {
			// Find IP address at start of lines followed by domains and optionally a comment.
			pattern := `(^\s*[\d.:]+)|(\s+[^\s#]+)|(#.+)`
			r := regexp.MustCompile(pattern)
			allItems := r.FindAllString(line, -1)
			someItems := []string{}
			for i, item := range allItems {
				item = strings.TrimSpace(item)
				if i == 0 || !strings.EqualFold(item, domain) {
					if i == 0 {
						item = fmt.Sprintf("%- 15s", item)
					}
					someItems = append(someItems, item)
				}
			}
			if len(someItems) > 1 {
				if !(len(someItems) == 2 && isComment(someItems[1])) {
					newLines = append(newLines, strings.Join(someItems, " "))
				}
			}

		}
		previousLine = line
	}

	// Join lines and remove trailing spaces.
	output := strings.TrimSpace(strings.Join(append(newLines, ""), "\n"))

	if !isBlank(output) {
		output += "\n"
	}

	return output
}

func addHostsEntries(ip string, domains []string, input string) string {
	output := removeHostsEntries(ip, domains, input)

	return output + fmt.Sprintf("%s # added by %s\n", formatIpAndDomains(ip, domains), mainScript)
}

func formatIpAndDomains(ip string, domains []string) string {
	return fmt.Sprintf("%- 15s %s", ip, strings.Join(domains, " "))
}

// @see https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
func confirm(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func isValidDomain(domain string) bool {
	return validateDomain(domain) == nil
}

func validateDomain(domain string) error {
	// An IP address is a valid domain, but we don't want that.
	err := validateIpAddress(domain)
	if err == nil || !validator.IsValidDomain(domain) {
		return fmt.Errorf("%q is not a valid domain", domain)
	}

	return nil
}

func isValidIpAddress(ip string) bool {
	return validateIpAddress(ip) == nil
}

func validateIpAddress(ip string) error {
	// https://stackoverflow.com/a/36760050
	r := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	if !r.MatchString(ip) {
		return fmt.Errorf("%q is not a IP address", ip)
	}

	return nil
}

func isRootUser() bool {
	return os.Geteuid() == 0
}

func requireRoot() {
	if !isRootUser() {
		log.Fatalf("Please run as root, i.e. with sudo")
	}
}
