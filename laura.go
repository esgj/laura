package laura

type Engine struct {
	classes map[string]*Class
}

type Class map[string]Token

type Token struct {
	count int
	Score float32
}

func New() *Engine {
	return &Engine{}
}