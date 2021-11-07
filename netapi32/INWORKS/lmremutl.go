package netapi32

type TIME_OF_DAY_INFO struct {
	tod_elapsedt  DWORD
	tod_msecs     DWORD
	tod_hours     DWORD
	tod_mins      DWORD
	tod_secs      DWORD
	tod_hunds     DWORD
	tod_timezone  LONG
	tod_tinterval DWORD
	tod_day       DWORD
	tod_month     DWORD
	tod_year      DWORD
	tod_weekday   DWORD
}

var (
	procNetRemoteComputerSupports = modNetApi32.NewProc("NetRemoteComputerSupports")
	procNetRemoteTOD              = modNetApi32.NewProc("NetRemoteTOD")
)

func NetRemoteComputerSupports() {

}

func NetRemoteTOD() {

}
