package netapi32

const (
	NET_API_STATUS_NERR_Success                      = 0
	NET_API_STATUS_NERR_InvalidComputer              = 2351
	NET_API_STATUS_NERR_NotPrimary                   = 2226
	NET_API_STATUS_NERR_SpeGroupOp                   = 2234
	NET_API_STATUS_NERR_LastAdmin                    = 2452
	NET_API_STATUS_NERR_BadPassword                  = 2203
	NET_API_STATUS_NERR_PasswordTooShort             = 2245
	NET_API_STATUS_NERR_UserNotFound                 = 2221
	NET_API_STATUS_ERROR_ACCESS_DENIED               = 5
	NET_API_STATUS_ERROR_NOT_ENOUGH_MEMORY           = 8
	NET_API_STATUS_ERROR_INVALID_PARAMETER           = 87
	NET_API_STATUS_ERROR_INVALID_NAME                = 123
	NET_API_STATUS_ERROR_INVALID_LEVEL               = 124
	NET_API_STATUS_ERROR_MORE_DATA                   = 234
	NET_API_STATUS_ERROR_SESSION_CREDENTIAL_CONFLICT = 1219
	NET_API_STATUS_RPC_S_SERVER_UNAVAILABLE          = 2147944122
	NET_API_STATUS_RPC_E_REMOTE_DISABLED             = 2147549468

	USER_PRIV_MASK = 0x3

	//User priveledge level
	USER_PRIV_GUEST = 0
	USER_PRIV_USER  = 1
	USER_PRIV_ADMIN = 2
)
