package global

func GetIsp(v int) string {
	switch v {
	case 1:
		return "移动"
	case 2:
		return "电信"
	case 3:
		return "联通"

	}
	return "其他"
}

func GetBelong(v int) string {

	switch v {
	case 2:

		return "招募"
	}
	return "自建"
}
func GetLineType(v int) string {

	switch v {
	case 1:

		return "IDC"
	case 2:
		return "ACDN"
	case 3:
		return "PCDN"
	}
	return ""
}
