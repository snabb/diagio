package diagio

// Count takes anything and calls its Count method and returns the result.
//
// Returns -1 if Count method does not exist.
//
// This is useful for passing around an ordinary [io.Reader] or an ordinary
// [io.Writer] and later calling Count on it.
func Count(i any) (n int64) {
	ci, ok := i.(interface {
		Count() int64
	})
	if !ok {
		return -1
	}
	return ci.Count()
}
