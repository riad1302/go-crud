package DataExporter

import (
	"encoding/csv"
	"fmt"
	"golang-crud/package/db"
	"log"
	"os"
	"time"
)

type CommissionRecord struct {
	Id               int
	Msisdn           string
	DeviceId         string
	ActivationDate   string
	InstallationTime string
	Shard            int
	created_at       string
}

func FetchData() error {

	records, err := queryCommissionRecords()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Printf("ID: %d, Device Id: %s,  ActivationDate: %s, Shared: %v\n", record.Id, record.DeviceId, record.ActivationDate, record.Shard)
	}

	if len(records) == 0 {
		fmt.Println("No records to export.")
		return nil
	}

	if err := csvGenerate(records); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data exported to exported_data.csv")

	// Update the Shard field to true in the database
	if err := updateShardFlag(records); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Shard updated to true for exported records.")

	return nil

}

func csvGenerate(records []CommissionRecord) error {
	// Create a CSV file for export
	fileName := generateFileName()
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Create a CSV writer
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// Write the CSV header
	csvWriter.Write([]string{"Timestamp", "MSISDN Number", "Device ID/ADID"})

	for _, record := range records {
		//fmt.Println("dsds", record.Id)
		csvWriter.Write([]string{record.InstallationTime + "/" + record.created_at, record.Msisdn, record.DeviceId})
	}

	return nil
}

func updateShardFlag(records []CommissionRecord) error {
	dbConn, err := db.DbConnection()
	if err != nil {
		return err
	}
	defer dbConn.Close()

	// Prepare an update statement
	updateStmt, err := dbConn.Prepare("UPDATE nrt_commission_records SET shard = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer updateStmt.Close()

	for _, record := range records {
		_, err := updateStmt.Exec(true, record.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func queryCommissionRecords() ([]CommissionRecord, error) {
	db, err := db.DbConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return nil, err
	}
	// Calculate the time one hour ago from the current time
	//oneHourAgo := time.Now().Add(-time.Hour)
	//
	//fmt.Println(oneHourAgo)

	//query := "SELECT id, msisdn, device_id, activation_date, installation_time, shard, created_at FROM nrt_commission_records WHERE created_at BETWEEN ? AND NOW() AND shard = ?"
	query := "SELECT id, msisdn, device_id, activation_date, installation_time, shard, created_at FROM nrt_commission_records WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 HOUR) AND shard = ?"
	rows, err := db.Query(query, false)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []CommissionRecord

	for rows.Next() {
		var record CommissionRecord
		if err := rows.Scan(&record.Id, &record.Msisdn, &record.DeviceId, &record.ActivationDate, &record.InstallationTime, &record.Shard, &record.created_at); err != nil {
			return nil, err
		}
		results = append(results, record)
	}

	return results, nil
}

func generateFileName() string {
	timestamp := time.Now().Format("20060102_150405") // Format the current time as "yyyyMMdd_HHMMSS"
	return fmt.Sprintf("csv/exported_data_%s.csv", timestamp)
}
