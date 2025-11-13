package main

import "fmt"

func StartQuizInfo() {
	fmt.Println("=== Quiz Informatique ===")
	fmt.Println("=== Quiz Informatique Générale ===")

	totalQuestions := 10
	score := 0
	// Question 1
	fmt.Println("Quel est le système d’exploitation libre le plus utilisé ?")
	fmt.Println("1. Windows")
	fmt.Println("2. macOS")
	fmt.Println("3. Linux")
	if CheckAnswer(3) {
		score++
	}
	// Question 2
	fmt.Println("Quel langage est compilé ?")
	fmt.Println("1. Python")
	fmt.Println("2. Go")
	fmt.Println("3. JavaScript")
	if CheckAnswer(2) {
		score++
	}
	// Question 3
	fmt.Println("Que signifie HTML ?")
	fmt.Println("1. HyperText Markup Language")
	fmt.Println("2. HighText Machine Learning")
	fmt.Println("3. Hyper Transfer Main Line")
	if CheckAnswer(1) {
		score++
	}
	// Question 4
	fmt.Println("Quel language de programmation est principalement utilisé pour le développement d'applications Android?")
	fmt.Println("1. Swift")
	fmt.Println("2. Java")
	fmt.Println("3. Python")
	if CheckAnswer(2) {
		score++
	}
	// Question 5
	fmt.Println("Que signifie CPU?")
	fmt.Println("1. Central Processing Unit")
	fmt.Println("2. Computer Power Unit")
	fmt.Println("3. Central Program Unit")
	if CheckAnswer(1) {
		score++
	}
	// Question 6
	fmt.Println("Quel protocole est utilisé pour sécuriser les communications sur Internet?")
	fmt.Println("1. HTTP")
	fmt.Println("2. FTP")
	fmt.Println("3. HTTPS")
	if CheckAnswer(3) {
		score++
	}
	// Question 7
	fmt.Println("Quel type de mémoire est volatile, c'est a dire qu'elle perd ses données lorsqu'on éteint l'ordinateur?")
	fmt.Println("1. ROM")
	fmt.Println("2. RAM")
	fmt.Println("3. HDD")
	if CheckAnswer(2) {
		score++
	}
	// Question 8
	fmt.Println("Que signifie SQL ?")
	fmt.Println("1. Structured Query Language")
	fmt.Println("2. Simple Query Language")
	fmt.Println("3. Sequential Question Logic")
	if CheckAnswer(1) {
		score++
	}
	// Question 9
	fmt.Println("Quelle structure de données utilise le principe LIFO (Last In, First Out)?")
	fmt.Println("1. File d'attente")
	fmt.Println("2. pile")
	fmt.Println("3. Liste chainée")
	if CheckAnswer(2) {
		score++
	}
	// Question 10
	fmt.Println("Quel est le role principal d'un pare-feu dans les réseaux informatiques ?")
	fmt.Println("1. Windows")
	fmt.Println("2. macOS")
	fmt.Println("3. Linux")
	if CheckAnswer(2) {
	}
	score++
	CalculateScore(score, totalQuestions)
}

func CheckAnswer(correctAnswer int) bool {
	var userAnswer int
	fmt.Print("Ta réponse (numéro) : ")
	fmt.Scan(&userAnswer)

	if userAnswer == correctAnswer {
		fmt.Println("✅ Correct !")
		return true
	} else {
		fmt.Println("❌ Mauvaise réponse.")
		return false
	}
}

func CalculateScore(score int, total int) {
	fmt.Printf("Tu as %d/%d bonnes réponses.\n", score, total)

	if score <= 4 {
		fmt.Println("🔰 Niveau : Apprenti codeur")
	} else if score == 5 {
		fmt.Println("💻 Niveau : Développeur en progression")
	} else {
		fmt.Println("🤯 Niveau : Développeur confirmé")
	}
}
