package repositories

// RepositoryResult : Data result
type RepositoryResult struct {
	Value        interface{} `json:"value"`
	Error        error       `json:"error"`
	RowsAffected int64       `json:"rows_affected"`
}
