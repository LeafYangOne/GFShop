// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"GFShop/internal/dao/internal"
)

// internalPraiseInfoDao is internal type for wrapping internal DAO implements.
type internalPraiseInfoDao = *internal.PraiseInfoDao

// praiseInfoDao is the data access object for table praise_info.
// You can define custom methods on it to extend its functionality as you wish.
type praiseInfoDao struct {
	internalPraiseInfoDao
}

var (
	// PraiseInfo is globally public accessible object for table praise_info operations.
	PraiseInfo = praiseInfoDao{
		internal.NewPraiseInfoDao(),
	}
)

// Fill with you ideas below.
