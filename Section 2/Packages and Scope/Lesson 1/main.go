package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Identifier struct {
	Type       string
	Visibility string
}

func fixIdentifierName(name, visibility string, needFix int) (string, error, int) {
	if visibility == "public" {
		runes := []rune(name)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		if string(runes) != (name) {
			needFix += 1
			return string(runes), fmt.Errorf("error"), needFix
		}
		return string(runes), nil, needFix
	}

	runes := []rune(name)
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}
	if string(runes) != (name) {
		needFix += 1
		return string(runes), fmt.Errorf("error"), needFix
	}
	return string(runes), nil, needFix
}

func main() {
	var packageDefinitions string
	var accessAttempts string

	fmt.Scanln(&packageDefinitions)
	fmt.Scanln(&accessAttempts)

	packages := make(map[string]map[string]Identifier)

	needFix := 0

	fmt.Println("=== PACKAGE VISIBILITY ANALYSIS ===")
	definitions := strings.Split(packageDefinitions, "|")
	for _, defin := range definitions {
		parts := strings.SplitN(defin, ":", 2)
		if len(parts) < 2 {
			continue
		}
		packageName := parts[0]
		identifiersStr := parts[1]

		if _, exists := packages[packageName]; !exists {
			packages[packageName] = make(map[string]Identifier)
		}
		fmt.Printf("Package: %s\n", packageName)

		identifiers := strings.Split(identifiersStr, ",")
		for _, id := range identifiers {
			fields := strings.Split(id, ":")
			if len(fields) != 3 {
				continue
			}
			name, err := "", error(nil)
			name, err, needFix = fixIdentifierName(fields[0], fields[2], needFix)
			packages[packageName][name] = Identifier{
				Type:       fields[1],
				Visibility: fields[2],
			}
			if err != nil {
				fmt.Printf("- %s (%s) -> %s (visibility: %s)\n", fields[0], fields[1], name, fields[2])
			} else {
				fmt.Printf("- %s (%s) (visibility: %s) ✓\n", name, fields[1], fields[2])
			}
		}
	}
	fmt.Println("=== ACCESS VALIDATION ===")
	attempts := strings.Split(accessAttempts, ",")
	allowedCount, deniedCount, notFoundCount := 0, 0, 0
	for _, attempt := range attempts {
		att := strings.Split(attempt, ".")
		pack := att[0]
		iden := att[1]
		founded := false
		for targetPackage, idMap := range packages {
			for correctedName := range idMap {
				if iden == correctedName || strings.EqualFold(iden, correctedName) {
					founded = true
					if unicode.IsUpper([]rune(correctedName)[0]) {
						fmt.Printf("%s accessing %s.%s: ✓ ALLOWED (exported)\n", pack, targetPackage, correctedName)
						allowedCount++
					} else {
						fmt.Printf("%s accessing %s.%s: ✗ DENIED (unexported)\n", pack, targetPackage, correctedName)
						deniedCount++
					}
				}
			}
		}
		if !founded {
			fmt.Printf("%s accessing [unknown_package].%s: ✗ NOT FOUND\n", pack, iden)
			notFoundCount++
		}
	}

	fmt.Println("=== SUMMARY ===")
	fmt.Printf("Total packages analyzed: %d\n", len(definitions))
	fmt.Printf("Total identifiers processed: %d\n", len(attempts))
	fmt.Printf("Identifiers requiring fixes: %d\n", needFix)
	fmt.Printf("Access attempts: %d\n", len(attempts))
	fmt.Printf("Allowed accesses: %d\n", allowedCount)
	fmt.Printf("Denied accesses: %d\n", deniedCount)
	if needFix > 0 {
		fmt.Println("Recommendation: Fix identifier naming to follow Go conventions")
	} else {
		fmt.Println("All identifiers follow proper Go naming conventions")
	}
}
