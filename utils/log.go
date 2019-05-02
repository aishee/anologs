package anologs

type Chunk struct {
  Data string `json:"data"`
  IsParam bool `json:"is_param"`
  freezed bool
}

type Log struct {
  text string
  chunks []*Chunk
}

func NLog(line string, sp Splits) *Log {
  log := Log{}
  log.text = line
  log.chunks = sp.Split(line)
  return &log
}

func (x *Log) String() string {
  s := ""
  for _, c := range x.chunks {
    s += c.Data
  }
  return s
}

func newChunk(d string) *Chunk {
  c := Chunk{}
  c.Data = d
  c.IsParam = false
  c.freezed = false
  return &c
}

func (x *Chunk) Clone() *Chunk {
  c := newChunk(x.Data)
  c.IsParam = x.IsParam
  c.freezed = x.freezed
  return c
}

func (x *Chunk) String() string {
  return x.Data
}

func (x *Chunk) equals(chunk *Chunk) bool {
  return x.Data == chunk.Data
}

func (x *Chunk) merge(chunk *Chunk) {
  if x.Data != chunk.Data {
    x.IsParam = true
    x.Data = "*"
  }
}
