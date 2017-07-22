package main

import (
	"math/rand"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
)

const (
	minX = 0
	minY = 0
	maxX = 578
	maxY = 416
)

func main() {
	oak.AddScene(
		"demo",
		func(string, interface{}) {
			layer := 0
			layerTxt := render.DefFont().NewIntText(&layer, 30, 20)
			layerTxt.SetLayer(100000000)
			render.Draw(layerTxt, 0)
			NewGopher(layer)
			layer++
			event.GlobalBind(func(int, interface{}) int {
				if oak.IsDown("K") {
					NewGopher(layer)
					layer++
				}
				return 0
			}, "EnterFrame")
		},
		func() bool {
			return true
		},
		func() (string, *oak.SceneResult) {
			return "demo", nil
		},
	)

	render.SetDrawStack(
		render.NewHeap(false),
		render.NewDrawFPS(),
		render.NewLogicFPS(),
	)
	oak.Init("demo")
}

type Gopher struct {
	entities.Doodad
	deltaX, deltaY float64
	rotation       int
}

func (g *Gopher) Init() event.CID {
	g.CID = event.NextID(g)
	return g.CID
}

func NewGopher(layer int) {
	goph := Gopher{}
	goph.Doodad = entities.NewDoodad(
		rand.Float64()*576,
		rand.Float64()*416,
		render.NewReverting(render.LoadSprite("raw\\gopher11.png")),
		goph.Init())
	goph.R.SetLayer(layer)
	goph.Bind(gophEnter, "EnterFrame")
	goph.deltaX = 4 * float64(rand.Intn(2)*2-1)
	goph.deltaY = 4 * float64(rand.Intn(2)*2-1)
	goph.rotation = rand.Intn(360)
	render.Draw(goph.R, 0)
}

func gophEnter(cid int, nothing interface{}) int {
	goph := event.GetEntity(cid).(*Gopher)
	// With rotation this gets very slow
	// consider commenting out this next line to compare with/without rotation
	// We could speed this up by caching all of the 360 rotation images,
	// but that would be a different benchmark.
	goph.R.(*render.Reverting).RevertAndModify(1, render.Rotate(goph.rotation))
	if goph.X() < minX || goph.X() > maxX {
		goph.deltaX *= -1
	}
	if goph.Y() < minY || goph.Y() > maxY {
		goph.deltaY *= -1
	}
	goph.SetPos(goph.deltaX+goph.X(), goph.deltaY+goph.Y())
	goph.rotation = (goph.rotation + 1) % 360
	return 0
}
