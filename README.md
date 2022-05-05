# MOVIESCREEN

An API service for fetching movie data.

The movie data are not meant to be accurate and were generated with https://mockaroo.com/

Codebase shows a demonstration of domain-driven design in Go.

## Usage

Before running or building the, create a `dotenv` file in the root directory of the project:

```shell
touch .env
```

The dotenv file is used to load environment variables into the build and have the form:

```shell
VARIABLE_NAME=value
```

Alternatively, the required values can be passed in via flags (though the dotenv file must still exist, even if empty).

Available flags and their alternative dotenv variable names can be gotten by running:

```shell
make help/api
```

***Note:***
***The following flags (or their dotenv counterparts) are required for the build to run:***
<li> -db-dsn </li>
<li> -smtp-host </li>
<li> -smtp-user </li>

<br>

Run `make help` to view the available rules for running, building and general operations.
