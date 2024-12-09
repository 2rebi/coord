package llmtools

import (
	"strings"

	"github.com/2rebi/coord/llm"
)

func TextFromContents(c *llm.Content) string {
	if c == nil {
		return ""
	}

	var sb strings.Builder

	for i := range c.Parts {
		if c.Parts[i].Type() == llm.SegmentTypeText {
			sb.WriteString(string(c.Parts[i].(llm.Text)))
		}
	}

	return sb.String()
}
