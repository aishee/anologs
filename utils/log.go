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
