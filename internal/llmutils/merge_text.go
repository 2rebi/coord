package llmutils

import "github.com/2rebi/coord/llm"

func Normalize(p []llm.Segment) []llm.Segment {
	if len(p) < 2 {
		return p
	}

	var new []llm.Segment = make([]llm.Segment, 0, len(p))
	for i := range p {
		if p[i].Type() == llm.SegmentTypeText &&
			len(new) > 0 &&
			new[len(new)-1].Type() == llm.SegmentTypeText {
			new[len(new)-1] = new[len(new)-1].(llm.Text) + p[i].(llm.Text)
		} else {
			if p[i].Type() == llm.SegmentTypeText {
				if p[i].(llm.Text) == "" {
					continue
				}
			}
			new = append(new, p[i])
		}
	}

	return new
}
