package notiftype

func (driver Driver) Is(target Driver) bool {
	return driver == target
}
