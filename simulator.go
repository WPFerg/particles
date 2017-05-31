package main

import (
	"log"

	"github.com/wpferg/particles/primitives"
)

var TIME_SLICE = 1e-6
var ITERATION_COUNT = 3000

type Simulation struct {
	inputChannels  []inputChannel
	outputChannels []outputChannel
}

type inputChannel chan [][]primitives.Subcube
type outputChannel chan []primitives.Particle

func (s *Simulation) Tick(subcubes *primitives.Subcubes) []primitives.Particle {
	for i, slice := range subcubes.Cubes {
		s.inputChannels[i] <- slice
	}

	var nextParticles []primitives.Particle

	for _, channel := range s.outputChannels {
		nextParticles = append(nextParticles, (<-channel)...)
	}

	return nextParticles
}

// A thread that takes some input, ticks the particles, then returns some output.
func startSimulationThread(input inputChannel, output outputChannel) {
	var subcubes [][]primitives.Subcube
	ok := true
	for ok {
		subcubes, ok = <-input
		if ok {
			var newPositions []primitives.Particle

			for y := range subcubes {
				for z := range subcubes[y] {
					newPositions = append(newPositions, (*subcubes[y][z].Tick(TIME_SLICE))...)
				}
			}

			output <- newPositions
		}
	}
}

func startSimulationThreads(threadCount int) ([]inputChannel, []outputChannel) {
	inputChannels := make([]inputChannel, threadCount)
	outputChannels := make([]outputChannel, threadCount)

	for i := 0; i < threadCount; i++ {
		inputChannels[i] = make(inputChannel)
		outputChannels[i] = make(outputChannel)
		go startSimulationThread(inputChannels[i], outputChannels[i])
	}

	return inputChannels, outputChannels
}

func Simulate(cube *primitives.Cube, subcubes *primitives.Subcubes, particles *[]primitives.Particle) {

	log.Println("Starting simulation.")

	currentIteration := *particles

	log.Println("Creating threads")
	inputChannels, outputChannels := startSimulationThreads(len(subcubes.Cubes))

	simulation := Simulation{
		inputChannels:  inputChannels,
		outputChannels: outputChannels,
	}

	for i := 0; i < ITERATION_COUNT; i++ {
		var nextIteration []primitives.Particle

		if i%50 == 0 {
			log.Println("Running iteration", i)
		}

		nextIteration = simulation.Tick(subcubes)

		(*cube).Collide(&nextIteration)
		currentIteration = nextIteration

		go SaveFile(i, currentIteration)

		subcubes.UpdateParticlePositions(&nextIteration)
	}

	for i := range inputChannels {
		close(inputChannels[i])
		close(outputChannels[i])
	}
}
