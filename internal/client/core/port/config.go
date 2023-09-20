package port

type Config interface {
	GetGroup(group string) (map[string]interface{}, error)
	GetField(group string, field string) (string, error) 
}