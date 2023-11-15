package module

// Application Configuration
type Configuration struct {
	Database struct {
		Host     string
		User     string
		Password string
		Dbname   string
		Port     int
	}
}
