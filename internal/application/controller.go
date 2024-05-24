package application

import (
	"main/pkg/kafka"
	"time"
)

type TaskController struct{}

func NewTaskController(kafka kafka.KafkaClient) TaskController {
	return TaskController{}
}

func (tc *TaskController) TestConsumeTopic(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from firstTopic", string(message.Value))
	time.Sleep(time.Second)

	return []byte("response"), nil
}

func (tc *TaskController) TestSecondTopic(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from second topic:", string(message.Value))
	time.Sleep(2 * time.Second)
	return []byte("response"), nil
}

func (tc *TaskController) TestThirdTopic(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from third topic:", string(message.Value))
	time.Sleep(3 * time.Second)
	return []byte("response"), nil
}

func (tc *TaskController) TestTopic4(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 4 Topic", string(message.Value))
	time.Sleep(time.Second)

	return []byte("response"), nil
}

func (tc *TaskController) TestTopic5(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 5 topic:", string(message.Value))
	time.Sleep(2 * time.Second)
	return []byte("response"), nil
}

func (tc *TaskController) TestTopic6(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 6 topic:", string(message.Value))
	time.Sleep(3 * time.Second)
	return []byte("response"), nil
}

func (tc *TaskController) TestTopic7(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 7 Topic", string(message.Value))
	time.Sleep(time.Second)

	return []byte("response"), nil
}

func (tc *TaskController) TestTopic8(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 8 topic:", string(message.Value))
	time.Sleep(2 * time.Second)
	return []byte("response"), nil
}

func (tc *TaskController) TestTopic9(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 9 topic:", string(message.Value))
	time.Sleep(3 * time.Second)
	return []byte("response"), nil
}

func (tc *TaskController) TestTopic10(message kafka.KafkaMessage) (response []byte, err error) {
	println("getting message from 10 Topic", string(message.Value))
	time.Sleep(time.Second)

	return []byte("response"), nil
}
