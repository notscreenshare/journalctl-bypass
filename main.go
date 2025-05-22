package main

import (
	"os"
	"strings"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	lockPath := homeDir + "/b.lock"
	journalsPath := "/var/log/journal/"

	journalsFolder, err := os.ReadDir(journalsPath)
	if err != nil {
		panic(err)
	}

	journalFile := journalsPath + journalsFolder[0].Name() + "/user-1000.journal"
	println(journalFile)

	journal, err := os.ReadFile(journalFile)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(lockPath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {

			err = os.WriteFile(lockPath, journal, 0755)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		oldJournal, err := os.ReadFile(lockPath)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(journalFile, oldJournal, 0755)
		if err != nil {
			panic(err)
		}

		err = os.Remove(lockPath)
		if err != nil {
			panic(err)
		}
	}
}
