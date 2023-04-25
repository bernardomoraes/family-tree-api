package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/bernardomoraes/family-tree/internal/entity"
	"github.com/bernardomoraes/family-tree/pkg/helpers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Relationship struct {
	DBDriver neo4j.DriverWithContext
}

func NewRelationship(driver neo4j.DriverWithContext) *Relationship {
	return &Relationship{DBDriver: driver}
}

func (r *Relationship) CreateIsParent(ctx context.Context, parent entity.Person, child entity.Person) error {
	session := helpers.NewSession(ctx, r.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"parentUUID": parent.UUID,
		"childUUID":  child.UUID,
	}

	dbResult, err := session.Run(ctx, `MATCH (parent:PERSON {uuid: $parentUUID}), (child:PERSON {uuid: $childUUID}) CREATE (parent)-[r:IS_PARENT]->(child) RETURN parent, child`, parameters)

	if err != nil {
		return err
	}

	resParsed, err := helpers.GetDbResponseParsed(ctx, dbResult, entity.Person{})
	if err != nil {
		return err
	}

	fmt.Println("resParsed:", resParsed)

	return nil
}

func (r *Relationship) CreateIsSpouse(ctx context.Context, person1 entity.Person, person2 entity.Person) error {
	session := helpers.NewSession(ctx, r.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"person1UUID": person1.UUID,
		"person2UUID": person2.UUID,
	}

	dbResult, err := session.Run(ctx, `MATCH (person1:PERSON {uuid: $person1UUID}), (person2:PERSON {uuid: $person2UUID}) CREATE (person1)-[r:IS_SPOUSE]->(person2) RETURN person1, person2`, parameters)

	if err != nil {
		return err
	}

	resParsed, err := helpers.GetDbResponseParsed(ctx, dbResult, entity.Person{})
	if err != nil {
		return err
	}

	fmt.Println("resParsed:", resParsed)

	return nil
}

func (r *Relationship) FindRelationship(ctx context.Context, relationship *entity.Relationship) (*entity.Relationship, error) {
	session := helpers.NewSession(ctx, r.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"parentUUID": relationship.Start,
		"childUUID":  relationship.End,
	}
	query := `MATCH (parent:PERSON {uuid: $parentUUID})-[r:IS_PARENT]->(child:PERSON {uuid: $childUUID}) 
	RETURN {
		start: parent.uuid, 
		end: child.uuid, 
		relation: type(r) 
	} as result`

	dbResult, err := session.Run(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	resParsed, err := helpers.GetDbResponse(ctx, dbResult, entity.Relationship{})
	if err != nil {
		return nil, err
	}

	if len(resParsed) == 0 {
		return nil, nil
	}

	fmt.Println("resParsed:", resParsed)

	return &resParsed[0], nil
}

func (r *Relationship) FindRelationshipsFromPerson(ctx context.Context, person entity.Person) ([]entity.Relationship, error) {
	session := helpers.NewSession(ctx, r.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"uuid": person.UUID,
	}
	query := `
	match (n:PERSON {uuid: $uuid}) 
	optional match (n)<-[:IS_PARENT]-(rp)
	optional match (n)-[:IS_PARENT]->(rc)
	with n, collect(distinct rp{.name, .uuid}) as parent, collect(distinct rc{.name, .uuid}) as child
	with n, n{parent, child} as relationships
	return n{.*, relationships}
	`

	dbResult, err := session.Run(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	resParsed, err := helpers.GetDbResponse(ctx, dbResult, entity.Relationship{})
	if err != nil {
		return nil, err
	}

	fmt.Println("resParsed:", resParsed)

	return resParsed, nil
}

func (r *Relationship) GetDegreeSeparation(ctx context.Context, person1 *entity.Person, person2 *entity.Person) (int, error) {
	session := helpers.NewSession(ctx, r.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"start": person1.UUID,
		"end":   person2.UUID,
	}
	query := `
	MATCH (p1:PERSON {uuid:$start})
	MATCH (p2:PERSON {uuid:$end})
	MATCH p=shortestPath((p1)-[*]-(p2))
	RETURN length(p) as bc;
	`

	dbResult, err := session.Run(ctx, query, parameters)
	if err != nil {
		return 0, err
	}
	record := dbResult.Record()
	if !dbResult.NextRecord(ctx, &record) {
		return 0, errors.New("no path")
	}
	fmt.Println(record)

	recordItem, found := dbResult.Record().Get(record.Keys[0])
	if !found {
		return 0, errors.New("no path")
	}

	fmt.Println("recordItem:", recordItem)

	return int(recordItem.(int64)), nil
}
