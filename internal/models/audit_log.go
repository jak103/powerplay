package models

// TODO This should tie to a middleware that auto updates the audit log
type AuditLogEntry struct {
	User   User
	Action string
	Table  string
	ID     int
}
