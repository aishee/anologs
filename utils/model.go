package anologs

type Model interface {
	read(log *Log) *Format
	count() int
	dump() ([]byte, error)
}
