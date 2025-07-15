package mappers

func Int64Ptr(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

func TimestampToMillis(timestamp int64) int64 {
	if timestamp == 0 {
		return 0
	}
	return timestamp
}