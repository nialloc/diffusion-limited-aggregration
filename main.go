package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	cParticles = 5000
	maxTree    = 400
	scale      = 80.0
	RADIUS     = float32(1)
	Width      = int32(800)
	Height     = int32(800)
)

// Particle used to hold information about individual particles
type Particle struct {
	x      int32
	y      int32
	colour rl.Color
	r      float32
}

var (
	// Walkers these are the random jiggling particles
	Walkers []Particle
	// Tree these are the fixed position ones
	Tree        []Particle
	TreeMap          = map[int32]int{}
	displayMode uint = 0x3 // bits to control what is displayed
)

// Draw particle to draw itself
func (p *Particle) Draw() {

	//rl.DrawPixel(p.x, p.y, p.colour)
	rl.DrawCircle(p.x, p.y, p.r, p.colour)

}

// Update - make the particle jiggle a bit
func (p *Particle) Update() {

	x := int32(float32(p.x) + 2.0*(0.5-rand.Float32()) + 0.5)
	y := int32(float32(p.y) + 2.0*(0.5-rand.Float32()) + 0.5)
	p.x = x
	p.y = y

}

func main() {

	rand.Seed(time.Now().UnixNano())

	rl.InitWindow(Width, Height, "niall@tigduv.com's Diffusion-Limited Aggregration")
	rl.SetTargetFPS(60)

	iteration := 0

	CreateInitialWalkers()
	CreateInitialTree()

	for _, element := range Tree {
		TreeMap[element.x+element.y*Width] = 1
	}

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		if rl.IsKeyPressed(rl.KeyZ) {
			displayMode++
		}
		if rl.IsKeyPressed(rl.KeyC) {
			Walkers = make([]Particle, 0)
		}

		// on left-click generate a bunch of new particles in a cirle at this radius from the centre
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			x := float64(rl.GetMousePosition().X)
			y := float64(rl.GetMousePosition().Y)

			dx := (x - float64(Width)/2)
			dy := (y - float64(Height)/2)
			dist := math.Sqrt(dx*dx + dy*dy)
			for i := 0; i < 100; i++ {

				angle := rand.Float64() * 2 * math.Pi
				x = float64(Width)/2 + dist*math.Cos(angle)
				y = float64(Height)/2 + dist*math.Sin(angle)
				Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Blue})
			}
		}

		if displayMode&1 != 0 {
			for _, particle := range Tree {
				particle.Draw()
			}
		}

		attempts := 0
		for attempts < 100 {

			for i := range Walkers {
				Walkers[i].Update()
			}
			attempts++
			next := make([]Particle, 0)
			for i := range Walkers {
				found := false

				walker := Walkers[i]
				p := walker.x + walker.y*Width
				for _, q := range []int32{p, p + 1, p - 1, p - Width, p - Width - 1, p - Width + 1, p + Width, p + Width - 1, p + Width + 1} {
					if TreeMap[q] > 0 {
						found = true
						Tree = append(Tree, walker)
						TreeMap[p]++
						break
					}
				}

				if !found {
					next = append(next, Walkers[i])
				}

			}
			Walkers = next
		}

		// maintain the number of particles at this level.
		// for len(Walkers) < cParticles {
		// 	particle := Particle{x:Width/2-100, y:int32(float64(Width)*rand.Float64()),r:RADIUS,colour:rl.DarkGreen}
		// 	Walkers = append(Walkers,particle)
		// }

		if displayMode&2 != 0 {
			for i := range Walkers {
				Walkers[i].Draw()
			}
		}

		rl.EndDrawing()

		text := fmt.Sprintf("  FPS:%4.0f iter %d Tree %d Walkers %d\n", rl.GetFPS(), iteration, len(Tree), len(Walkers))
		fmt.Println(text)
		rl.DrawText(text, 0, 0, 20, rl.Blue)
		iteration++
	}

	rl.CloseWindow()
}

func CreateInitialWalkers() {

	if true {
		for i := 0; i < cParticles; i++ {

			angle := rand.Float64() * 2 * math.Pi
			x := Width/2 + int32(float64(scale)*math.Cos(angle))
			y := Height/2 + int32(float64(scale)*math.Sin(angle))
			Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.DarkGreen})

			angle = rand.Float64() * 2 * math.Pi
			x = Width/2 + int32(float64(1.5*scale)*math.Cos(angle))
			y = Height/2 + int32(float64(1.5*scale)*math.Sin(angle))
			Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Blue})

			angle = rand.Float64() * 2 * math.Pi
			x = Width/2 + int32(float64(2*scale)*math.Cos(angle))
			y = Height/2 + int32(float64(2*scale)*math.Sin(angle))
			Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Red})

			angle = rand.Float64() * 2 * math.Pi
			x = Width/2 + int32(float64(2.5*scale)*math.Cos(angle))
			y = Height/2 + int32(float64(2.5*scale)*math.Sin(angle))
			Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Purple})

			angle = rand.Float64() * 2 * math.Pi
			x = Width/2 + int32(float64(3.0*scale)*math.Cos(angle))
			y = Height/2 + int32(float64(3.0*scale)*math.Sin(angle))
			Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Green})

			angle = rand.Float64() * 2 * math.Pi
			x = Width/2 + int32(float64(3.5*scale)*math.Cos(angle))
			y = Height/2 + int32(float64(3.5*scale)*math.Sin(angle))
			Walkers = append(Walkers, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Black})
		}

	}
}
func CreateInitialTree() {

	mode := 2
	if mode == 0 {

		Tree = append(Tree, Particle{x: Width / 2, y: Height / 2, r: RADIUS, colour: rl.Black})
	}

	// true will put some stationary elements near the edge of the window
	if mode == 1 {
		Tree = append(Tree, Particle{x: Width / 2, y: Height - 20, r: RADIUS, colour: rl.Black})
		Tree = append(Tree, Particle{x: Width / 2, y: 20, r: RADIUS, colour: rl.Black})
		Tree = append(Tree, Particle{x: 20, y: Height / 2, r: RADIUS, colour: rl.Black})
		Tree = append(Tree, Particle{x: Width - 20, y: Height / 2, r: RADIUS, colour: rl.Black})
	}

	// create a circlular pattern of tree elements
	if mode == 2 {
		for i := 0; i < maxTree; i++ {
			dist := float64(Width/2) * 0.80

			angle := rand.Float64() * 2 * math.Pi
			x := float64(Width)/2 + dist*math.Cos(angle)
			y := float64(Height)/2 + dist*math.Sin(angle)
			Tree = append(Tree, Particle{x: int32(x), y: int32(y), r: RADIUS, colour: rl.Black})

		}
	}
}
