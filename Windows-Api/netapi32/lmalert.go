package natapi32

type ADMIN_OTHER_INFO struct {
	alrtad_errcode    DWORD
	alrtad_numstrings DWORD
}

type ERRLOG_OTHER_INFO struct {
	alrter_errcode DWORD
	alrter_offset  DWORD
}

type PRINT_OTHER_INFO struct {
	alrtpr_jobid     DWORD
	alrtpr_status    DWORD
	alrtpr_submitted DWORD
	alrtpr_size      DWORD
}

type STD_ALERT struct {
}

type USER_OTHER_INFO struct {
	alrtus_errcode    DWORD
	alrtus_numstrings DWORD
}

var (
	proNetAlertRaiseEx = modNetApi32.NewProc("NetAlertRaiseEx")
	procNetAlertRaise  = modNetApi32.NewProc("NetAlertRaise")
)

func NetAlertRaiseEx() {

}

func NetAlertRaise() {
  
}
