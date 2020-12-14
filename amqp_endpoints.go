package users_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type AMQPEndpointFactory struct {
	userService UserService
}

type ErrorSt struct {
	Text string `json:"text"`
}

func NewAMQPEndpointFactory(userService UserService) *AMQPEndpointFactory {
	return &AMQPEndpointFactory{userService: userService}
}
func (fac *AMQPEndpointFactory) CreateUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateUserCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(&ErrorSt{
				err.Error(),
			})
		}
		resp, err := cmd.Exec(fac.userService)
		if err != nil {
			return AMQPError(&ErrorSt{
				err.Error(),
			})
		}
		return OK(resp)
	}
}

func OK(d interface{}) *amqp.Message {
	data, _ := json.Marshal(d)
	return &amqp.Message{Body: data}
}

func AMQPError(e interface{}) *amqp.Message {
	errObj, _ := json.Marshal(e)
	return &amqp.Message{Body: errObj}
}
