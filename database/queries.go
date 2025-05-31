package database

// CreateTableQueries contains the SQL queries for creating tables
var CreateTableQueries = []string{
	`CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT false
	)`,
	
	`CREATE INDEX IF NOT EXISTS idx_todos_deleted_at ON todos(deleted_at)`,
}

// CreateTables executes the table creation queries
func CreateTables() error {
	for _, query := range CreateTableQueries {
		if err := DB.Exec(query).Error; err != nil {
			return err
		}
	}
	return nil
} 