package main

import (
	"log"
	"strconv"

	"io/ioutil"

	"github.com/wpferg/particles/primitives"
)

func generateFileContents(data *[]primitives.Particle) string {
	workingString := "x,y,z\n"
	for _, particle := range *data {
		workingString += strconv.FormatFloat(particle.Position.X, 'E', -1, 64) + "," +
			strconv.FormatFloat(particle.Position.Y, 'E', -1, 64) + "," +
			strconv.FormatFloat(particle.Position.Z, 'E', -1, 64) + "\n"
	}
	return workingString
}

func SaveFile(iteration int, data []primitives.Particle) {
	err := ioutil.WriteFile("./results/iteration-"+strconv.Itoa(iteration)+".csv", []byte(generateFileContents(&data)), 0777)

	if err != nil {
		log.Println("Unable to open file for saving", err.Error())
	}
}
