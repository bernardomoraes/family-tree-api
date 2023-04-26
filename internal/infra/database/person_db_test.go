package database

import (
	"context"
	"fmt"

	"github.com/bernardomoraes/family-tree/internal/entity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var _ = Describe("Person repository", func() {
	const username = "neo4j"
	const password = "s3cr3t"

	var ctx context.Context
	var neo4jContainer testcontainers.Container
	var driver neo4j.DriverWithContext
	var repository *Person

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		neo4jContainer, err = startContainer(ctx, username, password)
		Expect(err).To(BeNil(), "Container should start")
		port, err := neo4jContainer.MappedPort(ctx, "7687")
		Expect(err).To(BeNil(), "Port should be resolved")
		address := fmt.Sprintf("bolt://localhost:%d", port.Int())
		driver, err = neo4j.NewDriverWithContext(address, neo4j.BasicAuth(username, password, ""))
		Expect(err).To(BeNil(), "Driver should be created")
		repository = &Person{
			DBDriver: driver,
		}
	})

	AfterEach(func() {
		driver.Close(ctx)
		Expect(neo4jContainer.Terminate(ctx)).To(BeNil(), "Container should stop")
	})

	It("Create person", func() {
		personName := "some-name"
		person, _ := entity.NewPerson(personName)

		createdPerson, err := repository.Create(ctx, person)

		Expect(err).To(BeNil(), "Person should be registered")
		Expect(createdPerson.UUID).To(Equal(person.UUID))

		session := driver.NewSession(ctx, neo4j.SessionConfig{})
		defer session.Close(ctx)

		result, err := session.Run(ctx, `MATCH (p:Person {uuid: $uuid}) RETURN p.name`, map[string]interface{}{"uuid": person.UUID})
		Expect(err).To(BeNil(), "Query should successfully run")
		if !result.Next(ctx) {
			Fail("Record should be retrieved")
		}
		persistedPerson, ok := result.Record().Get(result.Record().Keys[0])
		if !ok {
			Fail("Record should be retrieved")
		}
		Expect(persistedPerson.(map[string]interface{})["name"]).To(Equal(personName))
		Expect(persistedPerson.(map[string]interface{})["uuid"]).To(Equal(person.UUID))
	})
})

func startContainer(ctx context.Context, username, password string) (testcontainers.Container, error) {
	request := testcontainers.ContainerRequest{
		Image:        "neo4j",
		ExposedPorts: []string{"7687/tcp"},
		Env:          map[string]string{"NEO4J_AUTH": fmt.Sprintf("%s/%s", username, password)},
		WaitingFor:   wait.ForLog("Bolt enabled"),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}
