package anologs

import (
	"log"
	"regexp"
	"strings"
)

type Splits interface {
	Split(msg string) []*Chunk
}

type SSplits struct {
	delims    string
	regexList []*regexp.Regexp
	useRegex  bool
}

func NSplits() Splits {
	return NewSSplits()
}

func NewSSplits() *SSplits {
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
				for idx, m := range result {
					chunks[idx*2] = newChunk(chunk.Data[pos:m[0]])
					chunks[idx*2+1] = newChunk(chunk.Data[m[0]:m[1]])
					chunks[idx*2+1].freezed = true
				}
				chunks[len(chunks)-1] = newChunk(chunk.Data[pos:])
				return chunks
			}
		}
	}
	res := []*Chunk{chunk}
	return res
}

func (x *SSplits) splitByDelimiter(chunk *Chunk) []*Chunk {
	var res []*Chunk
	msg := chunk.Data
	for {
		idx := strings.IndexAny(msg, x.delims)
		if idx < 0 {
			if len(msg) > 0 {
				res = append(res, newChunk(msg))
			}
			break
		}
		fwd := idx + 1
		x1 := msg[:idx]
		x2 := msg[idx:fwd]
		x3 := msg[fwd:]
		if len(x1) > 0 {
			log.Print("Add x1: ", x1)
			res = append(res, newChunk(x1))
		}
		if len(x2) > 0 {
			log.Print("Add x2: ", x2)
			res = append(res, newChunk(x2))
		}
		msg = x3
		log.Print("Remain: ", msg)
	}
	return res
}

func (x *SSplits) Split(msg string) []*Chunk {
	chunk := newChunk(msg)
	prevLen := 0
	chunks := []*Chunk{chunk}
	for prevLen != len(chunks) {
		var tmp []*Chunk
		for _, b := range chunks {
			log.Println(b)
			if b.freezed {
				tmp = append(tmp, b)
			} else {
				tmp = append(tmp, x.splitByRegex(b)...)
			}
		}
		prevLen = len(chunks)
		chunks = tmp
	}
	var res []*Chunk
	for _, b := range chunks {
		if b.freezed {
			res = append(res, b)
		} else {
			res = append(res, x.splitByDelimiter(b)...)
		}
	}
	return res
}
