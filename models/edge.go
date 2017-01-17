package models

import (
	"github.com/astaxie/beego/logs"
)

type EdgeType uint8

type Edge struct {
	Id   uint64
	U    uint64
	V    uint64
	Type EdgeType
}

// enum of edge type
const (
	ACCOUNT2ROLE EdgeType = iota
	ROLE2AUTH
)

func (e *Edge) TableUnique() [][]string {
	return [][]string{
		[]string{"U", "V", "Type"},
	}
}

func SelectRoleV(edge_type EdgeType, edge_u uint64) (v uint64) {
	db := GetOrmer()
	var V uint64
	err := db.Raw("select V from api_edge where type = ? AND U = ?", edge_type, edge_u).QueryRow(&V)
	if err == nil {
		logs.Informational("Edge.V   : ", V)
	}
	return V
}

func SelectRoleU(edge_type EdgeType, edge_u uint64) (u uint64) {
	db := GetOrmer()
	var U uint64
	err := db.Raw("select U from api_edge where type = ? AND V = ?", edge_type, edge_u).QueryRow(&U)
	if err == nil {
		logs.Informational("Edge.U   : ", U)
	}
	return U
}
