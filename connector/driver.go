package connector

const (
	delay = iota
	cred_username
	cred_password
)

type Driver interface {
	Configuration() string
}
