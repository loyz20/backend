package audit

import "time"

type AuditLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`    // siapa yang melakukan perubahan
	Action    string    `json:"action"`     // create, update, delete
	TableName string    `json:"table_name"` // nama tabel/modul
	RecordID  uint      `json:"record_id"`  // id record yang diubah
	OldData   string    `json:"old_data"`   // json string sebelum update/delete
	NewData   string    `json:"new_data"`   // json string setelah create/update
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
