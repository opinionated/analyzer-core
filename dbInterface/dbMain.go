package dbInterface

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-cq/cq"
	_ "github.com/satori/go.uuid"
)

// keyQuery and taxQuery store the same information, differentiation is so search-field can be implicit
type KeyQuery struct {
	Text      string
	Sentiment string //ignoring for now
	Threshold float32
}

type TaxQuery struct {
	Name      string
	Sentiment string //ignoring for now
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

//Signifies error at certain id
type IDError struct {
	uuid    string
	message string
}

func (e *IDError) Error() string {
	return fmt.Sprintf("%s - %s", e.uuid, e.message)
}

//Initialize ?figure out what goes here
func Init() error {

}

//Don't need the driver for this, just using normal SQL calls
func Store(info ArticleInfo) error {
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
	_, err := stmt.Exec(info)

	//no error = nil
	return err
}

/**
 * @param id UUID, article id
 * @return error, in case of error
 * Remove article w/ id from database
 */
func Remove(id UUID) error {

	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		return err //Could not open!
	}
	defer db.Close()

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as uuid
		MATCH (articleID {id: uuid })
		DELETE articleID
	`)
	if err != nil {
		return err //Wat r u tryina remove dawg?
	}
	defer stmt.close()

	//Unwind json object and execute query
	_, err := stmt.Exec(id)
	//JUST DO IT^^^

	//no error = nil
	return err

}

/**
 * DB reset, for testing purposes !!!DONT USE UNLESS YOU'RE SERIOUSLY ARE SURE ABOUT THIS!!!
 */
func RemoveAll() error {
	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		return err //Could not open!
	}
	defer db.Close()

	//Create query
	stmt, err := db.Prepare(`
		MATCH (n) DETACH
		DELETE n
	`)
	if err != nil {
		return err //Wat r u tryina remove dawg?
	}
	defer stmt.close()

	//Execute query
	_, err := stmt.Exec()
	//JUST DO IT^^^

	//no error = nil
	return err

}

/**
 * @param id UUID, article id
 * @return ArticleInfo, first article with id {id}.
 * @return error, in case of error, or multiple articles listed under id
 * Gets article w/ id from database
 */
func Get(id UUID) (ArticleInfo, error) {

	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		return _, err //Could not open!
	}
	defer db.Close()

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as uuid
		MATCH (article {ID: uuid})
		RETURN article
	`)
	if err != nil {
		return _, err //Wat r u tryina git dawg?
	}
	defer stmt.close()

	//Execute query
	ret, err := stmt.Query(id)
	//JUST DO IT^^^
	if err != nil {
		return _, err //Something went wrong!!!
	}
	defer ret.close()

	//pull all articles and decode. Return the first (SHOULD NOT BE MORE THAN 1, return error otherwise)
	var article ArticleInfo
	var articleJSON []byte
	err := rows.Scan(&(articleJSON))
	if err != nil {
		return _, err //this one is weird
	}
	//DECODE json object
	err := json.Unmarshal(articleJSON, &article)
	if err != nil {
		return _, err
	}

	//Should not be any more rows, if there are return IDError on id
	for rows.Next() {
		return _, &IDError{id.String(), "Multiple entries in database for this ID!!"}
	}

	//no error = nil
	return article, err

}

/**
 * @param keyword Keyword, keyword
 * @return []UUID, set of articles ids with keywords
 * @return error, in case of error
 * Gets article ids w/ keyword from database
 */
func SearchByKeyword(keyword DBKeyword) ([]UUID, error) {
	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		return _, err //Could not open!
	}
	defer db.Close()

	if keyword.Threshold := nil {
		keyword.Threshold = 0
	}

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as kw
		MATCH (article)-[kw.Threshold]->(kw.Name)
		RETURN article.ID
	`)
	if err != nil {
		return _, err //Wat r u tryina git dawg?
	}
	defer stmt.close()

	//Execute query
	ret, err := stmt.Query(id)
	//JUST DO IT^^^
	if err != nil {
		return _, err //Something went wrong!!!
	}
	defer ret.close()

	//pull all articles and decode. Return the slice with all of them in it
	var aID UUID
	var idJSON []byte
	var set []UUID
	for rows.Next() {
		err := rows.Scan(&(idJSON)) //scan
		if err != nil {
			return _, err //this one is weird
		}
		err := json.Unmarshal(idJSON, &aID) //decode
		if err != nil {
			return _, err //more ways it can be weird
		}
		set = append(set, aID) //append
	}

	//no error = nil
	return set, err
}

/**
 * @param taxonomy DBTaxonomy, taxonomy
 * @return []UUID, set of articles with keywords
 * @return error, in case of error
 * Gets article ids w/ keyword from database
 */
func SearchByTaxonomy(taxonomy DBTaxonomy) ([]UUID, error) {
	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		return _, err //Could not open!
	}
	defer db.Close()

	if taxonomy.Threshold := nil {
		taxonomy.Threshold = 0
	}

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as tx
		MATCH (article)-[tx.Threshold]->(tx.Name)
		RETURN article.ID
	`)
	if err != nil {
		return _, err //Wat r u tryina git dawg?
	}
	defer stmt.close()

	//Execute query
	ret, err := stmt.Query(id)
	//JUST DO IT^^^
	if err != nil {
		return _, err //Something went wrong!!!
	}
	defer ret.close()

	//pull all articles and decode. Return the slice with all of them in it
	var aID UUID
	var idJSON []byte
	var set []UUID
	for rows.Next() {
		err := rows.Scan(&(idJSON)) //scan
		if err != nil {
			return _, err //this one is weird
		}
		err := json.Unmarshal(idJSON, &aID) //decode
		if err != nil {
			return _, err //more ways it can be weird
		}
		set = append(set, aID) //append
	}

	//no error = nil
	return set, err
}
