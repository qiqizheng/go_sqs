/*
   Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/
// snippet-start:[sqs.go.receive_messages]
package main

// snippet-start:[sqs.go.receive_messages.imports]
import (
	// "flag"
	// "bytes"
	"fmt"
	// "io/ioutil"
	// "net/http"
	// "net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// snippet-end:[sqs.go.receive_messages.imports]

// GetQueueURL gets the URL of an Amazon SQS queue
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	queueName is the name of the queue
//
// Output:
//
//	If success, the URL of the queue and nil
//	Otherwise, an empty string and an error from the call to
func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	// snippet-start:[sqs.go.receive_messages.queue_url]
	svc := sqs.New(sess)

	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	// snippet-end:[sqs.go.receive_messages.queue_url]
	if err != nil {
		return nil, err
	}

	return urlResult, nil
}

// GetMessages gets the messages from an Amazon SQS queue
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	queueURL is the URL of the queue
//	timeout is how long, in seconds, the message is unavailable to other consumers
//
// Output:
//
//	If success, the latest message and nil
//	Otherwise, nil and an error from the call to ReceiveMessage
func GetMessages(sess *session.Session, queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	// snippet-start:[sqs.go.receive_messages.call]
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   timeout,
	})
	// snippet-end:[sqs.go.receive_messages.call]
	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func DeleteMessage(sess *session.Session, queueURL *string, messageHandle *string) error {
	// snippet-start:[sqs.go.delete_message.call]
	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})
	// snippet-end:[sqs.go.delete_message.call]
	if err != nil {
		return err
	}

	return nil
}

func postShushu() {

	arr := []string{"hello", "world"}
	// arr["st"] = "uiui"
	str := strings.Join(arr, "&")
	fmt.Println(str) //

	fmt.Println(str)
	// var param [10]string{"appid": "ffbdeea3aba34e1aa5d9c406aee7be33"}
	// jsonStr := []byte(`{
	// 	"appid": "ffbdeea3aba34e1aa5d9c406aee7be33",
	// 	"debug": 1,
	// 	"data": {
	// 	  "#type": "track",
	// 	  "#event_name": "test",
	// 	  "#time": "2019-11-15 11:35:53.648",
	// 	  "properties": { "a": "123", "b": 2 },
	// 	  "#distinct_id": "1111"
	// 	}
	//   }`)

	// escapeUrl := url.QueryEscape(param)
	// // jsonStr := []byte(`{ "username": "auto", "password": "auto123123" }`)
	// url := "https://tareport.xiaoyouxiqun.com/sync_data"
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(escapeUrl))
	// req.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	// handle error
	// }
	// defer resp.Body.Close()
	// statuscode := resp.StatusCode
	// hea := resp.Header
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// fmt.Println(statuscode)
	// fmt.Println(hea)
}

func main() {

	postShushu()
	// snippet-start:[sqs.go.receive_messages.args]
	// queue := flag.String("q", "test.fifo", "The name of the queue")
	// timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	// flag.Parse()

	// if *queue == "" {
	// 	fmt.Println("You must supply the name of a queue (-q QUEUE)")
	// 	return
	// }

	// if *timeout < 0 {
	// 	*timeout = 0
	// }

	// if *timeout > 12*60*60 {
	// 	*timeout = 12 * 60 * 60
	// }
	// // snippet-end:[sqs.go.receive_messages.args]

	// // Create a session that gets credential values from ~/.aws/credentials
	// // and the default region from ~/.aws/config
	// // snippet-start:[sqs.go.receive_messages.sess]
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	// // snippet-end:[sqs.go.receive_messages.sess]

	// // Get URL of queue
	// urlResult, err := GetQueueURL(sess, queue)
	// if err != nil {
	// 	fmt.Println("Got an error getting the queue URL:")
	// 	fmt.Println(err)
	// 	return
	// }

	// // snippet-start:[sqs.go.receive_message.url]
	// queueURL := urlResult.QueueUrl
	// // snippet-end:[sqs.go.receive_message.url]

	// msgResult, err := GetMessages(sess, queueURL, timeout)
	// if err != nil {
	// 	fmt.Println("Got an error receiving messages:")
	// 	fmt.Println(err)
	// 	return
	// }
	// // fmt.Println(*msgResult.Messages[0].ReceiptHandle)
	// var ReceiptHandle string = *msgResult.Messages[0].ReceiptHandle

	// // fmt.Println(msgResult.Messages[0])

	// // fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)

	// // snippet-start:[sqs.go.receive_messages.print_handle]
	// // fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	// // snippet-end:[sqs.go.receive_messages.print_handle]

	// if ReceiptHandle != "" {
	// 	// messageHandle := flag.String("m", "AQEBU78+R/6vRukQ2Fg/wDE337CUqMwiqtsnf0BDg1oeOMv/dXWoTYGUz7td4QUmpb2+jXGc1ZRZFeNHgKAK3ZuEF2cLBwmEanBcvlpfYxk01xTtJmXZ7FTx4Ne9d1LOdUuIWWRAre8v3whGwM6CVGz5hSi15etxcQC13/ySx17o1cVmPl+9gSlid834/bqP2p9DI9Fg4JD+59xt2outqKJ92cgS5dKxxF1rnnT/qYcwLO4c7zAX/OBgbtGedB6aUn7Le13Szhz+4TAQA7iCiMh5aQ==", "The receipt handle of the message")

	// 	// var messageHandle string = new(string)
	// 	// messageHandle = ReceiptHandle
	// 	// fmt.Printf("%T", ReceiptHandle)
	// 	// fmt.Printf("%T", messageHandle)
	// 	// flag.Parse()

	// 	fmt.Println("delete Message " + *msgResult.Messages[0].ReceiptHandle)
	// 	DeleteMessage(sess, queueURL, &ReceiptHandle)

	// 	if err != nil {
	// 		fmt.Println("Got an error deleting the message:")
	// 		fmt.Println(err)
	// 		return
	// 	}

	// }

}

// snippet-end:[sqs.go.receive_messages]
