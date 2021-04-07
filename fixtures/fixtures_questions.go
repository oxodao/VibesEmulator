package fixtures

import (
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
)

func generateQuestions(prv *services.Provider) {
	questions := []models.Question{
		{
			Text: "Si vous deviez faire un boulot pendant 6 mois, lequel serait-ce ?",
			Answers: []models.Answer{
				{ Text: "Paysagiste" },
				{ Text: "Service civique" },
				{ Text: "Infermier" },
			},
		},
		{
			Text: "Accepterais-tu de te faire tatouer le logo d'une marque sur le front pour un million d'euros ?",
			Answers: []models.Answer{
				{ Text: "Oui" },
				{ Text: "Non" },
			},
		},
		{
			Text: "Tu es sur un roller coaster et tu te rends compte que tu va vomir. À ta gauche il y a ta mère, à ta droite ton conjoint. Sur qui vomis-tu ?",
			Answers: []models.Answer{
				{ Text: "Ma mère" },
				{ Text: "Mon conjoint" },
			},
		},
		{
			Text: "Les murs de ton appartement sont... ?",
			Answers: []models.Answer{
				{ Text: "Tout blancs" },
				{ Text: "Colorés" },
				{ Text: "Avec des motifs" },
			},
		},
		{
			Text: "Ton rendez-vous à l'aveugle a mal pris quelque chose que tu as dit et te jette son verre à la figure. Que fais-tu ?",
			Answers: []models.Answer{
				{ Text: "La même chose !" },
				{ Text: "Partir !" },
				{ Text: "Un autre verre, s'il vous plait !" },
			},
		},
		{
			Text: "Tu amarres ton bateau sur une magnifique île grècque et tu te rends compte qu'il s'agit d'une plage nudiste. Que fais-tu ?",
			Answers: []models.Answer{
				{ Text: "Demi tour !" },
				{ Text: "Un maillot de bain, quel maillot de bain ?!" },
				{ Text: "Un autre verre, s'il vous plait !" },
			},
		},
		{
			Text: "Tu passes à côté d'une vitrine de magasin et tu vois ton reflet. Que penses-tu ?",
			Answers: []models.Answer{
				{ Text: "Wow je suis beau !" },
				{ Text: "Hmm, pas mal" },
				{ Text: "Oh mon dieu, c'est vraiment moi ??" },
			},
		},
		{
			Text: "Star Trek est...",
			Answers: []models.Answer{
				{ Text: "... un classique" },
				{ Text: "... pour les nerds" },
				{ Text: "... on s'en fout !" },
			},
		},
		{
			Text: "Tu sens que ton meilleur ami et toi vous éloignez. Qu'est-ce que tu fais ?",
			Answers: []models.Answer{
				{ Text: "Je lui en parle" },
				{ Text: "Je laisse les choses se faire" },
			},
		},
		{
			Text: "Est-ce que tu as des tatouages",
			Answers: []models.Answer{
				{ Text: "Oui" },
				{ Text: "Non" },
			},
		},
		{
			Text: "Il y a un cirque itinérant dans votre ville et on te demande de donner de l'argent pour nourrir les animaux. Que fais-tu ?",
			Answers: []models.Answer{
				{ Text: "Bien sur !" },
				{ Text: "Non, c'est leur problème !" },
			},
		},
		{
			Text: "Est-ce que tu accueillerais une nuit un ancien camarade de classe ?",
			Answers: []models.Answer{
				{ Text: "Bien sur, pourquoi pas ?" },
				{ Text: "Non, ça serait suspect" },
			},
		},
		{
			Text: "Comment agis-tu avec quelqu'un qui te plait ?",
			Answers: []models.Answer{
				{ Text: "Je prends l'initiative" },
				{ Text: "Je n'arrive pas à dire un mot" },
				{ Text: "Je me force à faire l'effort" },
			},
		},
		{
			Text: "Où regardes-tu le dernier blockbuster ?",
			Answers: []models.Answer{
				{ Text: "Au cinéma" },
				{ Text: "A la maison" },
				{ Text: "Je ne le regarde pas" },
			},
		},
		{
			Text: "Quelle est ta position favorite pour dormir ?",
			Answers: []models.Answer{
				{ Text: "Sur le ventre" },
				{ Text: "Sur le dos" },
				{ Text: "Sur le côté" },
			},
		},
	}

	for _, question := range questions {
		err := prv.Dal.Questions.InsertQuestion(&question)
		if err != nil {
			panic(err)
		}
	}

}
