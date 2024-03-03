package interfaceRepository_auth_svc

type IAdminRepository interface{
	GetPassword( string) (string, error)
}