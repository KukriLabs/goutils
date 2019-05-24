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
