package core

// mock storage: using inmemory map

var Storage map[string]interface{}

func GetGlobalStorage() map[string]interface{} {
	if Storage == nil {
		Storage = make(map[string]interface{})
	}
	return Storage
}
