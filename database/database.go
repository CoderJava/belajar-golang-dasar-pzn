package database

import "fmt"

var connection string

// Otomatis terpanggil ketika package ini di-import oleh file lain (sebelum func main() terpanggil)
func init() {
	connection = "MySQL"
	fmt.Println("init")
}

func GetDatabase() string {
	return connection
}
