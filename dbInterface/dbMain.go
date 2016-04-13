// Package relationDB is used to find relations between articles
package relationDB

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	// for driver
	_ "gopkg.in/cq.v1"
)

// DBKeyword comment
type DBKeyword struct {
	Text      string  `json:"text"`
	Relevance float32 `json:"relevance,omitempty"`
}

// ArticleInfo comment
type ArticleInfo struct {
	// assumes that this is universally unique
	Identifier string `json:"ArticleIdentifier"`
}

// IDError when bad id
type IDError struct {
	uuid    string
	message string
}

// Error from IDError
func (e *IDError) Error() string {
	return fmt.Sprintf("%s - %s", e.uuid, e.message)
}

// private database for all the requests
var db *sql.DB

// Open a connection to the DB if one isn't already open
// you should turn off auth by settind dbms.security.auth_enabled = false
// in neo4j/data/dbms/auth
func Open(where string) error {
	if db != nil {
		return nil
	}

	tmp, err := sql.Open("neo4j-cypher", where)
	if err != nil {
		return err
	}

	db = tmp
	return nil
}

// Close the db
func Close() error {
	return db.Close()
}

// Store an article in the DB
// articleID should be a uuid for the article
// goes and checks that this is actually a uuid to prevent double inserts
func Store(articleID string) error {
	if info, err := GetByUUID(articleID); err != nil {
		return err
	} else if info.Identifier != "" {
		return fmt.Errorf("bad uuid %s", err.Error())
	}
	stmt, err := db.Prepare(`create (:Article {Identifier:{0}})`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(articleID)
	return err
}

// GetByUUID gets an article by its uuid
func GetByUUID(articleID string) (ArticleInfo, error) {
	stmt, err := db.Prepare(`match (n {Identifier: {0}}) return n`)
	if err != nil {
		return ArticleInfo{}, err
	}

	rows, err := stmt.Query(articleID)
	if err != nil {
		return ArticleInfo{}, fmt.Errorf("problem getting by uuid: %s", err.Error())
	}
	info := ArticleInfo{}
	err = rows.Scan(&info)
	if rows.Next() {
		return ArticleInfo{}, fmt.Errorf("uuid not unique")
	}
	return info, nil
}

// clear deletes all nodes from teh db, used most for testing
func clear() error {
	stmt, err := db.Prepare("match n delete n")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

/*
// Remove something by uuid
func Remove(id uuid.UUID) error {

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
	//defer stmt.close()

	//Unwind json object and execute query
	_, err2 := stmt.Exec(id)
	//JUST DO IT^^^

	//no error = nil
	return err2

}


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
	//defer stmt.close()

	//Execute query
	_, err2 := stmt.Exec()
	//JUST DO IT^^^

	//no error = nil
	return err2

}


func Get(id uuid.UUID) (ArticleInfo, error) {

	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		var empty ArticleInfo
		return empty, err //Could not open!
	}
	defer db.Close()

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as uuid
		MATCH (article {ID: uuid})
		RETURN article
	`)
	if err != nil {
		var empty ArticleInfo
		return empty, err //Wat r u tryina git dawg?
	}
	//defer stmt.close()

	//Execute query
	ret, err := stmt.Query(id)
	//JUST DO IT^^^
	if err != nil {
		var empty ArticleInfo
		return empty, err //Something went wrong!!!
	}
	//defer ret.close()

	//pull all articles and decode. Return the first (SHOULD NOT BE MORE THAN 1, return error otherwise)
	var article ArticleInfo
	var articleJSON []byte
	err2 := ret.Scan(&(articleJSON))
	if err2 != nil {
		var empty ArticleInfo
		return empty, err2 //this one is weird
	}
	//DECODE json object
	err3 := json.Unmarshal(articleJSON, &article)
	if err3 != nil {
		var empty ArticleInfo
		return empty, err3
	}

	//Should not be any more rows, if there are return IDError on id
	for ret.Next() {
		var empty ArticleInfo
		return empty, &IDError{id.String(), "Multiple entries in database for this ID!!"}
	}

	//no error = nil
	return article, err

}


func SearchByKeyword(keyword KeyQuery) ([]uuid.UUID, error) {
	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		var empty []uuid.UUID
		return empty, err //Could not open!
	}
	defer db.Close()

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as kw
		MATCH (article)-[kw.Threshold]->(kw.Name)
		RETURN article.ID
	`)
	if err != nil {
		var empty []uuid.UUID
		return empty, err //Wat r u tryina git dawg?
	}
	//defer stmt.close()

	//Execute query
	ret, err := stmt.Query(keyword)
	//JUST DO IT^^^
	if err != nil {
		var empty []uuid.UUID
		return empty, err //Something went wrong!!!
	}
	//defer ret.close()

	//pull all articles and decode. Return the slice with all of them in it
	var aID uuid.UUID
	var idJSON []byte
	var set []uuid.UUID
	for ret.Next() {
		err := ret.Scan(&(idJSON)) //scan
		if err != nil {
			var empty []uuid.UUID
			return empty, err //this one is weird
		}
		err2 := json.Unmarshal(idJSON, &aID) //decode
		if err2 != nil {
			var empty []uuid.UUID
			return empty, err2 //more ways it can be weird
		}
		set = append(set, aID) //append
	}

	//no error = nil
	return set, err
}

func SearchByTaxonomy(taxonomy TaxQuery) ([]uuid.UUID, error) {
	//connect to database
	db, err := sql.Open("neo4j-cypher", "localhost:7474") //port is up for debate
	if err != nil {
		var empty []uuid.UUID
		return empty, err //Could not open!
	}
	defer db.Close()

	//Create query
	stmt, err := db.Prepare(`
		UNWIND {{0}} as tx
		MATCH (article)-[tx.Threshold]->(tx.Name)
		RETURN article.ID
	`)
	if err != nil {
		var empty []uuid.UUID
		return empty, err //Wat r u tryina git dawg?
	}
	//defer stmt.close()

	//Execute query
	ret, err := stmt.Query(taxonomy)
	//JUST DO IT^^^
	if err != nil {
		var empty []uuid.UUID
		return empty, err //Something went wrong!!!
	}
	//defer ret.close()

	//pull all articles and decode. Return the slice with all of them in it
	var aID uuid.UUID
	var idJSON []byte
	var set []uuid.UUID
	for ret.Next() {
		err := ret.Scan(&(idJSON)) //scan
		if err != nil {
			var empty []uuid.UUID
			return empty, err //this one is weird
		}
		err2 := json.Unmarshal(idJSON, &aID) //decode
		if err2 != nil {
			var empty []uuid.UUID
			return empty, err2 //more ways it can be weird
		}
		set = append(set, aID) //append
	}

	//no error = nil
	return set, err
}
*/
