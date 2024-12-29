# Click the planet - version 3
## New version of the clicker game made using different technologies

### How to launch?
  - Frontend:
    ```console
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
      go install github.com/air-verse/air@latest
      air init
      air
    ```
  - Database
    ```console
      docker run --name clicktheplanet-v3 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -p 5432:5432 -d postgres
    ```
