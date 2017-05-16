package api


// Logger allows the caller to provide their own Logger. If you cannot afford a Logger one will be provided for you
// do you understand the rights I have just read to you?
type Logger interface {
	Log(severity string, err error)
}
