package eleven

import "fmt"

type Octopus struct {
	grid    *Grid
	xIdx    int
	yIdx    int
	energy  int
	flashed bool
}

func (o *Octopus) String() string {
	return fmt.Sprintf("%d", o.energy)
}

func (o *Octopus) Init(g *Grid, xIdx, yIdx, energy int) {
	o.grid = g
	o.xIdx = xIdx
	o.yIdx = yIdx
	o.energy = energy
}

func (o *Octopus) ResetFlash() {
	if o.flashed {
		o.energy = 0
	}
	o.flashed = false
}

func (o *Octopus) AddEnergy() {
	o.energy++
	if !o.flashed && o.energy > 9 {
		o.energy = 0
		o.flashed = true
		o.grid.Counter.Increment()

		emit(o.grid, o.xIdx-1, o.yIdx-1)
		emit(o.grid, o.xIdx, o.yIdx-1)
		emit(o.grid, o.xIdx+1, o.yIdx-1)
		emit(o.grid, o.xIdx-1, o.yIdx)
		emit(o.grid, o.xIdx+1, o.yIdx)
		emit(o.grid, o.xIdx-1, o.yIdx+1)
		emit(o.grid, o.xIdx, o.yIdx+1)
		emit(o.grid, o.xIdx+1, o.yIdx+1)
	}
}

func emit(g *Grid, targetX, targetY int) {
	if targetX < 0 || targetY < 0 || targetX > 9 || targetY > 9 {
		return
	}
	g.Octopi[targetX][targetY].AddEnergy()
}
