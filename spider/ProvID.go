package spider

func GetProv(id int) string {
	switch id {
	case 2:
		return "Beijing"
	case 3:
		return "Shanghai"
	case 4:
		return "Tianjing"
	case 5:
		return "Chongqing"
	case 8:
		return "Guangdong"
	case 17:
		return "Jiangsu"
	case 22:
		return "Anhui"
	case 33:
		return "Zhejiang"
	default:
		return "Not found"
	}
}
