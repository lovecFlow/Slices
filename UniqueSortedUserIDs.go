//Функция, которая возвращает отсортированный слайс, состоящий из уникальных идентификаторов userIDs. Обработка происходит in-place, то есть без выделения доп. памяти.
import "sort"

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 { //Проверяем на количество элементов в слайсе
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] }) //Сортируем 
	uniqPointer := 0 //Создаём перемнную с нулём 
	for i := 1; i < len(userIDs); i++ { //Ставим цикл на элементы слайса 
		if userIDs[uniqPointer] != userIDs[i] { //Если элементы uniqPointer не равны элементу i, 
			uniqPointer++ //то увеличиваем Pointer на один 
			userIDs[uniqPointer] = userIDs[i] //интерации проходят до момента когда индексы равны
		}
	}

	return userIDs[:uniqPointer+1] //Возвращаем отсортированный слайс 
}
