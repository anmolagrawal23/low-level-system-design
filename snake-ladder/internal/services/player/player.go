package player

type Stats struct {
	Won  int
	Loss int
}

type Player interface {
	GetStats() *Stats
	GetName() string
}

type GamePlayer struct {
	Id   int
	Info *Info
}

func (g *GamePlayer) GetName() string {
	return g.Info.Name
}

func (g *GamePlayer) GetStats() *Stats {
	return &Stats{0, 0}
}

type Info struct {
	Name  string
	Email string
}

func NewPlayer(id int, name string, email string) Player {
	info := &Info{
		Name:  name,
		Email: email,
	}
	player := &GamePlayer{Id: id, Info: info}
	return player
}
