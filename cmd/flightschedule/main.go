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
	"log"
	"os"
	"time"

	"github.com/bergotorino/go-launchpad/launchpad"

	"text/tabwriter"
)

func main() {
	rootDir := os.Getenv("SNAP_USER_DATA")
	if rootDir == "" {
		rootDir = os.Getenv("HOME")
	}
	rootDir += "/.go-launchpad"
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		err := os.Mkdir(rootDir, os.ModePerm)
		if err != nil {
			log.Fatal("Failed to create go-launchpad dir: ", err)
		}
	}
	sb := launchpad.SecretsFileBackend{File: rootDir + "/launchpad.secrets.json"}

	lp := launchpad.NewClient(nil, "Example Client")
	err := lp.LoginWith(&sb)
	if err != nil {
		log.Fatal("lp.Login: ", err)
		return
	}

	team, err := lp.People("snappy-hwe-team")
	if err != nil {
		log.Fatal("lp.People: ", err)
		return
	}

	snaps, err := lp.GitRepositories().GetRepositories(team.SelfLink)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	for _, s := range snaps {
		gitrepository, err := lp.GitRepositories().GetByPath(s.DisplayName[3:])
		if err != nil {
			log.Fatal("Failed to get git repository")
		}
		landingcandidates, err := gitrepository.LandingCandidates()
		if err != nil {
			log.Fatal("Failed to get landing targets")
		}
		for _, mp := range landingcandidates {
			fmt.Fprintf(w, " %s\t| %s\t| %s\t| %d\t| %s\n",
				s.Name,
				mp.RegistrantLink[33:], mp.WebLink,
				uint(time.Now().Sub(mp.DateCreated).Hours()/24+0.5),
				mp.QueueStatus)
		}
	}

	w.Flush()
}
