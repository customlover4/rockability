package player

import "time"


func (p *Player) FormatDate() string {
	started, _ := time.Parse(time.DateOnly, startDate)
	started.Add(time.Duration(p.day))

	return started.Format(time.DateOnly)
}
