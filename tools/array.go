package tools

import "gateway/models"

func Remove(list []*models.Essen, item *models.Essen) []*models.Essen {
	for i, v := range list {
		if v == item {
			copy(list[i:], list[i+1:])
			list[len(list)-1] = nil
			list = list[:len(list)-1]
		}
	}
	return list
}
