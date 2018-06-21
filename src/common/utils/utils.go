package utils

import (
	"../const/userconst"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-nsq"
	"log"
)

type (
	AuthResponse struct {
		FinalReply string            `bson:"finalReply" json:"finalReply"`
		Status     string            `bson:"status" json:"status"`
		Response   userconst.UserDoc `bson:"response" json:"response"`
	}
)

func PrintJSON(v interface{}) string {
	c, _ := json.Marshal(v)
	return fmt.Sprintf("%s", c)
}

func GetSHA1(s string) string {

	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash

}

func PublishMsgToNSQ(topic string, b *[]byte, nsqHost string) error {

	var err error

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(nsqHost, config)
	if err != nil {
		log.Printf(err.Error())
	}

	err = w.Publish(topic, *b)
	if err != nil {
		log.Printf(err.Error())
	}

	w.Stop()

	return err

}

func parseSQLRows(rows sql.Rows) map[int]map[string]string {

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	finalResult := map[int]map[string]string{}
	resID := 0

	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		tmpStruct := map[string]string{}

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			tmpStruct[col] = fmt.Sprintf("%s", v)
		}

		finalResult[resID] = tmpStruct
		resID++
	}

	return finalResult

}
