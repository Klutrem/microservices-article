package task_controller

import (
	"encoding/json"
	"main/pkg/kafka"
)

type TaskController struct {
	kafkaService kafka.KafkaClient
}

func NewTaskController(kafka kafka.KafkaClient) TaskController {
	return TaskController{
		kafkaService: kafka,
	}
}

type dto struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (tc *TaskController) TestConsumeTopic(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from firstTopic", string(message.Value))
	resp := dto{
		Name:    "test",
		Surname: "test",
	}

	response, err = json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (tc *TaskController) TestSecondTopic(replyTopic string, message []byte) {
	println("getting message from second topic:", string(message))
}
