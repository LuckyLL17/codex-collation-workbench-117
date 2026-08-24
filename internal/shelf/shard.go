package shelf

func Shard(digest string) string {
	if len(digest) < 2 {
		return "00"
	}
	return digest[:2]
}
