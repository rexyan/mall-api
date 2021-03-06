// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mall-api/internal/dao/internal"
)

// internalAdminUserDao is internal type for wrapping internal DAO implements.
type internalAdminUserDao = *internal.AdminUserDao

// adminUserDao is the data access object for table tb_newbee_mall_admin_user.
// You can define custom methods on it to extend its functionality as you wish.
type adminUserDao struct {
	internalAdminUserDao
}

var (
	// AdminUser is globally public accessible object for table tb_newbee_mall_admin_user operations.
	AdminUser = adminUserDao{
		internal.NewAdminUserDao(),
	}
)

// Fill with you ideas below.
