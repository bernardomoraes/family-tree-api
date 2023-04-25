package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/bernardomoraes/family-tree/internal/entity"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type AvailableDatabaseDrivers interface {
	neo4j.DriverWithContext
}

func GetDbResponseParsed[T any](ctx context.Context, result neo4j.ResultWithContext, resultType T) ([]T, error) {
	record := result.Record()
	response := []T{}

	for result.NextRecord(ctx, &record) {
		recordItem, found := record.Get(record.Keys[0])
		if !found {
			fmt.Println("Not found record item")
			continue
		}
		recordItemProps := recordItem.(neo4j.Node).Props
		recordItemProps["id"] = recordItem.(neo4j.Node).GetId()

		if recordItemProps["createdAt"] != nil {
			recordItemProps["createdAt"] = recordItemProps["createdAt"].(time.Time).Format(time.RFC3339)
		}

		if recordItemProps["updatedAt"] != nil {
			recordItemProps["updatedAt"] = recordItemProps["updatedAt"].(time.Time).Format(time.RFC3339)
		}

		err := mapstructure.Decode(recordItemProps, &resultType)
		if err != nil {
			fmt.Println("Error decoding record item")
			return nil, err
		}

		response = append(response, resultType)
	}
	return response, nil
}

func GetDbResponse[T any](ctx context.Context, result neo4j.ResultWithContext, resultType T) ([]T, error) {
	record := result.Record()
	response := []T{}

	for result.NextRecord(ctx, &record) {
		fmt.Println("record:", record)
		recordItem, found := record.Get(record.Keys[0])
		if !found {
			fmt.Println("Not found record item")
			continue
		}

		err := mapstructure.Decode(recordItem, &resultType)
		if err != nil {
			fmt.Println("Error decoding record item")
			return nil, err
		}
		response = append(response, resultType)
	}
	return response, nil
}
func GetDbResponseMap[T *entity.Person](ctx context.Context, result neo4j.ResultWithContext, resultType *entity.Person) ([]T, error) {
	record := result.Record()
	response := []T{}

	for result.NextRecord(ctx, &record) {
		fmt.Println("record:", record)
		recordItem, found := record.Get(record.Keys[0])

		if !found {
			fmt.Println("Not found record item")
			continue
		}
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			WeaklyTypedInput: false,
			Result:           &resultType,
			ZeroFields:       true,
		})
		decoder.Decode(recordItem)

		if err != nil {
			fmt.Println("Error decoding record item")
			return nil, err
		}
		response = append(response, resultType)
	}
	return response, nil
}

func Driver(target string, user string, password string) neo4j.DriverWithContext {

	token := neo4j.BasicAuth(user, password, "")

	driver, err := neo4j.NewDriverWithContext(target, token)
	if err != nil {
		panic(err)
	}
	return driver
}

func NewSession(ctx context.Context, driver neo4j.DriverWithContext, AccessMode neo4j.AccessMode) neo4j.SessionWithContext {
	session := (driver).NewSession(ctx, neo4j.SessionConfig{
		AccessMode: AccessMode,
	})
	return session
}
