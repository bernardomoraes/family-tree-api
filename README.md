# **Family Tree API**

This project aims to create an API to manage a family tree with the ability to create and manipulate Person entities and their parent-child relationships. Additionally, it provides an endpoint to retrieve a person's genealogical tree.

## **Core Requirements**
- [x]  Post endpoint that receives as a body param the person's name and creates a new Person entity.
- [x]  Post endpoint that receives as a body param the parent's and child's identifiers and creates a parent-child relationship between them.
- [x]  Get endpoint that returns a person's genealogical tree, including all possible ancestors up to their level.

## **Extra Functionalities**
> **Category:** Business Logic
- [x] Addition of nodes for children, siblings, and nieces/nephews.
- [x] Add relationships of kinship endpoint, such as cousins, spouse.
- [x] Implement Bacon's Number for a given pair of persons.
- [ ] Allow clients to choose the format of the response, based on the accept header.
- [ ] Add the spouse relationship between two members that share children (and modify the Bacon's Number between them).


> **Category:** Technical
- [ ] Create a CI pipeline with GitHub Actions.
- [ ] Prevent incestuous offspring, based on the restriction described in the challenge.
## **Deliverables**

- [x]  GitHub public repository with the core requirements implemented.
- [x]  Dockerfile for building a Docker image and a Docker Compose file to run the application, including the necessary services (database, environment variables, auxiliary services).
- [x]  Documentation file that explains how to use the API.

## **How to Run and Use**
### **Pre Requirements:**
- Make sure you have **[Docker](https://docs.docker.com/get-docker/)** and **[Docker Compose](https://docs.docker.com/compose/install/)** installed on your machine.
- Clone the repository: `git clone https://github.com/bernardomoraes/family-tree-api.git`

### **First Time Setup:**
1. In the project's root directory, create a file named `.env` with the following content, there is an example file in the repository:
  ```bash
  cp .env.example .env
  ```
2. Build the images:
  ```bash
  docker compose build --parallel
  ```

### **Running the API:**
**Note:** The API will be available by default at `http://localhost:8080`. (You can change the port in the `.env` file)
  ```bash
  # Using the default configuration
  docker compose up

  # Using a custom .env configuration file
  # Should be used if you want to change the default port dynamically 
  # by passing the PORT environment variable in the .env file
  docker compose --env-file ./.env up
  ```

### Extra:
> **[Air](https://github.com/cosmtrek/air)** to live-reloading. All the changes will be reflected in the API automatically, so you don't need to restart the container.

> **Neo4j browser** available at `http://localhost:7474/browser/`. You can use it to inspect the database and run Cypher queries.

> ~~**Swagger UI** available at `http://localhost:8080/docs/index.html`. You can use it to test the API endpoints.~~ (For some reason, the Swagger UI is not working properly with the API, I tried to fix it but I couldn't find the solution.)

## **Technical Decisions**
- **[Chi](https://go-chi.io/)** as the router.
- **[Neo4j](https://neo4j.com/)** as the database and interact with it using the **[Neo4j Go Driver](https://github.com/neo4j/neo4j-go-driver)**.
- **[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)** as the architectural pattern.
- **[Repository pattern](https://martinfowler.com/eaaCatalog/repository.html)** to abstract the database access.
- Implement Unit Tests and Integration Tests using **[Testify](https://github.com/stretchr/testify)**.
