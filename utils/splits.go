package anologs

type Splits interface {
  Split(msg string) []*Chunk
}

type SSplits struct {
  delims string
  regexList []*regexp.Regexp
  useRegex bool
}

func NSplits() SSplits {
  return NewSSplits()
}

func NSSplits() *SSplits {
  s := &SSplits{}
	s.delims = " \t!,:;[]{}()<>=|\\*\"'"
  s.useRegex = true

	heuristicsPatterns := []string{
		// DateTime
		`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.\d+`,
		`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`,
		`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`,
		// Date
		`\d{4}/\d{2}/\d{2}`,
		`\d{4}-\d{2}-\d{2}`,
		`\d{2}:\d{2}:\d{2}.\d+`,
		// Time
		`\d{2}:\d{2}:\d{2}`,
		// Mail address
		`[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*`,
		// IPv4 address
		`(\d{1,3}\.){3}\d{1,3}`,
	}

  s.regexList = make([]*regexp.Regexp, len(heuristicsPatterns))
  for idx, p := range heuristicsPatterns {
    s.regexList[idx] = regexp.MustCompile(p)
  }
  return s
}

func (x *SSplits) SetDelim(d string) {
  x.delims = d
}

func (x *SSplits) EnableRegex() {
  x.useRegex = true
}

func (x *SSplits) splitByRegex(chunk *Chunk) []*Chunk {
  if x.useRegex {
    for _, regex := range x.regexList {
      result := regex.FindAllStringIndex(chunk.Data, -1)
      if len(result) > 0 {
        pos := 0
        chunks := make([]*Chunk, len(result)*2+1)
      }
    }
  }
}
