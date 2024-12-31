# Click the planet - version 3

## New version of the clicker game made using different technologies

### How to launch?

- Frontend:
  ```console
    cd client
    npm i
  ```
  After that either:
  ```console
    npm run dev
  ```
  or
  ```console
    npm run build
    npm run start
  ```
- Backend
  ```console
    cd server
    go install github.com/air-verse/air@latest
    air init
    air
  ```
- Database
  ```console
    docker run --name clicktheplanet-v3 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=clicktheplanet-v3 -p 5432:5432 -d postgres
  ```
