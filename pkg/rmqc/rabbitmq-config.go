package rmqc

type RabbitMQConfig struct {
	Username       string
	Password       string
	Host           string
	Port           string
	VHost          string
	ConnectionName string
	Reconnect      Reconnect
}
