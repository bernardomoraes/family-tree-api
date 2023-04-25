package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/bernardomoraes/family-tree/internal/entity"
	"github.com/bernardomoraes/family-tree/pkg/helpers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Person struct {
	DBDriver neo4j.DriverWithContext
}

func NewPerson(driver neo4j.DriverWithContext) *Person {
	return &Person{DBDriver: driver}
}

func (p *Person) Create(ctx context.Context, person *entity.Person) (*entity.Person, error) {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"uuid": person.UUID,
		"name": person.Name,
	}

	dbResult, err := session.Run(ctx, "CREATE (person:PERSON {uuid: $uuid, name: $name, createdAt: datetime.statement()}) RETURN person",
		parameters)

	if err != nil {
		return nil, errors.New("Error when creating person: " + err.Error())
	}

	resParsed, err := helpers.GetDbResponseParsed(ctx, dbResult, entity.Person{})

	if err != nil {
		return nil, err
	}

	fmt.Println("resParsed:", resParsed)

	return &resParsed[0], nil
}

func (p *Person) FindByUUID(ctx context.Context, uuid string) (*entity.Person, error) {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	if uuid == "" {
		return nil, errors.New("uuid is empty")
	}
	findByUUIDQuery := `
		match (n:PERSON {uuid: $uuid}) 
		optional match (n)<-[:IS_PARENT]-(rp)
		optional match (n)-[:IS_PARENT]->(rc)
		with n, collect(distinct rp{.name, .uuid}) as parents, collect(distinct rc{.name, .uuid}) as childs
		with n, n{parents, childs} as relationships
		return n{.*, relationships}
	`

	result, err := session.Run(ctx, findByUUIDQuery, map[string]interface{}{"uuid": uuid})
	if err != nil {
		fmt.Println("Run Error")
		return nil, err
	}

	person, err := helpers.GetDbResponseMap(ctx, result, &entity.Person{})

	fmt.Println("person:", person[0])
	if err != nil {
		return nil, err
	}

	if len(person) == 0 {
		return nil, nil
	}
	return person[0], nil
}
func (p *Person) FindByName(ctx context.Context, name string) (*entity.Person, error) {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	if name == "" {
		return nil, errors.New("name is empty")
	}
	findByNameQuery := `
		match (n:PERSON {name: $name}) 
		optional match (n)<-[:IS_PARENT]-(rp)
		optional match (n)-[:IS_PARENT]->(rc)
		with n, collect(distinct rp{.name, .uuid}) as parents, collect(distinct rc{.name, .uuid}) as childs
		with n, n{parents, childs} as relationships
		return n{.*, relationships}
	`

	result, err := session.Run(ctx, findByNameQuery, map[string]interface{}{"name": name})
	if err != nil {
		fmt.Println("Run Error")
		return nil, err
	}

	person, err := helpers.GetDbResponse(ctx, result, entity.Person{})
	if err != nil {
		return nil, err
	}

	if len(person) == 0 {
		return nil, nil
	}
	return &person[0], nil
}

func (p *Person) Update(ctx context.Context, person *entity.Person) (*entity.Person, error) {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	queryResult, err := session.Run(ctx, "MATCH (u:PERSON {uuid: $uuid}) SET u.name = $name, u.updatedAt=datetime() RETURN u",
		map[string]interface{}{
			"uuid": person.UUID,
			"name": person.Name,
		})
	if err != nil {
		return nil, err
	}

	personParsed, err := helpers.GetDbResponseParsed(ctx, queryResult, entity.Person{})
	if err != nil {
		return nil, err
	}

	if len(personParsed) == 0 {
		return nil, nil
	}

	return &personParsed[0], err
}

func (p *Person) Delete(ctx context.Context, uuid string) error {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	_, err := session.Run(ctx, "MATCH (p:PERSON {uuid: $uuid}) DETACH DELETE p", map[string]interface{}{"uuid": uuid})
	if err != nil {
		return err
	}

	return nil
}

func (p *Person) FindAncestors(ctx context.Context, person *entity.Person) ([]*entity.Person, error) {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	// result, err := session.Run(ctx, "MATCH (p:PERSON {uuid: $uuid})-[:IS_PARENT_OF*1..]->(a:PERSON) RETURN a", map[string]interface{}{"uuid": uuid})
	query := `
		match (main:PERSON{uuid: $uuid})
		call {
				with main
				match (main)<-[r:IS_PARENT]-(first_ancestor:PERSON)
				match (main)-[ch:IS_PARENT]->(child:PERSON)
				with main, collect(distinct first_ancestor{.name, .uuid}) as parents, collect(distinct child{.name, .uuid}) as childs
				with main, main{parents, childs} as relationships
				return main{.name, .uuid, relationships} as res
			union
				with main
				match (ancestor:PERSON)-[:IS_PARENT*]->(main)
				with ancestor
				match (ancestor)<-[:IS_PARENT]-(pa:PERSON)
				with ancestor, pa
				match (ancestor)-[:IS_PARENT]->(ch:PERSON)
				with  ancestor, collect(distinct pa{.name, .uuid}) as parents, collect(distinct ch{.name, .uuid}) as childs
				with ancestor, ancestor{parents, childs} as relationships
				return ancestor{.name, .uuid, relationships} as res
		}
		return res
	`
	result, err := session.Run(ctx, query, map[string]interface{}{"uuid": person.UUID})
	if err != nil {
		fmt.Println("Run Error")
		return nil, err
	}

	personParsed, err := helpers.GetDbResponseMap(ctx, result, &entity.Person{})

	if err != nil {
		return nil, err
	}

	if len(personParsed) == 0 {
		return nil, nil
	}
	return personParsed, nil
}
