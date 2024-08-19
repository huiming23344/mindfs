package meta

type dataServer struct {
	Id     string
	Ip     string
	Port   int
	Chunks map[string]*Chunk
}
