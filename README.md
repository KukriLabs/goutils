Helpful common utilities I use across projects.

## Better Error Handling with sqlx Prepared Statements

Package         |   Utility               |
----------------|-------------------------|
`sqlxutils`     | `Stmt` & `NamedStmt`    |

When preparing multiple SQL statements with [`sqlx`](https://github.com/jmoiron/sqlx) error handling can be made less verbose with a single check at the end of all preparations.

### Example

```golang
package main

import (
  "github.com/kukrilabs/goutils/sqlxutils"
)

const (
  stmtOne       = `INSERT INTO ...`
  stmtTwo       = `SELECT * FROM ...`
  namedStmtOne  = `SELECT id FROM ...`
)

var (
  preparedStmtOne       *sqlx.Stmt
  preparedStmtTwo       *sqlx.Stmt
  preparedNamedStmtOne  *sqlx.NamedStmt
)

func main() {
  stmtPreparer := sqlxutils.NewStmt()
  namedStmtPreparer := sqlxutils.NewNamedStmt()

  preparedStmtOne = stmtPreparer.Prepare(stmtOne)
  preparedStmtTwo = stmtPreparer.Prepare(stmtTwo)
  preparedNamedStmtOne = namedStmtPreparer.Prepare(namedStmtOne)

  if stmtPreparer.Error() != nil {
    // One of the statements had an error
  }

  if namedStmtPreparer.Error() != nil {
    // One of the named statements had an error
  }
}
```

## Parse URLs or fail doing so

Package         |   Utility               |
----------------|-------------------------|
`urlutils`      | `MustParseURL`          |

Make sure a URL is successfully parsed or panic doing so.

## Add a `X-Request-ID` header to requests

Package         |   Utility               |
----------------|-------------------------|
`middleware`    | `RequestID`             |

Add a unique identifier to requests (useful for logging etc.) if the Request does not already have one. Access later by using the `RequestIDHeader` variable header name e.g. `request.Header.Get(RequestIDHeader)`

## Database connections with Exponential Back-off

Package         |   Utility               |
----------------|-------------------------|
`sqlxutils`     | `MustConnect`           |
`sqlxutils`     | `NewPostgresConfig`     |

Create a `*sqlx.DB` connection with some default connection information tailored for PostgreSQL with the `NewPostgresConfig` helper.

```golang
db := sqlxutils.MustConnect(sqlxutils.NewPostgresConfig("localhost:5432"))
```

## Easily JSON a HTTP response

Package         |   Utility               |
----------------|-------------------------|
`jsonutils`     | `JSONify`               |

Sets the `Content-Type` header to `application/json`, the relevant response code and marshals the input to JSON binary

## Unified Error Message `struct` for JSON responses

Package         |   Utility               |
----------------|-------------------------|
`jsonutils`     | `ErrorMessage`          |

Useful with the `JSONify` function to provide a uniform error message envelope

## Return a JSON "NotImplemented" response

Package         |   Utility               |
----------------|-------------------------|
`httputils`     | `NotImplemented`        |

Returns a `HTTP 501` response with a JSON "not implemented" message for stub endpoints

## Remove empty/white space strings from an array

Package         |   Utility               |
----------------|-------------------------|
`arrays`        | `TrimEmptyStrings`      |

Takes an input array of strings and returns an array with all empty or white space strings removed.

## Generate a string of a random length

Package         |   Utility               |
----------------|-------------------------|
`stringutils`   | `RandString`            |

Takes an input integer `n` and outputs a random alphanumeric string of length `n`
