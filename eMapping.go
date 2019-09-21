package flaw

var (
	errMap map[int]string
)

//error code
const (
	CustomCode = iota + 1
	UnmarshalCode
	MarshalCode
	BadReqCode
	ForbiddenCode
	NotFoundCode
	MiscCode
	ConnCode
	TxCode
	CreateCode
	SelectCode
	InsertCode
	UpdateCode
	DeleteCode
)

//msg constant
const (
	InvalidReqMsg  = "Invalid Request, Please contact system adminstrator for further clarification."
	ForbiddenMsg   = "You are not allowed to perform this operation. Please contact system adminstrator."
	BlandReqMsg    = "Blank request, Please provide input to process"
	NotFoundMsg    = "Record not found"
	ServerErrorMsg = "Sorry unable to process this request. Please Try Again"
)

func init() {
	errMap = map[int]string{
		UnmarshalCode: InvalidReqMsg,
		MarshalCode:   InvalidReqMsg,
		BadReqCode:    InvalidReqMsg,
		ForbiddenCode: ForbiddenMsg,
		NotFoundCode:  NotFoundMsg,
		MiscCode:      ServerErrorMsg,
		ConnCode:      ServerErrorMsg,
		TxCode:        ServerErrorMsg,
		CreateCode:    ServerErrorMsg,
		SelectCode:    ServerErrorMsg,
		InsertCode:    ServerErrorMsg,
		UpdateCode:    ServerErrorMsg,
		DeleteCode:    ServerErrorMsg,
	}
}
