package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/bernardomoraes/family-tree/internal/entity"
	"github.com/bernardomoraes/family-tree/pkg/helpers"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type User struct {
	DBDriver neo4j.DriverWithContext
}

func NewUser(driver neo4j.DriverWithContext) *User {
	return &User{DBDriver: driver}
}

func (p *User) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	session := helpers.NewSession(ctx, p.DBDriver, neo4j.AccessModeWrite)
	defer session.Close(ctx)

	parameters := map[string]interface{}{
		"uuid":     user.UUID,
		"email":    user.Email,
		"name":     user.Name,
		"password": user.Password,
	}
	fmt.Println("parameters:", parameters)
	dbResult, err := session.Run(ctx, "CREATE (user:USER {uuid: $uuid, name: $name, email: $email, password: $password}) RETURN user",
		parameters)

	if err != nil {
		return nil, errors.New("Error when creating user: " + err.Error())
	}

	resParsed, err := helpers.GetDbResponseParsed(ctx, dbResult, entity.User{})

	if err != nil {
		return nil, err
	}

	fmt.Println("resParsed:", resParsed)

	return &resParsed[0], nil
}

func (u *User) Update(ctx context.Context, user *entity.User) error {
	session := (u.DBDriver).NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)
	_, err := session.Run(ctx, "MATCH (u:USER {uuid: $uuid}) SET u.name = $name, u.email = $email, u.password = $password RETURN u", map[string]interface{}{"uuid": user.UUID, "name": user.Name, "email": user.Email, "password": user.Password})
	return err
}

func (p *User) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	session := p.DBDriver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})

	res, err := session.Run(ctx, "MATCH (u:USER {email:$email}) RETURN u LIMIT 1 ", map[string]interface{}{
		"email": email,
	})

	if err != nil {
		return nil, err
	}

	if res.Next(ctx) {
		result, ok := res.Record().Values[0].(neo4j.Node)
		if !ok {
			return nil, errors.New("unexpected type for user")
		}
		fmt.Println("retorno", result.GetProperties()["email"].(string))
		// return &entity.User{Email: user.GetProperties()["email"].(string)}, nil
		var user entity.User
		mapstructure.Decode(result.Props, &user)
		return &user, nil
	}

	return nil, nil
}

func (u *User) FindByUUID(ctx context.Context, uuid string) (*entity.User, error) {
	session := (u.DBDriver).NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close(ctx)
	fmt.Println(uuid)
	result, err := session.Run(ctx, "MATCH (u:USER {uuid: $uuid}) RETURN u", map[string]interface{}{"uuid": uuid})
	if err != nil {
		fmt.Println("Run Error")
		return nil, err
	}

	user, err := helpers.GetDbResponseParsed(ctx, result, entity.User{})
	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		return nil, errors.New("User not found")
	}
	return &user[0], nil
}

func (p *User) FindAll(ctx context.Context) ([]entity.User, error) {
	session := (p.DBDriver).NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})

	defer session.Close(ctx)
	result, err := session.Run(ctx, "MATCH (u:USER) RETURN u",
		map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	users, err := helpers.GetDbResponseParsed(ctx, result, entity.User{})

	if err != nil {
		fmt.Println("Error Parsing:", err)
		return nil, err
	}

	return users, nil
}

func (u *User) DeleteByUUID(ctx context.Context, uuid string) error {
	session := (u.DBDriver).NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)
	_, err := session.Run(ctx, "MATCH (u:USER {uuid: $uuid}) DETACH DELETE u", map[string]interface{}{"uuid": uuid})
	return err
}
