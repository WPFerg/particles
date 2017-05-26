package main

import (
	"log"
	"math/rand"

	"github.com/wpferg/particles/primitives"
)

func createParticles(count int) []primitives.Particle {
	particles := make([]primitives.Particle, count)

	for i := 0; i < count; i++ {
		particles[i] = primitives.Particle{
			Position: primitives.Point{rand.Float64(), rand.Float64(), rand.Float64()},
			Vector:   primitives.Vector{rand.Float64(), rand.Float64(), rand.Float64()},
		}
	}

	return particles
}

func main() {
	log.Println("Particle Simulationator")
	log.Println("Creating a cube")

	cube := primitives.Cube{
		PointA: primitives.Point{0, 0, 0},
		PointB: primitives.Point{1, 1, 1},
	}

	log.Println("Creating particles.")
	particles := createParticles(100)

	Simulate(&cube, &particles)
}
