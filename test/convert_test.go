package test

import (
	"cli88/convert"
	"testing"
)

func TestConvertPdf(t *testing.T) {
	convert.ReadPdf(
		"/Users/shenlvyu/personal_workspace/my_note/WORK/work/i18n.note_副本.pdf",
		"/Users/shenlvyu/personal_workspace/my_note/WORK/work/i18n副本.md",
	)
}
