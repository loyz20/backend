package audit

type UseCase interface {
	LogAction(userID uint, action, tableName string, recordID uint, oldData, newData string) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo}
}

func (u *usecase) LogAction(userID uint, action, tableName string, recordID uint, oldData, newData string) error {
	log := &AuditLog{
		UserID:    userID,
		Action:    action,
		TableName: tableName,
		RecordID:  recordID,
		OldData:   oldData,
		NewData:   newData,
	}
	return u.repo.CreateLog(log)
}
