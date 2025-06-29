// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"GFShop/internal/dao/internal"
)

// internalPermissionInfoDao is internal type for wrapping internal DAO implements.
type internalPermissionInfoDao = *internal.PermissionInfoDao

// permissionInfoDao is the data access object for table permission_info.
// You can define custom methods on it to extend its functionality as you wish.
type permissionInfoDao struct {
	internalPermissionInfoDao
}

var (
	// PermissionInfo is globally public accessible object for table permission_info operations.
	PermissionInfo = permissionInfoDao{
		internal.NewPermissionInfoDao(),
	}
)

// Fill with you ideas below.
