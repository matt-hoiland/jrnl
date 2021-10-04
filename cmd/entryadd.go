/*
Copyright © 2021 Matt Hoiland <matthew.proctor.hoiland@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kennygrant/sanitize"
	"github.com/matt-hoiland/jrnl/data"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// entryAddCmd represents the add command
var entryAddCmd = &cobra.Command{
	Use:   "add <title> [-t <tag> ...]",
	Short: "Add a new journal by title with optional tags",
	Long: `Entry add (jrnl entry add) creates a new journal entry preconfigured with the
optional tag set, timestamped for the moment of its creation, and titled with
the provide 'title'. The produced file will have time stamped, abbreviated,
sanitized short name derived from the title.

If successful, the filename is printed to stdout to be consumed in scripts.`,
	Args: cobra.ExactArgs(1),
	RunE: runEntryAddCmd,
}

func init() {
	entryCmd.AddCommand(entryAddCmd)
}

func runEntryAddCmd(cmd *cobra.Command, args []string) error {
	// Ignoring error
	// TODO: add logging support to file
	host, _ := os.Hostname()

	entry := &data.Entry{
		Metadata: data.Metadata{
			Title: args[0],
			Date:  time.Now().Format(data.TimeLayout),
			Host:  host,
		},
	}
	entry.Filename = fmt.Sprintf("%s_%s_%s.md", strings.Split(entry.Date, "T")[0], entry.GetWeekDay(), sanitizeTitle(entry.Title))

	fp, err := os.Create(entry.Filename)
	if err != nil {
		logrus.Error(err)
		return err
	}
	err = data.WriteEntry(fp, entry)
	if err != nil {
		logrus.Error(err)
		return err
	}
	fp.Close()
	fmt.Println(entry.Filename)
	return nil
}

func sanitizeTitle(title string) string {
	title = strings.ToLower(title)
	title = sanitize.BaseName(title)
	return title
}
