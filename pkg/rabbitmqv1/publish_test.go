package rabbitmqv1

import (
	"testing"
)

type ExampleMessage struct {
	Id string
}

func TestGetBytes(t *testing.T) {
	var comlexType = ExampleMessage{Id: "1"}
	var _, err = getBytes(comlexType)
	if err != nil {
		t.Errorf("Test Fail : Actual [%v]\tExptected [%v]\n", nil, err.Error())
	}

}
