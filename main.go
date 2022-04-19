package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// func main() {
// 	wordList := getWords()
// 	hardWords := [][]string{}
// 	for i, word := range wordList {
// 		fmt.Print(i, " ")
// 		a := solveWordle(word)

// 		if len(a) > 6 {
// 			hardWords = append(hardWords, []string{word, strconv.Itoa(len(a))})
// 		}
// 	}

// 	fmt.Println(len(hardWords), "hard words: ")

// 	f, err := os.Create("hardwords.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	for _, w := range hardWords {

// 		_, err := f.WriteString(w[0] + " " + w[1] + "\n")

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

func main() {
	a := solveWordle("frame")
	fmt.Println(a)
}

func solveWordle(solution string) []string {
	wordList := getWords()
	guess := "tares"
	guesses := []string{}
	iter := 1
	for guess != solution {
		nw := []string{}
		for _, wor := range wordList {
			if wor != guess {
				nw = append(nw, wor)
			}
		}
		wordList = nw
		// fmt.Println("guessing", guess)
		guesses = append(guesses, guess)
		resp := guessWord(guess, solution)
		// fmt.Println(resp)

		// Make list of letters not in word
		lettersNotInWord := []string{}
		for _, tile := range resp {
			if tile[1] == "grey" {
				lettersNotInWord = append(lettersNotInWord, tile[0])
			}
		}

		yellowLetters := []string{}
		for _, tile := range resp {
			if tile[1] == "yellow" {
				yellowLetters = append(yellowLetters, tile[0])
			}
		}
		possibleWords := []string{}

		if len(lettersNotInWord) == 1 && len(yellowLetters) == 0 && len(wordList) != 1 && iter < 5 {
			_possibleWords := possibleWords
			pos := 0
			for i, tile := range resp {
				if tile[1] == "grey" {
					pos = i
				}
			}

			possibleLetters := []string{}

			for _, word := range wordList {
				splitWord := strings.Split(word, "")
				a := splitWord[pos]
				possibleLetters = append(possibleLetters, a)
			}

			for _, word := range wordList {
				include := true
				splitWord := strings.Split(word, "")

				// Check if word contains any of the letters not in word
				for _, letter := range possibleLetters {
					if !contains(letter, splitWord) {
						include = false
					}
				}

				if include {
					_possibleWords = append(_possibleWords, word)
				}
			}

			if len(_possibleWords) == 0 {
				for _, word := range wordList {
					include := true
					splitWord := strings.Split(word, "")

					// Check if word contains any of the letters not in word
					for _, letter := range lettersNotInWord {
						if contains(letter, splitWord) {
							include = false
						}
					}

					// Check if word has matching letters in the yellow spots
					for i, letter := range splitWord {
						if resp[i][1] == "yellow" && resp[i][0] == letter {
							include = false
						}
					}

					for i, letter := range splitWord {
						if resp[i][1] == "green" && resp[i][0] != letter {
							include = false
						}
					}

					for _, letter := range yellowLetters {
						if !contains(letter, splitWord) {
							include = false
						}
					}

					if include {
						possibleWords = append(possibleWords, word)
					}
				}
			} else {
				for _, word := range wordList {
					include := true
					splitWord := strings.Split(word, "")

					// Check if word contains any of the letters not in word
					for _, letter := range lettersNotInWord {
						if contains(letter, splitWord) {
							include = false
						}
					}

					// Check if word has matching letters in the yellow spots
					for i, letter := range splitWord {
						if resp[i][1] == "yellow" && resp[i][0] == letter {
							include = false
						}
					}

					for i, letter := range splitWord {
						if resp[i][1] == "green" && resp[i][0] != letter {
							include = false
						}
					}

					for _, letter := range yellowLetters {
						if !contains(letter, splitWord) {
							include = false
						}
					}

					if include {
						possibleWords = append(possibleWords, word)
					}
				}
			}
		} else {
			for _, word := range wordList {
				include := true
				splitWord := strings.Split(word, "")

				// Check if word contains any of the letters not in word
				for _, letter := range lettersNotInWord {
					if contains(letter, splitWord) {
						include = false
					}
				}

				// Check if word has matching letters in the yellow spots
				for i, letter := range splitWord {
					if resp[i][1] == "yellow" && resp[i][0] == letter {
						include = false
					}
				}

				for i, letter := range splitWord {
					if resp[i][1] == "green" && resp[i][0] != letter {
						include = false
					}
				}

				for _, letter := range yellowLetters {
					if !contains(letter, splitWord) {
						include = false
					}
				}

				if include {
					possibleWords = append(possibleWords, word)
				}
			}
		}

		wordList = possibleWords
		guess = possibleWords[0]
		iter++
	}
	guesses = append(guesses, guess)
	fmt.Println("solved", solution, "in", iter, "guesses")

	return guesses
}

func guessWord(guess string, solution string) [][]string {
	guessSlice := strings.Split(strings.ToLower(guess), "")
	solutionSlice := strings.Split(strings.ToLower(solution), "")

	resp := [][]string{
		{guessSlice[0], "grey"},
		{guessSlice[1], "grey"},
		{guessSlice[2], "grey"},
		{guessSlice[3], "grey"},
		{guessSlice[4], "grey"},
	}

	for i := 0; i < 5; i++ {
		if guessSlice[i] == solutionSlice[i] {
			resp[i][1] = "green"
		} else if contains(guessSlice[i], solutionSlice) {
			resp[i][1] = "yellow"
		}
	}
	return resp
}

func getWords() []string {
	wordList := []string{}

	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Loaded Words")

	return wordList
}

func contains(str string, sl []string) bool {
	doesContain := false
	for _, s := range sl {
		if s == str {
			doesContain = true
		}
	}

	return doesContain
}
