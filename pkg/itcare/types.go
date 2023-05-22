package itcare

const (
	InstanceType string = "instance"
)

type CI interface {
	getID() int
	getType() string
}
