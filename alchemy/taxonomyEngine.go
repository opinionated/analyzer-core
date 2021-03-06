package alchemy

import (
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
	"strconv"
)

// TODO: docs, testing
type TaxonomyEngine struct {

	// Use the neo4j graph to score taxonomies against eachother
	db    *neoism.Database
	cache Neo4jCache
}

// Requests relation strength between two taxonomies from the db.
func (m *TaxonomyEngine) GetRelationStrength(main, related string) (float64, error) {
	// check if this request is cached
	// saves sending a net request
	// is now only a minor time saver, will get bigger with large data sets
	if v, isCached := m.cache.Get(main, related); isCached {
		return v, nil
	}

	cq := neoism.CypherQuery{
		Statement: `MATCH (a)-[r]-(b) 
		WHERE a.name={main} AND b.name={related} 
		RETURN r.cost`,
		Parameters: neoism.Props{"main": main, "related": related},
		Result:     []struct{}{},
	}

	// neoism seems to be broken so build request manually
	// these are the structures neoism uses internally
	type cypherRequest struct {
		Query      string                 `json:"query"`
		Parameters map[string]interface{} `json:"params"`
	}

	type cypherResult struct {
		Columns []string
		Data    [][]*json.RawMessage
	}

	result := cypherResult{}
	payload := cypherRequest{
		Query:      cq.Statement,
		Parameters: cq.Parameters,
	}

	ne := neoism.NeoError{}
	url := m.db.HrefCypher // get URL through db driver
	// send request
	resp, err := m.db.Session.Post(url, &payload, &result, &ne)
	if err != nil {
		fmt.Println("resp is:", resp)
		panic(err)
	}

	// if the results are filled, parse them
	// TODO: do more legitimate parsing, this is probably unsafe
	if len(result.Data) > 0 {
		// read in the first (and only) element
		rawMessage := result.Data[0][0]
		r, err := rawMessage.MarshalJSON()
		if err != nil {
			panic(err)
		}

		// be lazy and convert it to string then atoi the string
		s := string(r[:])
		score, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		// add this query to the cache (wouldn't hit here if it was in cache)
		m.cache.Add(main, related, float64(score))

		return float64(score), nil
	}

	// if results empty then there is no relation
	m.cache.Add(main, related, 0)

	return 0, nil
}

func (m *TaxonomyEngine) Start() {

	db, err := neoism.Connect("http://neo4j:root@localhost:7474/db/data/")
	if err != nil {
		fmt.Println("error in setup, db is:", db)
		panic(err)
	}

	m.db = db

	m.cache.Setup()
}
