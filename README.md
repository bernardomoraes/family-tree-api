# **Family Tree API**

This project aims to create an API to manage a family tree with the ability to create and manipulate Person entities and their parent-child relationships. Additionally, it provides an endpoint to retrieve a person's genealogical tree.

## **Core Requirements**

- [x]  Post endpoint that receives as a body param the person's name and creates a new Person entity.
- [ ]  Post endpoint that receives as a body param the parent's and child's identifiers and creates a parent-child relationship between them.
- [ ]  Get endpoint that returns a person's genealogical tree, including all possible ancestors up to their level.

## **Extra Functionalities**
> **Category:** Business Logic
- [ ] Addition of nodes for children, siblings, and nieces/nephews.
- [ ] Allow clients to choose the format of the response, based on the accept header.
- [ ] Add relationships of kinship, such as cousins.
- [ ] Implement Bacon's Number for a given pair of persons.
- [ ] Add the spouse relationship between two members that share children (and modify the Bacon's Number between them).


> **Category:** Technical
- [ ] Create a CI pipeline with GitHub Actions. (WIP)
- [ ] Prevent incestuous offspring, based on the restriction described in the challenge.
- [ ] Improve performance by using [Profiling](https://go.dev/blog/pprof) to the API.

## **Deliverables**

- [x]  GitHub public repository with the core requirements implemented.
- [x]  Dockerfile for building a Docker image and a Docker Compose file to run the application, including the necessary services (database, environment variables, auxiliary services).
- [ ]  Documentation file that explains how to use the API.

## **How to Run and Use**

- WIP

## **Final Considerations**

- I decided to use the **[Go](https://golang.org/)** programming language to implement the API.
- Use **[Chi](https://go-chi.io/)** as the router.
- Use **[Neo4j](https://neo4j.com/)** as the database and interact with it using the **[Neo4j Go Driver](https://github.com/neo4j/neo4j-go-driver)**.
- Use the Repository pattern to interact with the database.
- Implement Unit Tests and Integration Tests using **[Testify](https://github.com/stretchr/testify)**.
    

## **Future Improvements**
- WIP