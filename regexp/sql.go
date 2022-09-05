package regexp

const (
	ExpPreFixICase     = `(?i)`
	ExpSingleCondition = `\s*(=|in)\s*[\('"]{1}(.*?)[\)'"]{1}(([;\s\t]*$)|(\s+(AND|GROUP|ORDER|LIMIT)\s+))`
	ExpParseEntireSql  = `(?i)SELECT\s+(.*?)\s+FROM\s+(.*?)(([;\s\t]*$)|(\s+(WHERE\s+(.+)\s+(([;\s\t]*$)|((GROUP BY|ORDER BY|LIMIT)\s+)))))`
	ExpTableName       = `(?i)SELECT\s+.*?\s+FROM\s+(.*?)(([;\s\t]*$)|(\s+(WHERE|GROUP BY|ORDER BY|LIMIT)\s+))`
	ExpSelectFields    = `(?i)SELECT\s+(.*?)(([;\s\t]*$)|(\s+FROM\s+))`
	ExpWhere           = `(?i)WHERE\s+(.+?)\s+(([;\s\t]*$)|((GROUP BY|ORDER BY|LIMIT)\s+))`
)

func init() {
}

func MatchSqlWhere(sql string) (string, error) {
	sql
}

func SearchWhereCondition(field string) (value interface{}) {
	getExp := func(field string) string {
		return ExpPreFixICase + field + ExpCondition
	}

	return
}
