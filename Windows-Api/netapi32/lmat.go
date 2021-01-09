package netapi32

type AT_ENUM struct {
	JobId       DWORD
	JobTime     DWORD_PTR
	DaysOfMonth DWORD
	DaysOfWeek  UCHAR
	Flags       UCHAR
	Command     LPWSTR
}

type AT_INFO struct {
	JobTime     DWORD_PTR
	DaysOfMonth DWORD
	DaysOfWeek  UCHAR
	Flags       UCHAR
	Command     LPWSTR
}

var (
	procNetScheduleJobAdd     = modNetApi32.NewProc("NetScheduleJobAdd")
	procNetScheduleJobDel     = modNetApi32.NewProc("NetScheduleJobDel")
	procNetScheduleJobEnum    = modNetApi32.NewProc("NetScheduleJobEnum")
	procNetScheduleJobGetInfo = modNetApi32.NewProc("NetScheduleJobGetInfo")
)
