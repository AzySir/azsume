package tablepkg

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func FormatList(sections []string) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Service Principal"})
	for idx, section := range sections {
		t.AppendRows([]table.Row{
			{idx + 1, section},
		})
	}
	// t.AppendSeparator()
	t.AppendFooter(table.Row{"Total", len(sections)})
	t.Render()

	return t
}
