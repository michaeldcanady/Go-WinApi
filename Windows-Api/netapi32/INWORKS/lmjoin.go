package netapi32

var (
	procNetRequestProvisioningPackageInstall = modNetApi32.NewProc("NetRequestProvisioningPackageInstall")
	procNetRemoveAlternateComputerName       = modNetApi32.NewProc("NetRemoveAlternateComputerName")
	procNetCreateProvisioningPackage         = modNetApi32.NewProc("NetCreateProvisioningPackage")
	procNetAddAlternateComputerName          = modNetApi32.NewProc("NetAddAlternateComputerName")
	procNetProvisionComputerAccount          = modNetApi32.NewProc("NetProvisionComputerAccount")
	procNetRequestOfflineDomainJoin          = modNetApi32.NewProc("NetRequestOfflineDomainJoin")
	procNetEnumerateComputerNames            = modNetApi32.NewProc("NetEnumerateComputerNames")
	procNetFreeAadJoinInformation            = modNetApi32.NewProc("NetFreeAadJoinInformation")
	procNetSetPrimaryComputerName            = modNetApi32.NewProc("NetSetPrimaryComputerName")
	procNetGetAadJoinInformation             = modNetApi32.NewProc("NetGetAadJoinInformation")
	procNetRenameMachineInDomain             = modNetApi32.NewProc("NetRenameMachineInDomain")
	procNetGetJoinInformation                = modNetApi32.NewProc("NetGetJoinInformation")
	procNetGetJoinableOUs                    = modNetApi32.NewProc("NetGetJoinableOUs")
	procNetUnjoinDomain                      = modNetApi32.NewProc("NetUnjoinDomain")
	procNetValidateName                      = modNetApi32.NewProc("NetValidateName")
	procNetJoinDomain                        = modNetApi32.NewProc("NetJoinDomain")
)

type DSREG_USER_INFO struct {
	pszUserEmail   LPWSTR
	pszUserKeyId   LPWSTR
	pszUserKeyName LPWSTR
}

type DSREG_JOIN_INFO struct {
	joinType              DSREG_JOIN_TYPE
	pJoinCertificate      PCCERT_CONTEXT
	pszDeviceId           LPWSTR
	pszIdpDomain          LPWSTR
	pszTenantId           LPWSTR
	pszJoinUserEmail      LPWSTR
	pszTenantDisplayName  LPWSTR
	pszMdmEnrollmentUrl   LPWSTR
	pszMdmTermsOfUseUrl   LPWSTR
	pszMdmComplianceUrl   LPWSTR
	pszUserSettingSyncUrl LPWSTR
	pUserInfo             *DSREG_USER_INFO
}

type NETSETUP_PROVISIONING_PARAMS struct {
	dwVersion           DWORD
	lpDomain            LPCWSTR
	lpHostName          LPCWSTR
	lpMachineAccountOU  LPCWSTR
	lpDcName            LPCWSTR
	dwProvisionOptions  DWORD
	aCertTemplateNames  *LPCWSTR
	cCertTemplateNames  DWORD
	aMachinePolicyNames *LPCWSTR
	cMachinePolicyNames DWORD
	aMachinePolicyPaths *LPCWSTR
	cMachinePolicyPaths DWORD
	lpNetbiosName       LPWSTR
	lpSiteName          LPWSTR
	lpPrimaryDNSDomain  LPWSTR
}

func NetRequestProvisioningPackageInstall() {
}
func NetRemoveAlternateComputerName() {
}
func NetCreateProvisioningPackage() {
}
func NetAddAlternateComputerName() {
}
func NetProvisionComputerAccount() {
}
func NetRequestOfflineDomainJoin() {
}
func NetEnumerateComputerNames() {
}
func NetFreeAadJoinInformation() {
}
func NetSetPrimaryComputerName() {
}
func NetGetAadJoinInformation() {
}
func NetRenameMachineInDomain() {
}
func NetGetJoinInformation() {
}
func NetGetJoinableOUs() {
}
func NetUnjoinDomain() {
}
func NetValidateName() {
}
func NetJoinDomain() {
}
