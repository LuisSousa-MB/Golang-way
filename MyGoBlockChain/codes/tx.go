package codes

type ITransaction interface {
	GetMsg() string
	SetMsg(string)
	GetTimeStamp() string
	SetTimeStamp(string)
	GetValue() float64
	SetValue(float64)
}

type Transaction struct {
	Msg, TimeStamp, Sender, Receiver, Hash string
	Value                                  float64
}

type Data struct {
	Transactions []Transaction
}
