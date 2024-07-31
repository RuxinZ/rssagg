# GO RSS AGGREGATOR

## Progress

### 1. Set up HTTP server

1. Create env file
   - Use `godotenv.Load()` to load the variables from `.env` file
   - Use `os.Getenv()` to get the value of PORT
2. Create route and server
   - Create a `chi.NewRouter`
   - Use `router.Use` to add the built-in cors.Handler middleware
   - Create a new http.Server and add the port and the main router to it
   - Test the server
3. Create two JSON helper functions
   - `respondWithJSON(w http.ResponseWriter, status int, payload interface{})` and `respondWithError(w http.ResponseWriter, status int, msg string)`
   - They write an HTTP response with a status code, an `application/json` content type, and a JSON body
4. Add a health check handler and an error handler
   - Add a handler for `GET /v1/healthz` requests. It should return a 200 status code.
   - Add a handler for `GET /v1/err` requests.It should return a 500 status code

### 2. Set up PostgreSQL database & Create Users

Tools added:

1. sqlc - generates Go code from SQL queries
2. Goose - for database migration
3. database/sql - part of Go's standard library. It provides a way to connect to a SQL database, execute queries, and scan the results into Go types.

Steps:

1. Create a migration file in `sql/schema` called `001_users.sql`
2. `cd` into the `sql/schema` directory and run: `goose postgres [DB_URL here] up`
3. Write a query to create a user: Create a file called `sqlc.yaml` in the root directory.
4. Generate the Go code: Run `sqlc generate` from the root of the project.
5. Open a connection to the database, and store it in a config struct
6. Create an HTTP handler to create a user
7. Test the handler with an HTTP client

### 3. API key

1. Add an "api key" column to the users table
   - Use a new migration file in the sql/schema directory to add a new column to the users table.
   - Generate a valid default API keys (256-bit hex values) because we already have some users in the database
2. Create an API key for new users
   - Update the "create user" SQL query and run `sqlc generate`
3. Add a new SQL query to get a user by their API key
4. Run sqlc generate to generate new Go code
   - Remember: each time you update your queries or schema you'll need to regenerate your Go code with sqlc generate. If you update the schema you'll also need to migrate your database up (and maybe down).
5. Add an endpoint to return the current user

### 4. Create feed

1. Create a feeds table
2. Add a query to create feed
3. Create some authentication middleware

   - Move the logic for authenticating a user from `handlerGetUser` to a newly created `middlewareAuth`
   - This middleware returns a callback function that accepts standard HTTP request params `w http.ResponseWriter, r *http.Request` and returns a function that we define as type `authedHandler`, `authedHandler` also accepts a `database.User` param
   - Use the middleware<br>
     `v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))  `

4. Add a handler for creating feed & Endpoint: POST /v1/feeds
