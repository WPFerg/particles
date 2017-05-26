package main

import (
	"log"

	"github.com/wpferg/particles/primitives"
)

func Simulate(cube *primitives.Cube, particles *[]primitives.Particle) {
	MAX_ITERATIONS := 100

	log.Println("Starting simulation.")

	currentIteration := *particles

	for i := 0; i < MAX_ITERATIONS; i++ {

		for j, particle := range currentIteration {
			otherParticles := append(currentIteration[:j], currentIteration[j+1:]...)
			particle.Tick(0.1, &otherParticles)
		}

		currentIteration = (*cube).Collide(currentIteration)

		SaveFile(i, currentIteration)
	}
}
