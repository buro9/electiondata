package electiondata

import "encoding/json"

func Parse2016(file string) (Referendum, error) {
	ref := Referendum{
		Year:     2016,
		Question: "Should the United Kingdom remain a member of the European Union or leave the European Union?",
	}

	b, err := readFile(file)
	if err != nil {
		return ref, err
	}

	//Areas

	var areas []struct {
		ID      string `json:"areaCode"`
		Name    string `json:"areaName"`
		Answers []struct {
			Label string `json:"answerLabel"`
			Votes int    `json:"answerValue"`
		} `json:"answers"`
	}
	err = json.Unmarshal(b, &areas)
	if err != nil {
		return ref, err
	}

	for _, area := range areas {
		var ca ConstituencyAnswer
		ca.ID = area.ID
		ca.Name = area.Name

		for _, answer := range area.Answers {
			ca.Answers = append(
				ca.Answers,
				Answer{
					Answer: answer.Label,
					Votes:  answer.Votes,
				},
			)
		}

		ref.ConstituencyAnswers = append(ref.ConstituencyAnswers, ca)
	}

	return ref, nil
}
