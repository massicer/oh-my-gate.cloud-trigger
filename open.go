package helloworld

import (
        "encoding/json"
        "context"
        "cloud.google.com/go/pubsub"
        "fmt"
        "net/http"
        "os"
        "strconv"
        "log"
)

const OPEN_ACTION = "Open"

type OpenMessage struct {
	Id int64
	Action string
}

func Open(w http.ResponseWriter, r *http.Request) {

        ctx := context.Background()

        var topic_name string = os.Getenv("TOPIC_NAME")
        var project_id string =  os.Getenv("GCP_PROJECT_ID")

        client, err := pubsub.NewClient(ctx, project_id)
	if err != nil {
                erMsg := fmt.Sprintf("Could not create client: %v", err)
                return_response(w, erMsg, 500)
                return
        }
        log.Print("Client created")

        topic := client.Topic(topic_name)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
                erMsg := fmt.Sprintf("Could not check if topic exists: %v", err)
                log.Print(erMsg)
                return_response(w, erMsg, 500)
                return
	}else 
	        if !exists {
                        log.Printf("Topic %v doesn't exist - creating it", topic_name)
                        _, err = client.CreateTopic(ctx, topic_name)
                        if err != nil {
                                erMsg := fmt.Sprintf("Could not create topic: %v", err)
                                log.Print(erMsg)
                                return_response(w, erMsg, 500)
                                return
                        }
        }
        log.Print("Topic exists")

        pin_param := r.URL.Query().Get("pin")
        log.Print("Pin param extracted: ", pin_param)
        var gate_id, _ = strconv.ParseInt(pin_param, 10, 32)
        var pub_msg = OpenMessage{
                Id: gate_id,
                Action: OPEN_ACTION,
        }
        encoded_msg, _ := json.Marshal(pub_msg)
        var msg =  &pubsub.Message{
                Data: []byte(encoded_msg),
        }
        
        if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
                log.Printf("Message %v unable to publish to topic %s", msg, topic_name)
		return_response(w, fmt.Sprintf("Could not publish message: %v", err), 500)
	}else {
                log.Printf("Message %v published to topic %s", msg, topic_name)
                return_response(w, "Opened", 200)
        }
}
        
      

func return_response(w http.ResponseWriter, message string, status int) {
        w.WriteHeader(status)
        fmt.Fprint(w, message)
}
