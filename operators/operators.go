package operators

const (
	AND = "AND"
	OR  = "OR"
)

const (
	EQUALS              = "eq"
	NOT_EQUALS          = "neq"
	GREATER_THAN        = "gt"
	GREATER_THAN_EQUALS = "gte"
	LESS_THAN           = "lt"
	LESS_THAN_EQUAL     = "lte"
	LIKE                = "like"
	NOT_LIKE            = "nlike"
	IGNORE_LIKE         = "ilike"
	NOT_IGNORE_LIKE     = "nilike"
	IN                  = "in"
	NOT_IN              = "nin"
	IS                  = "is"
	IS_NOT              = "isnot"
	BETWEEN             = "between"
	NOT_BETWEEN         = "nbetween"
	OVERLAP             = "overlap"
	CONTAINS            = "contains"
	CONTAINED           = "contained"
)

var SQL = map[string]string{
	EQUALS:              "=",
	NOT_EQUALS:          "<>",
	GREATER_THAN:        ">",
	GREATER_THAN_EQUALS: ">=",
	LESS_THAN:           "<",
	LESS_THAN_EQUAL:     "<=",
	LIKE:                "LIKE",
	NOT_LIKE:            "NOT LIKE",
	IGNORE_LIKE:         "iLIKE",
	NOT_IGNORE_LIKE:     "NOT iLIKE",
	IN:                  "IN",
	NOT_IN:              "NOT IN",
	IS:                  "=",
	IS_NOT:              "<>",
	BETWEEN:             "IN",
	NOT_BETWEEN:         "NOT IN",
	OVERLAP:             "=",
	CONTAINS:            "LIKE",
	CONTAINED:           "LIKE",
}
