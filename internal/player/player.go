package player

type Player struct {
	name               string
	side               string
	victory, countGame int
}

// func New(name string, side string) (error, Player) {
// 	if stringWorker.Checkname(name) && stringWorker.CheckSide(side) {
// 		p := Player{side: side, name: name, victory: 0, countGame: 0}
// 		return nil, p
// 	} else {
// 		return errors.New("Incorrect parameters for creating a player"), Player{}
// 	}
// }

func New(name string, side string) Player {
	p := Player{side: side, name: name, victory: 0, countGame: 0}
	return p
}

func (p *Player) GetSideBool() string {
	return p.side
}

func (p *Player) GetName() string {
	return p.name
}
