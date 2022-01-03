package winapi

import(

)

const(
  //Errors
	NO_ERROR				                       = 0x00000000
	ERROR_SESSION_CREDENTIAL_CONFLICT	     = 0x000004C3
  ERROR_BAD_USERNAME                     = 0x0000089A
  ERROR_NOT_SUPPORTED                    = 0x00000032
  ERROR_CANCELLED                        = 0x000004c7
  ERROR_RETRY                            = 0x000004d5
  ERROR_MORE_DATA                        = 0x000000ea
  ERROR_INVALID_ADDRESS                  = 0x000001e7
  ERROR_INVALID_PARAMETER                = 0x00000057
  ERROR_INVALID_PASSWORD                 = 0x00000056
  ERROR_ACCESS_DENIED                    = 0x00000005
  ERROR_BUSY                             = 0x000000aa
  ERROR_UNEXP_NET_ERR                    = 0x0000003b
  ERROR_NOT_ENOUGH_MEMORY                = 0x00000008
  ERROR_NO_NETWORK                       = 0x000004c6
  ERROR_EXTENDED_ERROR                   = 0x000004b8
  ERROR_INVALID_LEVEL                    = 0x0000007c
  ERROR_INVALID_HANDLE                   = 0x00000006
  ERROR_ALREADY_INITIALIZED              = 0x000004df
  ERROR_NO_MORE_DEVICES                  = 0x000004e0
  // Connections
  ERROR_NOT_CONNECTED                    = 0x000008ca
  ERROR_OPEN_FILES                       = 0x00000961
  ERROR_DEVICE_IN_USE                    = 0x00000964
  ERROR_BAD_NET_NAME                     = 0x00000043
  ERROR_BAD_DEVICE                       = 0x000004b0
  ERROR_ALREADY_ASSIGNED                 = 0x00000000
  ERROR_GEN_FAILURE                      = 0x00000055
  ERROR_CONNECTION_UNAVAIL               = 0x000004b1
  ERROR_NO_NET_OR_BAD_PATH               = 0x000004b3
  ERROR_BAD_PROVIDER                     = 0x000004b4
  ERROR_CANNOT_OPEN_PROFILE              = 0x000004b5
  ERROR_BAD_PROFILE                      = 0x000004b6
  ERROR_BAD_DEV_TYPE                     = 0x00000042
  ERROR_DEVICE_ALREADY_REMEMBERED        = 0x000004b2
  ERROR_CONNECTED_OTHER_PASSWORD         = 0x0000083c
  ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT = 0x0000083d
)
