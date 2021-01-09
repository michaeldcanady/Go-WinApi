package netapi32

type MSG_INFO_0 struct {
	msgi0_name LPWSTR
}

type MSG_INFO_1 struct {
	msgi1_name         LPWSTR
	msgi1_forward_flag DWORD
	msgi1_forward      LPWSTR
}

var (
	procNetMessageNameGetInfo = modNetApi32.NewProc("NetMessageNameGetInfo")
	procNetMessageBufferSend  = modNetApi32.NewProc("NetMessageBufferSend")
	procNetMessageNameEnum    = modNetApi32.NewProc("NetMessageNameEnum")
	procNetMessageNameAdd     = modNetApi32.NewProc("NetMessageNameAdd")
	procNetMessageNameDel     = modNetApi32.NewProc("NetMessageNameDel")
)

func NetMessageNameGetInfo() {

}
func NetMessageBufferSend() {

}
func NetMessageNameEnum() {

}
func NetMessageNameAdd() {

}
func NetMessageNameDel() {

}
