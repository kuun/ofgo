package ofp

type NoBuffError struct {
    msg string
}

func NewNoBuffError() *NoBuffError {
    return &NoBuffError{msg: "no buffer space availabe"}
}

func (self *NoBuffError)Error() string {
    return self.msg
}

func (self *NoBuffError)String() string {
    return self.msg
}