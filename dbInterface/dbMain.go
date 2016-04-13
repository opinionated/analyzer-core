// Package relationDB is used to find relations between articles
package relationDB

import (
	"fmt"
	// use neoism
	"gopkg.in/jmcvetta/neoism.v1"
)

// DBKeyword comment
type DBKeyword struct {
	Text      string  `json:"text"`
	Relevance float32 `json:"relevance,omitempty"`
}

// ArticleInfo comment
type ArticleInfo struct {
	// assumes that this is universally unique
	Identifier string `json:"n.Identifier"`
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
//var db *sql.DB
var db *neoism.Database

// Open a connection to the DB if one isn't already open
// you should turn off auth by settind dbms.security.auth_enabled = false
// in neo4j/data/dbms/auth
func Open(where string) error {
	if db != nil {
		return nil
	}

	tmp, err := neoism.Connect(where)
	if err != nil {
		return err
	}

	db = tmp
	return nil
}

// Close the db
func Close() error {
	return nil
}

// Store an article in the DB
// articleID should be a uuid for the article
// goes and checks that this is actually a uuid to prevent double inserts
func Store(articleID string) error {
	if info, err := GetByUUID(articleID); err != nil {
		return fmt.Errorf("bad uuid: %s", err.Error())
	} else if info.Identifier != "" {
		return fmt.Errorf("uuid not unique")
	}

	cq := neoism.CypherQuery{
		Statement:  `create (:Article {Identifier:{Identifier}})`,
		Parameters: neoism.Props{"Identifier": articleID},
		Result:     nil,
	}
	return db.Cypher(&cq)
}

// GetByUUID gets an article by its uuid
func GetByUUID(articleID string) (ArticleInfo, error) {
	result := []ArticleInfo{}

	cq := neoism.CypherQuery{
		Statement:  `match (n {Identifier: {Identifier} }) return n, n.Identifier`,
		Parameters: neoism.Props{"Identifier": articleID},
		Result:     &result,
	}

	err := db.Cypher(&cq)
	if err != nil {
		return ArticleInfo{}, err
	}

	fmt.Println(result)

	if len(result) > 1 {
		return result[0], fmt.Errorf("too many articles returned!\n")
	}
	if len(result) > 0 {
		return result[0], nil
	}

	// nothing
	return ArticleInfo{}, nil

}

// StrengthBetween find how closely two nodes are related by some label
// finds all paths between and sums up the output
func StrengthBetween(startID string, endID string, label string) (float32, int, error) {
	result := []struct {
		Score float32 `json:"total"`
		Count int     `json:"count"`
	}{}

	cq := neoism.CypherQuery{
		Statement: `
	match (start:Article {Identifier: {startID}}),(end:Article {Identifier: {endID}}) 
	match p = (start)-[rel_s]-(mid:Keyword)-[rel_e]-(end) with collect(p) as paths

	return reduce(o_s = 0, path in paths 
	| o_s + reduce(s = 0, rel in relationships(path) | s + rel.Relevance)) as total, length(paths) as count
						`,
		Parameters: neoism.Props{"startID": startID, "endID": endID, "label": label},
		Result:     &result,
	}
	/*match (start:Article {Identifier: {startID}}),(end:Article {Identifier: {endID}})
	match p = (start)-[:Keyword]-(mid:Keyword)-[:Keyword]-(end) with collect(p) as paths
	return reduce(o_s = 0, path in paths | o_s +
	  reduce(s = 0, rel in relationships(path) | s + rel.Relevance)) as total,
		count(paths) as count
	*/

	//match (start:Article {Identifier: {startID}})-[rel_s]-(mid:Keyword)-[rel_e]-(end:Article {Identifier: {endID}})
	//return mid, mid.Text, rel_s.Relation, start.Identifier

	//where rel_s.Relevance > 0.5 and rel_e.Relevance > 0.5
	err := db.Cypher(&cq)
	if err != nil {
		return 0, 0, err
	}
	if len(result) != 1 {
		return 0, 0, fmt.Errorf("result is too long")
	}
	return result[0].Score, result[0].Count, err
}

// InsertRelations inserts an array of relations named by keyword
// assumes that values has Text, Relevance
func InsertRelations(articleID string, keyword string, values interface{}) error {

	cq := neoism.CypherQuery{
		Statement: `
			match (start:Article {Identifier: {articleID}})
			unwind {relations} as relations
			foreach (relation in relations | 
			merge (end:Keyword {Text: relation.Text})
			create unique (start)-[:Relation {Relevance: relation.Relevance}]->(end)
			)
	`,
		Parameters: neoism.Props{"articleID": articleID, "keyword": keyword, "relations": values},
	}

	//create unique (start)-[:Relation {Relevance: relation.Relevance}]-(:Keyword {Text: relation.Text})
	err := db.Cypher(&cq)
	return err
}

// clear deletes all nodes from teh db, used most for testing
func clear() error {
	cq := neoism.CypherQuery{
		Statement: `
		match (node) optional match (node)-[edge]-() 
		delete node, edge`,
	}

	return db.Cypher(&cq)
}

// TODO(@max): get by type ie (n)-[relevance > thresh]-(:{Keyword} {Text = {text} })
