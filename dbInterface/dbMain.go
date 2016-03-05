package dbInterface

import (
	"encoding/json"
	"database/sql"
	_ "github.com/satori/go.uuid"
	_ "github.com/go-cq/cq"
)

// keyQuery and taxQuery store the same information, differentiation is so search-field can be implicit
type KeyQuery struct {
	Text      string
	Sentiment string
	Threshold float32
}

type TaxQuery struct {
	Name      string
	Sentiment string
	Threshold float32
}

type DBKeyword struct {
	Text      string  `json:"text"`
	Sentiment string  `json:"sentiment,omitempty"`
	Relevance float32 `json:"relevance,omitempty"`
}

type DBTaxonomy struct {
	Label     string  `json:"label"`
	Sentiment string  `json:"sentiment"`
	Score     float32 `json:"score,omitempty"`
}

type ArticleInfo struct {
	ID         int          `json:"ID,omitempty"`
	Author     string       `json:"Author,omitempty"`
	Body       string       `json:"BodyFilename"`
	Keywords   []DBKeyword  `json:"Keywords,omitempty"`
	Taxonomies []DBTaxonomy `json:"Taxonomies,omitempty"`
}

//Don't need the driver for this, just using normal SQL calls
func store(info ArticleInfo) error {
	//create uuid for article
	info.id = uuid.NewV4()

	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
    if err != nil {
        return err //oh no!
    }
    defer db.Close()

	//create query, use prepared statements
    stmt, err := db.Prepare(`
		UNWIND {{0}} as article
		MERGE (a:Article {id:article.ID, body:article.Body, author:article.Body})
		FOREACH (kw IN article.Keywords |
		   MERGE (k:Keyword {text:kw.Name})
		   MERGE (k)-[:KEYWORD_OF {sentiment:kw.Sentiment, relevance:kw.Relevance}]->(a)
		   MERGE (t:Taxonomy {label:kw.Label})
		   MERGE (t)-[:TAXONOMY_OF {sentiment:kw.Sentiment, score:kw.Score}]->(a)
		)
	`)
    if err != nil {
        return err //oops
    }
	defer stmt.close()

	//pass in the json object to be unwound and execute the statement
	_, err := stmt.Exec(ArticleInfo)

	//no error = nil
	return err
}
