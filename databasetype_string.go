// Code generated by "stringer -type DatabaseType database.go"; DO NOT EDIT.

package dog

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PostGres-0]
	_ = x[Redis-1]
	_ = x[MySQL-2]
}

const _DatabaseType_name = "PostGresRedisMySQL"

var _DatabaseType_index = [...]uint8{0, 8, 13, 18}

func (i DatabaseType) String() string {
	if i < 0 || i >= DatabaseType(len(_DatabaseType_index)-1) {
		return "DatabaseType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DatabaseType_name[_DatabaseType_index[i]:_DatabaseType_index[i+1]]
}
