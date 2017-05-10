package electiondata

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

// Parse2015 parses the results.csv for 2015
func Parse2015(file string) (Election, error) {
	election := Election{
		Year: 2015,
	}

	b, err := readFile(file)
	if err != nil {
		return election, err
	}

	r := csv.NewReader(bytes.NewReader(b))
	records, err := r.ReadAll()
	if err != nil {
		return election, err
	}

	const (
		FORENAME = iota
		SURNAME
		DESCRIPTION
		SEAT
		PANO
		VOTE
		PERCENT
		CHANGE
		FOO1
		INCUMBENT
		FOO2
		SEAT_ID
		REGION_ID
		COUNTY
		REGION
		COUNTRY
		SEAT_TYPE
		PARTY_NAME
		PARTY
	)

	var cur Constituency
	for ii, record := range records {
		if ii == 0 {
			// header row
			continue
		}

		var elected bool
		if ii == 1 {
			cur.ID = record[SEAT_ID]
			cur.Name = record[SEAT]
			elected = true
		} else if cur.Name != record[SEAT] {
			// add the last one, and work on the new
			election.Constituencies = append(election.Constituencies, cur)
			cur = Constituency{}
			cur.ID = record[SEAT_ID]
			cur.Name = record[SEAT]
			elected = true
		}

		var res Result
		res.Candidate = fmt.Sprintf("%s, %s", record[SURNAME], record[FORENAME])
		res.Elected = elected
		res.Incumbent = (record[INCUMBENT] == "MP")
		res.Party = record[PARTY]
		res.SanitizeParty()
		votes, err := strconv.Atoi(strings.Replace(record[VOTE], ",", "", -1))
		if err != nil {
			return election, err
		}
		res.Votes = votes
		cur.Results = append(cur.Results, res)
	}
	election.Constituencies = append(election.Constituencies, cur)

	return election, nil
}
