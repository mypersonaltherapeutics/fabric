package mpc


type Channel interface {

	Send(payload []byte, endpoint string) error

	Receive(timeout int) ([]byte, error)

}


