package person

import (
	_ "embed"
)

// PostgreSQL
var (
	//go:embed sql/pg/tbl_person.sql
	PGPersonSQL string
	//go:embed sql/pg/tbl_named_id.sql
	PGNamedIDSQL string
	//go:embed sql/pg/tbl_person_name_id.sql
	PGPersonNameID string
)
