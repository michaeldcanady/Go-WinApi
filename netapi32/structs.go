package netapi32

import "time"

type USER_INFO_1 struct {
	usri1_name         *uint16
	usri1_password     *uint16
	usri1_password_age uint32
	usri1_priv         uint32
	usri1_home_dir     *uint16
	usri1_comment      *uint16
	usri1_flags        uint32
	usri1_script_path  *uint16
}

type LocalUser1 struct {
	Username      string
	PasswordAge   time.Duration
	Priviledge    string
	HomeDirectory string
	Comment       string
	Flags         []string
	ScriptPath    string
}

type USER_INFO_2 struct {
	Usri2_name           *uint16
	Usri2_password       *uint16
	Usri2_password_age   uint32
	Usri2_priv           uint32
	Usri2_home_dir       *uint16
	Usri2_comment        *uint16
	Usri2_flags          uint32
	Usri2_script_path    *uint16
	Usri2_auth_flags     uint32
	Usri2_full_name      *uint16
	Usri2_usr_comment    *uint16
	Usri2_parms          *uint16
	Usri2_workstations   *uint16
	Usri2_last_logon     uint32
	Usri2_last_logoff    uint32
	Usri2_acct_expires   uint32
	Usri2_max_storage    uint32
	Usri2_units_per_week uint32
	Usri2_logon_hours    uintptr
	Usri2_bad_pw_count   uint32
	Usri2_num_logons     uint32
	Usri2_logon_server   *uint16
	Usri2_country_code   uint32
	Usri2_code_page      uint32
}
