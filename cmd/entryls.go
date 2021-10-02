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

	"github.com/spf13/cobra"
)

var any bool

// entryLsCmd represents the ls command
var entryLsCmd = &cobra.Command{
	Use:   "ls [-t <tag> ...] [--any]",
	Short: "List journal entries optionally filtered by a tag set",
	Long: `Entry ls (jrnl entry ls) lists all tracked journal entries by default. A set of
tags can be provided through the repeated '-t <tag>' flag which filters down
only entries which have all given tags. The '--any' option permits any entry to
pass the filter if it has only one of the given tags.`,
	Run: runEntryLsCmd,
}

func init() {
	entryCmd.AddCommand(entryLsCmd)

	entryLsCmd.Flags().BoolVar(&any, "any", false, "match journal entries which contain any one of the provided tags")
}

func runEntryLsCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("TODO: implement entry ls, args: %v, tags: %v, any: %v\n", args, tags, any)
}
