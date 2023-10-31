package main

import (
	"encoding/csv"
	"fmt"
	"golang-crud/controller/DataExporter"
	"golang-crud/package/redis"
	"log"
	"os"
)

func main() {
	//db, err := db.DbConnection()
	//if err != nil {
	//	log.Printf("Error %s when getting db connection", err)
	//	return
	//}
	//
	//defer db.Close()
	//
	//log.Printf("Successfully connected to database")

	// Export data to CSV
	err := DataExporter.FetchData()
	if err != nil {
		log.Printf("Error %s when exporting data to CSV", err)
		return
	}

	//CSVWriteAll()

	redisClient, err := redis.RedisConnection()
	if err != nil {
		log.Printf("Failed to connect to redis: %s", err)
		return
	}

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}

func CSVWriteAll() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}
