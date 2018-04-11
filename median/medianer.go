package median

// Medianer wraps median functionality
type Medianer interface {
	AddDelay(int)
	GetMedian() float64
}
