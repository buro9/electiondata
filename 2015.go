package electiondata

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

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
		switch record[PARTY] {
		case "Con":
			res.Party = CON
		case "DUP":
			res.Party = DUP
		case "Green":
			res.Party = GRN
		case "Ind":
			res.Party = IND
		case "Lab":
			res.Party = LAB
		case "LD":
			res.Party = LIB
		case "PC":
			res.Party = PC
		case "SDLP":
			res.Party = SDLP
		case "SF":
			res.Party = SF
		case "SNP":
			res.Party = SNP
		case "Spk":
			res.Party = SPK
		case "UKIP":
			res.Party = UKIP
		case "UUP":
			res.Party = UUP
		default:
			res.Party = record[PARTY]
		}
		votes, err := strconv.Atoi(strings.Replace(record[VOTE], ",", "", -1))
		if err != nil {
			return election, err
		}
		res.Votes = votes
		cur.Results = append(cur.Results, res)
	}
	election.Constituencies = append(election.Constituencies, cur)

	for _, c := range election.Constituencies {
		for _, r := range c.Results {
			if r.Elected {
				// winning parties sanitisation test
				switch r.Party {
				case CON:
				case DUP:
				case GRN:
				case IND:
				case LAB:
				case LIB:
				case PC:
				case SDLP:
				case SF:
				case SNP:
				case SPK:
				case UKIP:
				case UUP:
				default:
					return election,
						fmt.Errorf("%s not recognised for %s\n", r.Party, c.Name)
				}
			}
		}
	}

	return election, nil
}
