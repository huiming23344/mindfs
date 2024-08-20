package meta

type DataServer struct {
	Id     string
	Ip     string
	Port   int
	Chunks map[string]*Chunk
}
