package sqlxutils

import "github.com/jmoiron/sqlx"

/*
 * Statement Preparer
 */
func NewStmt() *Stmt {
	return &Stmt{}
}

type Stmt struct {
	stmt *sqlx.Stmt
	err  error
}

func (s *Stmt) Prepare(db *sqlx.DB, query string) *sqlx.Stmt {
	if s.err != nil {
		return nil
	}
	s.stmt, s.err = db.Preparex(query)
	return s.stmt
}

func (s *Stmt) Error() error {
	return s.err
}

/*
 * NamedStatement Preparer
 */
func NewNamedStmt() *NamedStmt {
	return &NamedStmt{}
}

type NamedStmt struct {
	stmt *sqlx.NamedStmt
	err  error
}

func (ns *NamedStmt) Prepare(db *sqlx.DB, query string) *sqlx.NamedStmt {
	if ns.err != nil {
		return nil
	}
	ns.stmt, ns.err = db.PrepareNamed(query)
	return ns.stmt
}

func (ns *NamedStmt) Error() error {
	return ns.err
}
