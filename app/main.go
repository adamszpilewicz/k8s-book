package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"log"
	"net/http"
	"os"
)

var (
	appPort = 8888
)

// Person struct represents the structure of the JSON object.
type Person struct {
	Name    string  `parquet:"name=name, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY" json:"name"`
	Age     int32   `parquet:"name=age, type=INT32" json:"age"`
	Active  bool    `parquet:"name=active, type=BOOLEAN" json:"active"`
	Height  float32 `parquet:"name=height, type=FLOAT" json:"height"`
	Country string  `parquet:"name=country, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY" json:"country"`
}

func writeParquet(people []Person, filename string) error {
	fw, err := local.NewLocalFileWriter(filename)
	if err != nil {
		return err
	}
	defer fw.Close()

	pw, err := writer.NewParquetWriter(fw, new(Person), 4)
	if err != nil {
		return err
	}
	defer pw.WriteStop()

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for _, person := range people {
		if err = pw.Write(person); err != nil {
			return err
		}
	}

	return nil
}

func setupRoutes(r *gin.Engine) {
	r.GET("/status", func(c *gin.Context) {
		c.String(200, "healthy")
	})

	r.GET("/", func(c *gin.Context) {
		envVar := os.Getenv("MY_ENV_VAR")
		c.String(200, "ENV VARIABLE: "+envVar)
	})

	r.POST("/parquet", func(c *gin.Context) {
		var people []Person
		if err := c.ShouldBindJSON(&people); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := writeParquet(people, "tmp.parquet"); err != nil {
			log.Println("Failed to write parquet file", err)
			return
		}

		data, err := os.ReadFile("tmp.parquet")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the parquet file."})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename=converted.parquet")
		c.Data(http.StatusOK, "application/octet-stream", data)

		// Delete the temporary file
		if err := os.Remove("tmp.parquet"); err != nil {
			log.Println("Failed to remove temporary file", err)
		}
	})
}

func main() {
	r := gin.Default()
	setupRoutes(r)

	log.Printf("Starting server on port %d\n", appPort)
	portGin := fmt.Sprintf(":%d", appPort)
	if err := r.Run(portGin); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
