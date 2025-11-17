package main

import "fmt"

func SectionEntrainement() {
	for {
		fmt.Println("=== Section Entraînement ===")
		fmt.Println("1. Informatique (bases générales)")
		fmt.Println("2. Cybersécurité")
		fmt.Println("3. IA / Data")
		fmt.Println("0. Retour au menu")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			EntrainementInformatique()
		case 2:
			EntrainementCyber()
		case 3:
			EntrainementIAData()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func EntrainementInformatique() {
	fmt.Println("\n=== Entraînement – Informatique générale ===\n")

	fmt.Println("1) Linux : le système d’exploitation libre le plus utilisé")
	fmt.Println("Linux est un système d’exploitation open-source. Il est très utilisé :")
	fmt.Println("- dans les serveurs,")
	fmt.Println("- dans les infrastructures cloud,")
	fmt.Println("- par les développeurs.")
	fmt.Println("Il est gratuit, modifiable et réputé pour sa stabilité.\n")

	fmt.Println("2) Langages compilés vs interprétés")
	fmt.Println("Un langage compilé (comme Go ou C) est traduit en langage machine AVANT d’être exécuté.")
	fmt.Println("Résultat : il est plus rapide et plus performant.")
	fmt.Println("Les langages interprétés (Python, JavaScript) sont exécutés ligne par ligne.\n")

	fmt.Println("3) HTML : HyperText Markup Language")
	fmt.Println("HTML n'est pas un langage de programmation mais un langage de structure.")
	fmt.Println("Il sert à organiser le contenu des pages web : titres, paragraphes, images, liens, etc.\n")

	fmt.Println("4) Java : langage principal pour Android")
	fmt.Println("Android utilise principalement Java et Kotlin.")
	fmt.Println("Java est très utilisé dans les applications mobiles et d'entreprise.\n")

	fmt.Println("5) CPU : Central Processing Unit")
	fmt.Println("Le CPU est le processeur, le cerveau de l’ordinateur.")
	fmt.Println("Il exécute les instructions des programmes.\n")

	fmt.Println("6) HTTPS : version sécurisée du HTTP")
	fmt.Println("HTTPS chiffre les données entre ton navigateur et un site web.")
	fmt.Println("Il empêche l’interception et la modification des informations.\n")

	fmt.Println("7) RAM : mémoire volatile")
	fmt.Println("La RAM perd son contenu lorsqu’on éteint l’ordinateur.")
	fmt.Println("Elle est utilisée pour stocker temporairement les données en cours d’utilisation.\n")

	fmt.Println("8) SQL : Structured Query Language")
	fmt.Println("SQL est le langage standard pour interagir avec une base de données :")
	fmt.Println("- sélectionner,")
	fmt.Println("- filtrer,")
	fmt.Println("- trier,")
	fmt.Println("- modifier des données.\n")

	fmt.Println("9) Structure LIFO : la pile (Stack)")
	fmt.Println("LIFO = Last In, First Out.")
	fmt.Println("Le dernier élément ajouté est le premier retiré.")
	fmt.Println("Exemple : pile d’assiettes.\n")

	fmt.Println("10) Pare-feu : filtrage du trafic réseau")
	fmt.Println("Un pare-feu (firewall) contrôle les communications entrantes et sortantes.")
	fmt.Println("Il bloque les connexions dangereuses et protège le réseau.\n")
}

func EntrainementCyber() {
	fmt.Println("\n=== Entraînement – Cybersécurité ===\n")

	fmt.Println("1) Mot de passe sécurisé : mélange lettres/chiffres/symboles")
	fmt.Println("Un mot de passe fort doit être long et complexe.")
	fmt.Println("Utiliser sa date de naissance ou un mot simple est très risqué.\n")

	fmt.Println("2) Phishing : faux e-mail frauduleux")
	fmt.Println("Le phishing vise à tromper l’utilisateur pour voler :")
	fmt.Println("- mots de passe,")
	fmt.Println("- numéros bancaires,")
	fmt.Println("- identités.")
	fmt.Println("On imite un e-mail officiel pour pousser à cliquer.\n")

	fmt.Println("3) Antivirus : détecte et supprime les malwares")
	fmt.Println("Il surveille les fichiers, les programmes et le comportement du système.")
	fmt.Println("Il empêche l’installation de logiciels malveillants.\n")

	fmt.Println("4) HTTPS : protocole sécurisé pour les sites")
	fmt.Println("HTTPS chiffre les données échangées.")
	fmt.Println("Toujours vérifier le cadenas dans le navigateur.\n")

	fmt.Println("5) Clé USB non chiffrée : très vulnérable")
	fmt.Println("Une clé USB peut être perdue facilement.")
	fmt.Println("Si elle n’est pas chiffrée, n’importe qui peut lire son contenu.\n")

	fmt.Println("6) Ransomware : logiciel qui chiffre vos données")
	fmt.Println("Le pirate demande ensuite une rançon pour les restituer.")
	fmt.Println("Attaque très courante sur les entreprises.\n")

	fmt.Println("7) Authentification à deux facteurs (2FA / MFA)")
	fmt.Println("Le 2FA ajoute une confirmation supplémentaire :")
	fmt.Println("- SMS,")
	fmt.Println("- application d'authentification,")
	fmt.Println("- clé physique.")
	fmt.Println("Même si le mot de passe est volé, le compte reste protégé.\n")

	fmt.Println("8) Cliquer sur un lien inconnu : comportement dangereux")
	fmt.Println("C’est une des principales origines d’infections : phishing, virus, trojans.\n")

	fmt.Println("9) Wi-Fi public non sécurisé : interception de données")
	fmt.Println("Un pirate peut espionner :")
	fmt.Println("- mots de passe,")
	fmt.Println("- messages,")
	fmt.Println("- pages visitées.")
	fmt.Println("Toujours utiliser un VPN.\n")

	fmt.Println("10) Logiciel de chiffrement : rend un fichier illisible sans clé")
	fmt.Println("Les outils de chiffrement protègent les données sensibles.")
	fmt.Println("Même si un fichier est volé, il reste inutilisable.\n")
}

func EntrainementIAData() {
	fmt.Println("\n=== Entraînement – IA / Data ===\n")

	fmt.Println("1) Modèle d’IA : programme qui apprend à partir de données")
	fmt.Println("Un modèle analyse des données pour reconnaître des schémas et faire des prédictions.\n")

	fmt.Println("2) Machine Learning : apprentissage automatique")
	fmt.Println("Un programme apprend grâce à des exemples plutôt que par des règles codées.\n")

	fmt.Println("3) Big Data : analyser de très grands volumes de données")
	fmt.Println("Objectif : extraire de l’information utile depuis d’énormes quantités de données.\n")

	fmt.Println("4) Feature : variable d’entrée du modèle")
	fmt.Println("Une feature peut être :")
	fmt.Println("- l’âge d’un client,")
	fmt.Println("- une température,")
	fmt.Println("- un mot dans un texte.")
	fmt.Println("Elle sert au modèle pour prédire quelque chose.\n")

	fmt.Println("5) Classification : prédire une catégorie")
	fmt.Println("Exemples classiques :")
	fmt.Println("- spam ou non spam,")
	fmt.Println("- chien / chat,")
	fmt.Println("- client qui part / reste.\n")

	fmt.Println("6) Dataset : ensemble organisé de données")
	fmt.Println("Un dataset contient :")
	fmt.Println("- des données d’entrée,")
	fmt.Println("- des réponses (pour l’entraînement),")
	fmt.Println("- parfois des métadonnées.\n")

	fmt.Println("7) Réseau de neurones : imitation simplifiée du cerveau humain")
	fmt.Println("Il apprend en ajustant des connexions appelées pondérations.")
	fmt.Println("Très utilisé pour la vision, le son, et le langage.\n")

	fmt.Println("8) Overfitting : le modèle apprend trop bien les données d’entraînement")
	fmt.Println("Il mémorise au lieu de généraliser.")
	fmt.Println("Il devient mauvais sur de nouvelles données.\n")

	fmt.Println("9) Python : langage n°1 de l’IA")
	fmt.Println("Il dispose de bibliothèques puissantes :")
	fmt.Println("- NumPy,")
	fmt.Println("- Pandas,")
	fmt.Println("- TensorFlow,")
	fmt.Println("- PyTorch.\n")

	fmt.Println("10) Base de données : stocker et organiser des informations")
	fmt.Println("Les bases permettent :")
	fmt.Println("- de sauvegarder,")
	fmt.Println("- de rechercher,")
	fmt.Println("- d’analyser des données efficacement.\n")
}
