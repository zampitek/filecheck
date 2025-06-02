package config

import (
	"fmt"
	"os"

	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/err"
)

// ExecRule filters the files in a []internal.FileInfo slice and executes the specified action on them.
func ExecRule(rule RuleSpec, age AgeRules, size SizeRules, action string, files []internal.FileInfo, confirm bool) {
	var matched []internal.FileInfo
	for _, file := range files {
		if matches(*rule.Filters, file, age, size) {
			matched = append(matched, file)
		}
	}

	if confirm {
		var confirmation string
		fmt.Printf("%d files have been detected by the rule \"%s\". Do you want to perform the action \"%s\"? [y/N] ", len(matched), rule.Name, action)
		fmt.Scanln(&confirmation)

		if confirmation != "y" {
			return
		}
	}
	fmt.Printf("Performing \"%s\" on %d files...\n", action, len(matched))

	for _, file := range matched {
		actionFunc, e := matchAction(action)
		if e != nil {
			err.ExitWithError(e.Error())
		}

		res := actionFunc(file)
		if res != nil {
			err.ExitWithError(e.Error())
		}
	}
}

func matches(filters Filters, file internal.FileInfo, age AgeRules, size SizeRules) bool {
	if filters.Age != nil && file.LastAccess == int16(*filters.Age) {
		return true
	}

	if filters.AgeCategory != nil {
		switch *filters.AgeCategory {
		case 0:
			if file.LastAccess < int16(age.Low) {
				return true
			}
		case 1:
			if file.LastAccess < int16(age.Medium) {
				return true
			}
		case 2:
			if file.LastAccess >= int16(age.Medium) {
				return true
			}
		}
	}

	if filters.Size != nil && file.LastAccess == int16(*filters.Size) {
		return true
	}

	if filters.SizeCategory != nil {
		switch *filters.SizeCategory {
		case 0:
			if file.Size < size.Low {
				return true
			}
		case 1:
			if file.Size < size.Medium {
				return true
			}
		case 2:
			if file.Size >= size.Medium {
				return true
			}
		}
	}

	return false
}

func matchAction(action string) (func(internal.FileInfo) error, error) {
	switch action {
	case "delete":
		return deleteFile, nil
	default:
		return nil, err.Wrap("invalid config action", err.ErrInvalidConfig)
	}
}

func deleteFile(file internal.FileInfo) error {
	path := file.Path

	e := os.Remove(path)
	if e != nil {
		return err.Wrap(fmt.Sprintf("failed deleting file %s", file.Name), e)
	}

	return nil
}
