package main

import "fmt"

func StartQuizCyber() {
	fmt.Println("―――――Quiz Cybersecurité ―――――")
	
	
	totalQuestions := 10
	score := 0

	// Question 1
	fmt.Println("Quel est le meilleur moyen de créer un mot de passe sécurisé ?")
	fmt.Println("1. Utiliser sa date de naissance ")
	fmt.Println("2. Utiliser un mot unique facile à retenir")
	fmt.Println("3. Combiner lettres,chiffres et symboles")
	if CheckAnswer(3) {
		score++
	}
	// Question 2
	fmt.Println("Quel type d'attaque consiste a tromper un utilisateur via un faux e-mail ?")
	fmt.Println("1. Phishing")
	fmt.Println("2. DDoS")
	fmt.Println("3. Ransomware")
	if CheckAnswer(1) {
		score++
	}
	// Question 3
	fmt.Println("Quel est le role d'un antivirus ?")
	fmt.Println("1. Booster les performances du PC")
	fmt.Println("2. Détecter et supprimer les logiciels malveillants")
	fmt.Println("3. Bloquer les e-mails")
	if CheckAnswer(2) {
		score++
	}
	// Question 4
	fmt.Println("Quel est protocole est le plus sécurisé pour naviguer sur un site web  ?")
	fmt.Println("1. HTTP")
	fmt.Println("2. FTP")
	fmt.Println("3. HTTPS")
	if CheckAnswer(3) {
		score++
	}
	// Question 5
	fmt.Println("Quel support est le plus vulnerable a la perte ou au vol de données?")
	fmt.Println("1. Clé USB non chiffrée")
	fmt.Println("2. Disque dur externe chiffré")
	fmt.Println("3. Cloud sécurisé")
	if CheckAnswer(1) {
		score++
	}
	// Question 6
	fmt.Println("Quel terme décrit un logiciel qui demande une rancon pour lib&rer les données?")
	fmt.Println("1. Spyware")
	fmt.Println("2. Ransomware")
	fmt.Println("3. Adware")
	if CheckAnswer(2) {
		score++
	}
	// Question 7
	fmt.Println(" Quelle action renforce le plus la sécurité d'un compte en ligne?")
	fmt.Println("1. Utiliser le meme mot de passe partout")
	fmt.Println("2. Activer l'authentification a deux facteur")
	fmt.Println("3. Partager son mot de passe avec un collègue de confiance")
	if CheckAnswer(2) {
		score++
	}
	// Question 8
	fmt.Println("Quel comportement est risqué en cybersecurité?")
	fmt.Println("1. Mettre régulièrement a jour ses logiciels")
	fmt.Println("2. Cliquer sur un lien recu par un inconnu")
	fmt.Println("3. Sauvegarder ses données")
	if CheckAnswer(2) {
		score++
	}
	// Question 9
	fmt.Println(" l est le principal risque lié au WI-Fi public non sécurisé?")
	fmt.Println("1. Une consommation excessive de batterie")
	fmt.Println("2. L'interception de vos données")
	fmt.Println("3. Une connexion plus lente")
	if CheckAnswer(2) {
		score++
	}
	// Question 10
	fmt.Println(" Quel outil permet de chiffrer des fichiers pour les rendres illisible sans clé ?")
	fmt.Println("1. Un logiciel de chiffrement")
	fmt.Println("2. Un lecteur multimedia")
	fmt.Println("3. Un navigateur web")
	if CheckAnswer(1) {
	}
	score++
	CalculateScore(score, totalQuestions)
 }

