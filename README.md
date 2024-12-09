# ZopDev

Zop is a comprehensive tool for managing cloud infrastructure. It consists of three main components:

1. **zop-api**: Backend API service.
2. **zop-ui**: User interface for managing and monitoring cloud resources.
3. **zop-cli**: Command-line interface for developers and admins.

---

## Installation

### Prerequisites

- Docker installed on your system.

---

### Running Locally

#### zop-api
Run the following command to pull and start the Docker image for the zop-api:

```bash
    docker run -d -p 8000:8000 --name zop-api zop.dev/zop-api:v0.0.1
```

#### zop-ui
Run the following command to pull and start the Docker image for the zop-ui:
```bash
    docker run -d -p 3000:3000 -e NEXT_PUBLIC_API_BASE_URL='http://localhost:8000' --name zop-ui zop.dev/zop-ui:v0.0.1
```

> **Note:** The environment variable `NEXT_PUBLIC_API_BASE_URL` is used by zop-ui to connect to the zop-api. Ensure that the value matches the API's running base URL.
#### zop-cli

Run the following command install zop-cli:
```bash
   go install zop.dev/clizop@latest
```

> **Note:** Set the environment variable `ZOP_API_URL`, used by zop-cli to connect to the zop-api. Ensure that the value matches the API's running base URL.

### zop-api

#### Endpoints

1. **POST /cloud-accounts**  
   Add a new cloud account.

    - **Request Body**
      ```json
      {
          "name": "zop cloud Account",
          "provider": "gcp",
          "credentials": {}
      }
      ```

    - **Description**  
      Use this endpoint to add a cloud account by specifying its name, provider, and credentials.

---

2. **GET /cloud-accounts**  
   List all cloud accounts.

    - **Sample Response**
      ```json
      {
          "data": [
              {
                  "name": "zop cloud Account",
                  "id": 1,
                  "provider": "gcp",
                  "providerId": "caramel-park-443607-n22wqe5331",
                  "providerDetails": null,
                  "createdAt": "2024-12-09T09:48:27Z",
                  "updatedAt": "2024-12-09T09:48:27Z"
              }
          ]
      }
      ```

    - **Description**  
      This endpoint returns a list of all registered cloud accounts with their details.
