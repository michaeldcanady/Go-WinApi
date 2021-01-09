package netapi32

type SERVER_INFO_100 struct {
	sv100_platform_id DWORD
	sv100_name        LMSTR
}

type SERVER_INFO_101 struct {
	sv101_platform_id   DWORD
	sv101_name          LMSTR
	sv101_version_major DWORD
	sv101_version_minor DWORD
	sv101_type          DWORD
	sv101_comment       LMSTR
}

type SERVER_INFO_102 struct {
	sv102_platform_id   DWORD
	sv102_name          LMSTR
	sv102_version_major DWORD
	sv102_version_minor DWORD
	sv102_type          DWORD
	sv102_comment       LMSTR
	sv102_users         DWORD
	sv102_disc          LONG
	sv102_hidden        BOOL
	sv102_announce      DWORD
	sv102_anndelta      DWORD
	sv102_licenses      DWORD
	sv102_userpath      LMSTR
}

type SERVER_INFO_1005 struct {
	sv1005_comment LMSTR
}

type SERVER_INFO_1010 struct {
	sv1010_disc LONG
}

type SERVER_INFO_1016 struct {
	sv1016_hidden BOOL
}

type SERVER_INFO_1017 struct {
	sv1017_announce DWORD
}

type SERVER_INFO_1018 struct {
	sv1018_anndelta DWORD
}

type SERVER_INFO_1107 struct {
	sv1107_users DWORD
}

type SERVER_INFO_1501 struct {
	sv1501_sessopens DWORD
}

type SERVER_INFO_1502 struct {
	sv1502_sessvcs DWORD
}

type SERVER_INFO_1503 struct {
	sv1503_opensearch DWORD
}

type SERVER_INFO_1506 struct {
	sv1506_maxworkitems DWORD
}

type SERVER_INFO_1509 struct {
	sv1509_maxrawbuflen DWORD
}

type SERVER_INFO_1510 struct {
	sv1510_sessusers DWORD
}

type SERVER_INFO_1511 struct {
	sv1511_sessconns DWORD
}

type SERVER_INFO_1512 struct {
	SERVER_INFO_1512 DWORD
}

type SERVER_INFO_1513 struct {
	sv1513_maxpagedmemoryusage DWORD
}

type SERVER_INFO_1515 struct {
	sv1515_enableforcedlogoff BOOL
}

type SERVER_INFO_1516 struct {
	sv1516_timesource BOOL
}

var (
	procNetServerComputerNameAdd = modNetApi32.NewProc("NetServerComputerNameAdd")
	procNetServerComputerNameDel = modNetApi32.NewProc("NetServerComputerNameDel")
	procNetServerDiskEnum        = modNetApi32.NewProc("NetServerDiskEnum")
	procNetServerEnum            = modNetApi32.NewProc("NetServerEnum")
	procNetServerGetInfo         = modNetApi32.NewProc("NetServerGetInfo")
	procNetServerSetInfo         = modNetApi32.NewProc("NetServerSetInfo")
	procNetServerTransportAdd    = modNetApi32.NewProc("NetServerTransportAdd")
	procNetServerTransportAddEx  = modNetApi32.NewProc("NetServerTransportAddEx")
	procNetServerTransportDel    = modNetApi32.NewProc("NetServerTransportDel")
	procNetServerTransportEnum   = modNetApi32.NewProc("NetServerTransportEnum")
)
