package repository

import _interface "main/features/test/model/interface"

func NewMemoryTestRepository() _interface.IMemoryTestRepository {
	return &MemoryTestRepository{}
}
