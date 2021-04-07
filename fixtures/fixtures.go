package fixtures

import (
	"fmt"
	"github.com/oxodao/vibes/services"
)

func GenerateFakeData(prv *services.Provider) {
	fmt.Println("Generating fake data")

	generateQuestions(prv)

}
