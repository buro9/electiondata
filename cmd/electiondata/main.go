package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/buro9/electiondata"
)

func main() {
	electiondata.RegisterFlags()
	flag.Parse()

	if (electiondata.Raw2010 == nil || *electiondata.Raw2010 == "") &&
		(electiondata.Raw2015 == nil || *electiondata.Raw2015 == "") {
		fmt.Print("One of 2010 or 2015 must be supplied:\n\n")
		flag.Usage()
		os.Exit(1)
	}

	var results electiondata.Results

	if *electiondata.Raw2010 != "" {
		election, err := electiondata.Parse2010(*electiondata.Raw2010)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		cons, cands, votes := election.Totals()
		fmt.Printf(
			"2010: %d votes across %d candidates in %d constituencies\n",
			votes,
			cands,
			cons,
		)
		results.Elections = append(results.Elections, election)
	}

	if *electiondata.Raw2015 != "" {
		election, err := electiondata.Parse2015(*electiondata.Raw2015)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		cons, cands, votes := election.Totals()
		fmt.Printf(
			"2015: %d votes across %d candidates in %d constituencies\n",
			votes,
			cands,
			cons,
		)
		results.Elections = append(results.Elections, election)
	}

	b, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	file, err := os.Create("out/elections.json")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
