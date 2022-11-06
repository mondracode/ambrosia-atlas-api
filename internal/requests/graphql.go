package requests

type GraphQLQuery struct {
	Query     string                 `json:"query" binding:"required"`
	Variables map[string]interface{} `json:"variables"`
}
