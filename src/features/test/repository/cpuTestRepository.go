package repository

import (
	_interface "main/features/test/model/interface"
)

func NewCpuTestRepository() _interface.ICpuTestRepository {
	return &CpuTestRepository{}
}
