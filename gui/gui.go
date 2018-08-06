package main

import (
	"fmt"
	"log"
	"math"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	. "github.com/wgoodall01/nbody/sim"
)

// Gui manages GUI simulation
type Gui struct {
	// sim is the simulation which the Gui will display
	sim Simulation

	// win is the window to which simulation will be drawn
	win *pixelgl.Window

	// canv is a buffer which dots are drawn to
	canv *pixel.PictureData
}

func NewGui(sim Simulation) *Gui {
	return &Gui{
		sim: sim,
	}
}

func (g *Gui) Start() {
	pixelgl.Run(g.run)
}

func (g *Gui) placeDot(pos pixel.Vec) {
	if g.canv.Bounds().Contains(pos) {
		g.canv.Pix[g.canv.Index(pos)] = colornames.White
	}
}

func (g *Gui) placeLine(min, max pixel.Vec) {
	// Clamp to integer pixels
	min = min.Map(math.Floor)
	max = max.Map(math.Floor)

	w, h := max.Sub(min).XY()
	slope := h / w

	y := min.Y
	for x := min.X; x < max.X; x++ {
		g.placeDot(pixel.V(x, y))
		y += slope
	}
}

func (g *Gui) run() {
	cfg := pixelgl.WindowConfig{
		Title:     "nbody",
		Bounds:    pixel.R(0, 0, 500, 500),
		Resizable: true,
		VSync:     true,
	}

	var err error
	g.win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal("Couldn't create window.")
	}

	// Various counters and timers
	frames := 0 // reset every second
	seconds := time.Tick(time.Second)
	start := time.Now()
	last := start

	// Resized?
	var oldBounds pixel.Rect

	for !g.win.Closed() {
		// Clear window to black.
		g.win.Clear(colornames.Black)

		// Clear canv to black
		if g.win.Bounds() != oldBounds {
			oldBounds = g.win.Bounds()
			fmt.Printf("Resizing to %v\n", g.win.Bounds())
			g.canv = pixel.MakePictureData(g.win.Bounds())
		}
		for i := range g.canv.Pix {
			g.canv.Pix[i] = colornames.Black
		}

		// Time setup
		deltaTime := time.Since(last).Seconds()
		last = time.Now()
		// time := time.Since(start).Seconds()

		// Update FPS counter
		frames++
		select {
		case <-seconds:
			g.win.SetTitle(fmt.Sprintf("%s @ %d fps, %0.2f ms/frame", cfg.Title, frames, (deltaTime * 1000)))
			frames = 0
		default:
		}

		// Get all masses from the simulation
		bounds := g.canv.Bounds()
		orig := bounds.Center()
		g.sim.Tick(deltaTime)
		g.sim.WriteFrame(func(m Mass) bool {
			pos := pixel.V(orig.X+m.Pos.X, orig.Y+m.Pos.Y)
			g.placeDot(pos)
			return true
		})

		// Update window.
		canvSprite := pixel.NewSprite(g.canv, g.canv.Bounds())
		canvSprite.Draw(g.win, pixel.IM.Moved(g.win.Bounds().Center()))
		g.win.Update()
	}
}
