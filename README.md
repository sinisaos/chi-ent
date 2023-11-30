Example application using [Ent](https://entgo.io/). Library is easy to use and really powerful (especially ``FK`` and ``M2M ``relations) with code generation.

## How to use

Clone the repo.
```
$ git clone https://github.com/sinisaos/chi-ent.git
$ cd chi-ent
$ cp .env.example .env && rm .env.example # WARNING: change the DSN to your credentials and database name
$ make server
```

## Generate the schema
To generate changes in the Ent schema.
```
$ make generate
```

## Run tests
To run tests.
```
$ make tests
```

After application is running you can visit ``localhost:8000/swagger/`` and use interactive Swagger documentation which was generated using [Entoas](https://github.com/ent/contrib/tree/master/entoas), with minor modifications to use authorization.

## Service layer

This application uses the [Chi](https://github.com/go-chi/chi) router, but since all database interaction is in the [service layer](https://github.com/sinisaos/chi-ent/tree/main/pkg/service), you can use any other router or framework for web part.