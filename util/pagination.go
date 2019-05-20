package util


func GetLimit(page int) (begin, end int) {
	begin = 0

	if page > 0 {
		begin = (page - 1) * 20
	}

	end = begin + 20
	return begin, end
}
