package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const bucketSize float64 = 2000.0

type bucket struct {
	Volume float64
	Items  [][]string
}

func main() {
	csvFile, _ := os.Open(os.Args[1])
	r := csv.NewReader(csvFile)

	//Track the current bucket
	currentBucket := 0
	//build a list of all buckets
	var buckets []bucket

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for len(records) > 0 {
		for i := 0; i < len(records); i++ {
			record := records[i]
			//Initialize a new bucket if necessary.
			if len(buckets) <= currentBucket {
				buckets = append(buckets, bucket{
					Volume: 0,
				})
			}

			trimmed := strings.Trim(record[1], " ")
			weight, err := strconv.ParseFloat(trimmed, 64)

			if err != nil {
				log.Fatal(err)
			}

			theBucket := buckets[currentBucket]
			currentVolume := theBucket.Volume
			newVolume := weight + currentVolume
			//If the item fits, put it in the bucket
			//and remove it from the candidate list.
			if (newVolume) <= bucketSize {
				theBucket.Volume = newVolume
				theBucket.Items = append(theBucket.Items, record)
				records = append(records[:i], records[i+1:]...)
				buckets[currentBucket] = theBucket
				i-- //manually decrement i so we don't out of bounds.
			}
		}
		currentBucket = currentBucket + 1
	}

	fmt.Printf("Create %d new bins\n", len(buckets))
	for i, bucket := range buckets {
		fmt.Printf("Bin %d: %f\n", i, bucket.Volume)
		for _, item := range bucket.Items {
			fmt.Println(item)
		}
	}

}
