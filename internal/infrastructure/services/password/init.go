package password

const (
	saltLength = 16
	keyLength  = 32
	iterations = 1
	memory     = 64 << 10
	threads    = 4
)

type service struct {
}

func New() *service {
	return &service{}
}
