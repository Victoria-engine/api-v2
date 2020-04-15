# Victoria API

![Build](https://github.com/Victoria-engine/api-v2/workflows/Go/badge.svg)

## Install & Running the API

Make sure you have Go and Make installed on your machine.

Victoria API uses `postgres` for the database, so make sure you have instance of a postgres db running locally and set the creditials on the `.env.sample`.

After that rename the file to `.env` and fill the rest of the variables.

Then run, to run the API locally:

```bash
make run
```

## Testing

```bash
make test
```

## Contribution
0. (Issues from the Github project board are welcome for everyone)
1. Create a branch and a Pull Request with the changes or Fork the project.
2. Make sure the issue is consise.
3. Wait for review and thank for contributing!


### Creating a new service:

The architecture of the API is inspired by the [Gorsk](https://github.com/ribice/gorsk) with some changes.

The Victoria api is divided into small unit services that represent at most tables in the database.
Each is made by

```
   |- post             (Service name, equal to the database table)
      |- transport     
         | http.go     (http methods and routes for this service)
      |- repository    
         | post.go     (post repository that talks to the db)
      | post.go        (services implementation)
      | presenter.go   (HTTP response data for the client)
      | service.go     (interface for the service and repository)
```





