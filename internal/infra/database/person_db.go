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

	result, err := session.Run(ctx, "MATCH (p:PERSON {uuid: $uuid}) RETURN p", map[string]interface{}{"uuid": uuid})
	if err != nil {
		fmt.Println("Run Error")
		return nil, err
	}

	person, err := helpers.GetDbResponseParsed(ctx, result, entity.Person{})
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
