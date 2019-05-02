package anologs

import (
	"encoding/json"
	"time"
)

type Format struct {
	Count     int       `json:"count"`
	Chunks    []Chunk   `json:"chunks"`
	Timestamp time.Time `json:"timestamp"`
}

func NFormat(log *Log) *Format {
	fmt := &Format{}
	fmt.Timestamp = time.Now()
	fmt.Count = 1
	fmt.Chunks = make([]Chunk, len(log.chunks))
	for idx, x := range log.chunks {
		fmt.Chunks[idx].Data = x.Data
	}
	return fmt
}

func (x *Format) matchRatio(log *Log) float64 {
	if len(log.chunks) == len(x.Chunks) {
		matched := 0
		for idx, chunk := range x.Chunks {
			if chunk.equals(log.chunks[idx]) {
				matched++
			}
		}
		return float64(matched) / float64(len(x.Chunks))
	}
	return 0
}

func (x *Format) merge(log *Log) {
	for idx, t := range log.chunks {
		x.Chunks[idx].merge(t)
	}
	x.Count++
}

func (x *Format) String() string {
	s := ""
	for _, c := range x.Chunks {
		s += c.String()
	}
	return s
}

type SModel struct {
	Formats []*Format
}

func findCloseFormat(formats []*Format, log *Log) (*Format, float64) {
	maxIdx := -1
	maxScore := 0.0
	for idx, fmt := range formats {
		score := fmt.matchRatio(log)
		if score > maxScore {
			maxIdx = idx
			maxScore = score
		}
	}
	if maxIdx < 0 {
		return nil, 0
	}
	return formats[maxIdx], maxScore
}

func (x *SModel) read(log *Log) *Format {
	if len(log.chunks) == 0 {
		return nil
	}

	fmt, maxScore := findCloseFormat(x.Formats, log)
	if len(x.Formats) == 0 || maxScore < 0.7 {
		fmt = NFormat(log)
		x.Formats = append(x.Formats, fmt)
	}
	return fmt
}

func (x *SModel) count() int {
	sum := 0
	for _, fmt := range x.Formats {
		sum += fmt.Count
	}
	return sum
}

func (x *SModel) formats() []*Format {
	return x.Formats
}

func (x *SModel) dump() ([]byte, error) {
	data, err := json.Marshal(x)
	return data, err
}
