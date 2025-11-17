package main

import "fmt"

func StartQuizIAData() {
	fmt.Println("――――― Quiz IA et Data ―――――")
	
	
	totalQuestions := 10
	score := 0

	// Question 1
	fmt.Println("Qu'est ce qu'un modèle d'intelligence artificielle ?")
	fmt.Println("1. Un programme capable d'apprednre a partir de données")
	fmt.Println("2. Un robot physique autonome")
	fmt.Println("3. Un systeme qui n'utilise aucune données")
	if CheckAnswer(1) {
		score++
	}
	// Question 2
	fmt.Println("Que signifie machine learning ?")
	fmt.Println("1. L'utilisation d'ordinateurs sans algorithmes")
	fmt.Println("2. L'apprentissage automatique a partir de données")
	fmt.Println("3. La création d'images a partir de texte")
	if CheckAnswer(2) {
		score++
	}
	// Question 3
	fmt.Println("Quel est l'objectif principal du Big Data ?")
	fmt.Println("1. Réduire la quantité de données disponible")
	fmt.Println("2. Gérer et analyser de grands volumes de données")
	fmt.Println("3.Construire des ordinateurs plus petits")
	if CheckAnswer(2) {
		score++
	}
	// Question 4
	fmt.Println("Que représente uen feature (variable) en data science ?")
	fmt.Println("1. Une donnée d'entrée utilisée par le modèle")
	fmt.Println("2. Le code source complet du programme")
	fmt.Println("3. La réponse finale du modèle")
	if CheckAnswer(1) {
		score++
	}
	// Question 5
	fmt.Println("A quoi sert un algorithme de classification?")
	fmt.Println("1. Deviner le prochaion nombre d'une suite")
	fmt.Println("2. Prédire une categorie ou une classe")
	fmt.Println("3. Améliorer la vitesse du processeur")
	if CheckAnswer(2) {
		score++
	}
	// Question 6
	fmt.Println("Qu'est ce qu'un Dataset?")
	fmt.Println("1. Un type de réseau informatique")
	fmt.Println("2. Un ensemble organisé de données")
	fmt.Println("3. Un modèle d'IA pré-entrainé")
	if CheckAnswer(2) {
		score++
	}
	// Question 7
	fmt.Println(" Que fait un réseau de neruones artificels?")
	fmt.Println("1. Imite de manière simplifiée le fonctionnement du cerveay humain")
	fmt.Println("2. Stocke des fichiers dans le cloud")
	fmt.Println("3. Crée des circuits électroniques")
	if CheckAnswer(1) {
		score++
	}
	// Question 8
	fmt.Println("Quand parle-t-on d'overfitting (surapprentissage)?")
	fmt.Println("1. Le modèle oublie toutes les données")
	fmt.Println("2. Le modèle devient trop simple")
	fmt.Println("3. Le modèle apprend trop bien les données d'entrainement et généralise mal")
	if CheckAnswer(3) {
		score++
	}
	// Question 9
	fmt.Println("Quel language est le plus répandu pour l'IA?")
	fmt.Println("1. html")
	fmt.Println("2. Bash")
	fmt.Println("3. Python")
	if CheckAnswer(3) {
		score++
	}
	// Question 10
	fmt.Println(" A quoi sert une base de données ?")
	fmt.Println("1. A exécuter du code Go")
	fmt.Println("2. A stocker et organiser des informations")
	fmt.Println("3. A produire des images")
	if CheckAnswer(2) {
	}
	score++
	CalculateScore(score, totalQuestions)
 }
