/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"fmt"
	"github.com/bergotorino/go-launchpad/launchpad"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

// Retrun bug id and bug title out of Task title which format is:
// Bug $ID in $PACKAGE ($ISTRIBUTION): "$TITLE"
func getIdAndTitle(title string) (string, string) {
	cut := strings.Index(title, ":")

	one := title[:cut]
	two := title[cut:]

	res := two
	id := strings.Split(one, " ")[1][1:]
	return id, res[3 : len(res)-1]
}

func shorten(s string, length uint) string {
	l := len(s)
	if uint(l) <= length {
		return s
	} else {
		return s[:length] + "..."
	}
}

type Source struct {
	Dist string
	Pkg  string
}

func readConfigFile(file string) ([]Source, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(dat), "\n")

	var src []Source
	for _, l := range lines {
		// skip lines starting with '#'
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "#") {
			continue
		}
		if l == "" {
			continue
		}

		data := strings.Split(l, "/")
		data = append(data, "")

		src = append(src, Source{Dist: data[0], Pkg: data[1]})
	}

	return src, nil
}

const maxEntries = 10
const maxTitleLength = 60

func getBugsFor(lp *launchpad.Launchpad, src Source) {
	distribution, err := lp.Distributions(src.Dist)
	if err != nil {
		log.Fatal("lp.Distributions: ", err)
		return
	}

	var bugs []launchpad.BugTask
	if src.Pkg != "" {

		pkg, err := distribution.GetSourcePackage("bluez")
		if err != nil {
			log.Fatal("lp.SourcePackages: ", err)
			return
		}
		bugs, err = pkg.SearchBugs()
		if err != nil {
			log.Fatal("lp.GetBugs ", err)
			return
		}
	} else {
		bugs, err = distribution.SearchTasks()
		if err != nil {
			log.Fatal("lp.GetBugs ", err)
			return
		}
	}

	fmt.Printf("\nRecent %s bugs:\n\n", src.Dist+"/"+src.Pkg)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for i, bug := range bugs {
		id, title := getIdAndTitle(bug.Title)
		fmt.Fprintf(w, " %s\t| %s\t| %s\n",
			id, shorten(title, maxTitleLength), bug.WebLink)
		if i > maxEntries {
			break
		}
	}
	w.Flush()
}

func main() {

	// Figure out where to read data from
	rootDir := os.Getenv("SNAP_DATA")
	if rootDir == "" {
		rootDir = os.Getenv("HOME")
	}
	configFile := rootDir + "/.go-launchpad/bugsurfer.config"
	secretsFile := rootDir + "/.go-launchpad/launchpad.secrets.json"

	src, err := readConfigFile(configFile)
	if err != nil {
		fmt.Println("Error reading config file")
	}

	// Create a place to save and load the credentials
	// The file will be created if it not exists.
	sb := launchpad.SecretsFileBackend{File: secretsFile}

	// Get a handle to the Launchpad client. All further requests
	// are proxied through it.
	lp := launchpad.NewClient(nil, "Example Client")

	// Loginto Launchpad using the previously created secrets backend
	err = lp.LoginWith(&sb)
	if err != nil {
		log.Fatal("lp.Login: ", err)
		return
	}

	person, err := lp.Me()
	if err != nil {
		log.Fatal("lp.People ", err)
		return
	}
	bgs, err := person.SearchTasks()
	if err != nil {
		log.Fatal("lp.GetBugs ", err)
		return
	}

	fmt.Println("Bugs assigned to you:\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for i, bug := range bgs {
		id, title := getIdAndTitle(bug.Title)
		fmt.Fprintf(w, " %s\t| %s\t| %s\n",
			id, shorten(title, maxTitleLength), bug.WebLink)
		if i > maxEntries {
			break
		}
	}
	w.Flush()

	for _, s := range src {
		getBugsFor(lp, s)
	}
}
