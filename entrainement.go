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

	fmt.Println("1) Systèmes d’exploitation")
	fmt.Println("Un système d’exploitation (OS) est le logiciel qui permet à ton ordinateur de fonctionner.")
	fmt.Println("Il gère la mémoire, les fichiers, les programmes et les périphériques (clavier, souris, réseau...).")
	fmt.Println("Parmi les systèmes d’exploitation les plus connus, on trouve :")
	fmt.Println("- Windows : très répandu sur les PC personnels et en entreprise ;")
	fmt.Println("- macOS : le système d’Apple, utilisé sur les Mac ;")
	fmt.Println("- Linux : un système libre, très utilisé sur les serveurs, dans le cloud, et par les développeurs.\n")

	fmt.Println("2) Langages de programmation")
	fmt.Println("Un langage de programmation permet de donner des instructions à un ordinateur.")
	fmt.Println("Chaque langage est adapté à certains usages :")
	fmt.Println("- Go : très efficace pour les applications réseau, le cloud, les services web ;")
	fmt.Println("- Python : très utilisé pour l’automatisation, la data, l’IA ;")
	fmt.Println("- JavaScript : le langage du web côté navigateur ;")
	fmt.Println("- Java / C# : souvent utilisés pour les applications d’entreprise.\n")

	fmt.Println("3) Matériel : CPU, RAM, stockage")
	fmt.Println("Le CPU (Central Processing Unit) est le 'cerveau' de l’ordinateur. Il exécute les instructions.")
	fmt.Println("La RAM est une mémoire rapide et temporaire : elle stocke les données des programmes en cours d’exécution.")
	fmt.Println("Le stockage (HDD ou SSD) sert à conserver les fichiers même quand l’ordinateur est éteint.")
	fmt.Println("Un SSD est plus rapide qu’un disque dur (HDD) et rend l’ordinateur beaucoup plus réactif.\n")

	fmt.Println("4) Réseau et Internet")
	fmt.Println("Sur un réseau, chaque appareil communique via une adresse IP.")
	fmt.Println("Des protocoles comme HTTP et HTTPS sont utilisés pour afficher les pages web.")
	fmt.Println("HTTPS est la version sécurisée de HTTP : les données sont chiffrées entre ton navigateur et le serveur.\n")

	fmt.Println("Cette section te donne une base pour mieux comprendre les questions du quiz Informatique.\n")
}

func EntrainementCyber() {
	fmt.Println("\n=== Entraînement – Cybersécurité ===\n")

	fmt.Println("1) Qu’est-ce que la cybersécurité ?")
	fmt.Println("La cybersécurité regroupe l’ensemble des pratiques, outils et méthodes qui protègent les systèmes")
	fmt.Println("informatiques, les réseaux et les données contre les attaques et les accès non autorisés.\n")

	fmt.Println("2) Menaces courantes")
	fmt.Println("- Phishing : un faux e-mail ou message qui imite un service de confiance (banque, réseau social, etc.)")
	fmt.Println("  pour te pousser à cliquer sur un lien ou à donner tes identifiants.")
	fmt.Println("- Malware : logiciel malveillant (virus, cheval de Troie, ransomware) installé sur ta machine.")
	fmt.Println("- Ransomware : un type de malware qui chiffre tes fichiers et demande une rançon pour les récupérer.\n")

	fmt.Println("3) Moyens de protection")
	fmt.Println("- Pare-feu (firewall) : filtre le trafic réseau entrant et sortant pour bloquer ce qui est suspect ;")
	fmt.Println("- Antivirus / antimalware : analyse les fichiers et les programmes pour détecter des comportements malveillants ;")
	fmt.Println("- Mots de passe forts : longs, avec lettres, chiffres et caractères spéciaux, différents pour chaque service ;")
	fmt.Println("- Authentification multifacteur (MFA) : ajoute une étape (code SMS, application, clé physique) en plus du mot de passe.\n")

	fmt.Println("4) Bonnes pratiques au quotidien")
	fmt.Println("- Ne jamais cliquer sur un lien douteux ou une pièce jointe d’un expéditeur inconnu ;")
	fmt.Println("- Vérifier l’adresse du site (URL) avant de saisir un mot de passe, surtout pour la banque ou les services sensibles ;")
	fmt.Println("- Mettre à jour régulièrement son système et ses logiciels ;")
	fmt.Println("- Sauvegarder ses données importantes sur un support externe ou dans le cloud.\n")

	fmt.Println("Comprendre ces notions t’aidera à répondre aux questions sur les attaques, les défenses et les bonnes pratiques.\n")
}

func EntrainementIAData() {
	fmt.Println("\n=== Entraînement – IA / Data ===\n")

	fmt.Println("1) Données : la base de tout")
	fmt.Println("L’IA et la Data reposent sur la collecte et l’analyse de données.")
	fmt.Println("- Données structurées : organisées dans des tableaux (lignes/colonnes), comme dans une base SQL ;")
	fmt.Println("- Données non structurées : textes, images, vidéos, sons, qui ne rentrent pas dans un simple tableau.\n")

	fmt.Println("2) Rôle de SQL")
	fmt.Println("SQL (Structured Query Language) est le langage standard pour interroger les bases de données.")
	fmt.Println("Il permet par exemple de sélectionner certaines lignes, filtrer, trier, faire des statistiques simples.")
	fmt.Println("Savoir lire et écrire des requêtes SQL est une compétence clé dans les métiers de la data.\n")

	fmt.Println("3) IA, Machine Learning et Deep Learning")
	fmt.Println("L’Intelligence Artificielle (IA) regroupe différentes techniques permettant à une machine d’imiter")
	fmt.Println("certaines capacités humaines (reconnaissance d’image, compréhension du langage, prédiction...).")
	fmt.Println("- Machine Learning : l’ordinateur apprend à partir d’exemples (données d’entraînement).")
	fmt.Println("- Deep Learning : sous-domaine du Machine Learning qui utilise des réseaux de neurones profonds,")
	fmt.Println("  très efficaces pour l’image, le son ou le texte.\n")

	fmt.Println("4) Outils fréquents")
	fmt.Println("- Python : langage principal pour la data et l’IA ;")
	fmt.Println("- Bibliothèques : Pandas (manipulation de données), Scikit-Learn (Machine Learning),")
	fmt.Println("  TensorFlow / PyTorch (Deep Learning).\n")

	fmt.Println("Cette section t’aide à comprendre les mots-clés que tu verras dans le quiz IA / Data.\n")
}
