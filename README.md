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
- [ ] Prevent incestuous offspring.

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

### **Using the API:**

#### **Endpoints:**
- **POST** `/person`
  - **Description:** Creates a new Person entity.
  - **Body:**
    ```json
    {
      "name": "string"
    }
    ```
  - **Response:**
    ```json
    {
      "uuid": "string",
      "name": "string"
    }
    ```

- **POST** `/relationship`
  - **Description:** Creates a new parent-child relationship between two Person entities.
  - **Body:**
    ```json
    {
      "parent": "parent_uuid",
      "child": "child_uuid"
    }
    ```
  - **Response:**
    ```json
    {
      "parent_uuid": "string",
      "child_uuid": "string"
    }
    ```
- **GET** `/person/{person_uuid}`
  - **Description:** Returns a person and their parents and children.
  - **Response:**
    ```json
    {
      "uuid": "string",
      "name": "string",
      "parents": [
        {
          "uuid": "string",
          "name": "string"
        }
      ],
      "children": [
        {
          "uuid": "string",
          "name": "string"
        }
      ]
    }
    ```
- **DELETE** `/person/{person_uuid}`
  - **Description:** Deletes a person and their relationships.
  - **Response:**
    ```json
    {
      "message": "string"
    }
    ```
- **GET** `/person/{person_uuid}/ancestors`
  - **Description:** Returns a person and their ancestors.
  - **Response:**
    ```json
    {
      "uuid": "string",
      "name": "string",
      "ancestors": [
        {
          "uuid": "string",
          "name": "string",
          "relationships": {
            "parents": [
              {
                "uuid": "string",
                "name": "string"
              }
            ],
            "childs": {
              "uuid": "string",
              "name": "string"
            }
          }
        }
      ]
    }
    ```
- **GET** `/person/{person_uuid}/family`
- **Description:** Returns a person and their family (Ancestors with their childs).
- **Response:**
  ```json
  {
    "uuid": "string",
    "name": "string",
    "family": [
      {
        "uuid": "string",
        "name": "string",
        "relationships": {
          "parents": [
            {
              "uuid": "string",
              "name": "string"
            }
          ],
          "childs": {
            "uuid": "string",
            "name": "string"
          }
        }
      }
    ]
  }
  ```

- **GET** `/relationship/{person1_uuid}/bacon_number/{person2_uuid}`
  - **Description:** Returns the Bacon's Number between two persons.
  - **Response:**
    ```json
    {
      "bacon_number": "int"
    }
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


## **Final Considerations**

### **What went well**

- 100% of the entities are covered by tests.

- I was able to implement all the core requirements and some of the extra functionalities.

- I was able to implement the API using the **Clean Architecture** and the Repository Pattern, I think it's a good way to organize the code and make it more testable.

- The containers are working properly and it's easy to use.

- I was able to implement the live-reloading using **Air**.

- I was able to implement the **Neo4j** database and interact with it using the **Neo4j Go Driver**.

- To implement the Accept Header should be easy, I just need to create a middleware to check the Accept Header and return the response in the correct format, but I decided to focus on the core requirements and the extra functionalities.

### **Future Improvements**
- I had some problems with the Neo4j Go Driver, I couldn't find a way to mock the database connection, so I had to create a Neo4j instance in the integration tests, and this will take some time so I decided to not implement the tests for the repository layer. But there is a example of how it would be in the person_db_test.go file.

- Implement more tests and raise the code coverage. As I decided to focus on the core requirements and the extra functionalities with the clean architecture and the repository pattern there was no time to implement tests to use cases and repository.

- For some reason, the Swaggo (Swagger for Golang) was not working properly with the API, I tried to fix it but I couldn't find the solution. I noticed that other people had the same problem, so I think it's a problem with the Swaggo itself. I will try to take a look again in the future.

- Implement API Authentication and Authorization, all the basic configuration is already done, I just need to implement the authentication and authorization middleware.
