# Courses Rest Api
REST API that handles CRUD about courses, made using Go, Echo, PostgreSQL, and JWT.
This project has clean architecture folder structure, which is based on [this github repository](https://github.com/bxcodec/go-clean-arch). The database diagram for this API can be seen at [this link](https://drive.google.com/file/d/1ERygjzOBWGYA6V3QEc9002SgqKUH65co/view?usp=share_link). The API documentation can be accessed at [this link](https://documenter.getpostman.com/view/457088/2s935vnLhb).

## How to run the code
1. Create a new database for this API. No need to manually create the tables in the new database because the tables will be created automatically after executing `go run .`(Step 5).
2. Copy and rename file `example.env` into `.env`.
3. Open file `.env` and change `PORT`, `SECRET`, and `DATABASE_URL` into the appropriate port, database url, and secret for generating JWT token.
4. Open terminal, go into root directory of this code, and run `go mod tidy`.
5. Then run `go run .`.
6. Press `Ctrl + C` to terminate the API.