// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"GFShop/internal/dao/internal"
)

// internalFileInfoDao is internal type for wrapping internal DAO implements.
type internalFileInfoDao = *internal.FileInfoDao

// fileInfoDao is the data access object for table file_info.
// You can define custom methods on it to extend its functionality as you wish.
type fileInfoDao struct {
	internalFileInfoDao
}

var (
	// FileInfo is globally public accessible object for table file_info operations.
	FileInfo = fileInfoDao{
		internal.NewFileInfoDao(),
	}
)

// Fill with you ideas below.
