package squirrel

import sq "github.com/Masterminds/squirrel"

func basicSelectQuery() (string, []interface{}, error) {
	users := sq.Select("*").From("users").Join("emails USING (email_id_")

	active := users.Where(sq.Eq{"deleted_at": nil})

	return active.ToSql()
}
