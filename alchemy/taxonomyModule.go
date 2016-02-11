package alchemy

import (
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
	"strconv"
)

// TODO: test this
// TODO: fix the neo4j api because its all messed up
// TODO: think about if we want this kind of stuff in the pipeline
//       or in here

// TaxonomyEngine ranks articles by taxonomy, a high level grouping. An
// example taxonomy might be "government and politics/elections". This
// is a fairly coarse method.
type TaxonomyEngine struct {
	// Use the neo4j graph database for relation between taxonomies.
	// The graph is undirected, where nodes hold taxonomies and edges
	// are relationships between two taxonomies. Edge weight is relation
	// strength between each node.
	db    *neoism.Database // db driver for neo4j graph db
	cache Neo4jCache       // simple cache to store neo requests
}

// ScoreTaxonomy scores the main taxonomy against the related taxonomy.
func (engine *TaxonomyEngine) ScoreTaxonomy(main, related []Taxonomy) (float64, error) {

	totalScore := 0.0

	// loop through each taxonomy
	for _, mainTax := range main {
		for _, relatedTax := range related {

			if mainTax.Label == relatedTax.Label {
				// give equivalent taxonomies a high score
				totalScore += 5.0 * float64(mainTax.Score*relatedTax.Score)
				continue
			}

			// fetch from DB
			score, err := engine.getRelationStrength(mainTax.Label, relatedTax.Label)
			if err != nil {
				// TODO: handle this error better
				return 0.0, err
			}

			// factor in the taxonomy/article strengths
			score *= float64(mainTax.Score * relatedTax.Score)
			totalScore += score
		}
	}

	return totalScore, nil
}

// getRelationStrength of two taxonomies from the neo4j graph db.
func (engine *TaxonomyEngine) getRelationStrength(main, related string) (float64, error) {
	// check if this request is cached
	// saves sending a net request
	// is now only a minor time saver, will get bigger with large data sets
	if v, isCached := engine.cache.Get(main, related); isCached {
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
	url := engine.db.HrefCypher // get URL through db driver
	// send request
	resp, err := engine.db.Session.Post(url, &payload, &result, &ne)
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
		engine.cache.Add(main, related, float64(score))

		return float64(score), nil
	}

	// if results empty then there is no relation
	engine.cache.Add(main, related, 0)

	return 0, nil
}

// Setup connects to the graph server and sets up the cache.
func (m *TaxonomyEngine) Setup() {
	db, err := neoism.Connect("http://neo4j:root@localhost:7474/db/data/")
	if err != nil {
		fmt.Println("error in setup, db is:", db)
		panic(err)
	}

	m.db = db

	m.cache.Setup()
}
