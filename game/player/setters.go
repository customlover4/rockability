package player

import "game/game/player/region"

func (p *Player) SetName(name string) {
	p.playerName = name
}

func (p *Player) SetGroupName(name string) {
	p.groupName = name
}

func (p *Player) SetRegion(region region.Region) {
	p.Region = region
}
