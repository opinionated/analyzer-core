package dbInterface

import{
	"analyzer/structures"
	"net/http"
	"strconv"
	"gopkg.in/cq.v1"
}



//Add relationship between two articles given score
func AddRelationship(id1 int, id2 int, score float64){
	query := "
		MATCH (a:artilcleInfo), (b:articleInfo)
		WHERE a.id = '"+id1+"' AND b.id = '"+id2+"'
		CREATE a-[r:SCORE{'score':'"score"'}]->b
	"
}
//Add new node
func AddArticle(article articleInfo){
	query := "
		CREATE article:articleInfo
	"
}