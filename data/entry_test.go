package data_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/matt-hoiland/jrnl/data"
	"github.com/stretchr/testify/assert"
)

func TestEntry_WriteEntry(t *testing.T) {
	tests := []struct {
		Name     string
		Entry    *data.Entry
		Expected string
	}{
		{
			Name: "Simple Entry",
			Entry: &data.Entry{
				Metadata: data.Metadata{
					Title:    "Simple Entry",
					Date:     "2021-01-01T00:00:00-06:00",
					Filename: "2021-01-01_Tu_simple_entry.md",
				},
			},
			Expected: "# Simple Entry\n" +
				"\n" +
				"```json\n" +
				"{\n" +
				"  \"title\": \"Simple Entry\",\n" +
				"  \"date\": \"2021-01-01T00:00:00-06:00\",\n" +
				"  \"filename\": \"2021-01-01_Tu_simple_entry.md\"\n" +
				"}\n" +
				"```\n" +
				"\n",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			wr := &strings.Builder{}
			err := data.WriteEntry(wr, test.Entry)
			assert.Nil(t, err)
			fmt.Print(wr.String())
			assert.Equal(t, test.Expected, wr.String())
		})
	}
}
