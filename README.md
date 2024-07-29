# GO RSS AGGREGATOR

## Progress

### Set up

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
