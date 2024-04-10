package utils

type Cache struct {
	maps map[string]any
}

func (s *Cache) GetKey(value any) string {
	if s.maps == nil {
		return ""
	}
	for key, tmpConn := range s.maps {
		if tmpConn == value {
			return key
		}
	}
	return ""
}
func (s *Cache) Get(key string) any {
	if s.maps == nil {
		return nil
	}
	return s.maps[key]
}
func (s *Cache) Add(key string, value any) {
	if s.maps == nil {
		s.maps = make(map[string]any)
	}
	s.maps[key] = value
}

func (s *Cache) Remove(value any) {
	if s.maps == nil {
		s.maps = make(map[string]any)
	}
	clientId := s.GetKey(value)

	delete(s.maps, clientId)
}

// 检查是否授权了
func (s *Cache) CheckExist(value any) bool {
	if s.maps == nil {
		return false
	}
	key := s.GetKey(value)
	return key != ""
}

func (s *Cache) CheckExistByKey(key string) bool {
	if s.maps == nil {
		return false
	}
	return s.maps[key] != nil
}

func (s *Cache) GetAll() map[string]any {
	return s.maps
}
