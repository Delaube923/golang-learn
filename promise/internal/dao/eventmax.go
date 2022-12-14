// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"promise/internal/dao/internal"
)

// internalEventmaxDao is internal type for wrapping internal DAO implements.
type internalEventmaxDao = *internal.EventmaxDao

// eventmaxDao is the data access object for table eventmax.
// You can define custom methods on it to extend its functionality as you wish.
type eventmaxDao struct {
	internalEventmaxDao
}

var (
	// Eventmax is globally public accessible object for table eventmax operations.
	Eventmax = eventmaxDao{
		internal.NewEventmaxDao(),
	}
)

// Fill with you ideas below.
