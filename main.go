package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	fmt.Println()
	color.HiBlue("   ************************************")
	color.HiBlue("   *              PowerZ              *")
	color.HiBlue("   ************************************")
	fmt.Println()

	minBase := 1
	var maxBase int
	var power int
	var questions int

	// Collect input parameters
	askQuestion(color.YellowString("-> You want to test the power of? "), &power)
	askQuestion(color.YellowString("-> Numbers should be from 1 to? "), &maxBase)
	askQuestion(color.YellowString("-> How many questions? "), &questions)

	// Recap input
	fmt.Println()
	color.Yellow("   OK, so I will ask you:")
	fmt.Printf("     * %s questions\n", color.YellowString("%d", questions))
	fmt.Printf("     * about the power of %s\n", color.YellowString("%d", power))
	fmt.Printf("     * using bases in the range %s\n", color.YellowString("[%d-%d]", minBase, maxBase))
	fmt.Println()

	// Wait for ENTER to start
	time.Sleep(time.Second)
	color.HiBlue("-> Press ENTER to begin!")
	_, err := fmt.Scanln()
	handleExitError(err)

	// Main questions loop
	start := time.Now()
	correct := 0
	var guess int
	for i := 0; i < questions; i++ {
		base := rand.Intn(maxBase) + minBase
		answer := int(math.Pow(float64(base), float64(power)))

		askQuestion(color.CyanString("-> %d to the power of %d? ", base, power), &guess)

		fmt.Println()
		if answer == guess {
			correct++
			color.Green("   CORRECT !!!")
		} else {
			color.Red("   WRONG !!! %d to the power of %d is %d", base, power, answer)
		}
		fmt.Println()
	}
	end := time.Now()

	// Working out final details
	duration := end.Sub(start)
	percent := float64(correct) / float64(questions) * float64(100)

	// Print out results
	color.HiBlue("   ************************************")
	color.HiGreen("   You got %d correct", correct)
	color.HiRed("   You got %d wrong", questions-correct)
	color.Blue("   Final Score: %.2f%%", percent)
	color.Blue("   You took %s to answer", duration)
	color.HiBlue("   ************************************")
	fmt.Println()
}

func handleExitError(err error) {
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

func askQuestion(text string, toScan ...any) {
	fmt.Printf("%s", text)
	_, err := fmt.Scanln(toScan...)
	handleExitError(err)
}
