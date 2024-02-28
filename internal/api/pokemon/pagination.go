package pokemon

type PaginationManager struct {
	Next     *string
	Previous *string
}

func NewPaginationManager() *PaginationManager {
	return &PaginationManager{}
}

func (pm *PaginationManager) Update(next, previous *string) {
	pm.Next = next
	pm.Previous = previous
}

func (pm *PaginationManager) GetNext() *string {
	return pm.Next
}

func (pm *PaginationManager) GetPrevious() *string {
	return pm.Previous
}
