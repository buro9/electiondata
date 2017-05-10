package electiondata

import "strconv"

func (r *Results) ToCSV() ([][]string, error) {
	records := [][]string{
		{"year", "id", "name", "candidate", "party", "votes", "incumbent", "elected"},
	}

	for _, e := range r.Elections {
		for _, c := range e.Constituencies {
			for _, r := range c.Results {
				records = append(
					records,
					[]string{
						strconv.Itoa(e.Year),
						c.ID,
						c.Name,
						r.Candidate,
						r.Party,
						strconv.Itoa(r.Votes),
						strconv.FormatBool(r.Incumbent),
						strconv.FormatBool(r.Elected),
					},
				)
			}
		}
	}

	return records, nil
}
