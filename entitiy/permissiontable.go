package permissiontable

type PermissionTable struct{

	 Element map[string]int 
}

func CreatePermissionTable() (permissionTable *PermissionTable){
	permissionTable = &PermissionTable{nil}
	permissionTable.Element = make(map[string]int)
	permissionTable.Element["GET"] = 1
	permissionTable.Element["POST"] = 2
	permissionTable.Element["PUT"] = 4
	permissionTable.Element["DELETE"] = 8
	permissionTable.Element["CONNECT"] = 16
	permissionTable.Element["HEAD"] = 32
	permissionTable.Element["PATCH"] = 64
	permissionTable.Element["OPTION"] = 128
	permissionTable.Element["TRACE"] = 256
	return permissionTable
}
func (permissionTable PermissionTable) GetCodeNumber(permissions []string) int64 {
	var sum int64 = 0
	for i := 0; i < len(permissions); i++ {
		sum = sum + permissionTable.Element[permissions[i]]
	}
	return sum
}
func (permissionTable PermissionTable) GetPermissions(codeNumber int64) string {

	var permissionsCodeStrr string = strconv.FormatInt(codeNumber, 2)
	var permissionsCodeInt int = strconv.Atoi(permissionsCode)

	fmt.Println("num, ", permissionsCodeInt)

	return permissionsCode
}
