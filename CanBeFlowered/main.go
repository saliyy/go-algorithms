package main

import "fmt"

func main() {
	//
	flowered := []int{0}
	fmt.Println(canPlaceFlowers(flowered, 1))
}

var (
	isNotPlanted = 0
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	placesToPlantCount := 0

	if len(flowerbed) == 1 && flowerbed[0] == isNotPlanted && n == 1 {
		return true
	}

	for i := 0; i < len(flowerbed)-1; i += 1 {

		prevIndex := i - 1
		nextIndex := i + 1

		if flowerbed[i] == isNotPlanted {

			if prevIndex == -1 {
				if flowerbed[nextIndex] == isNotPlanted {
					flowerbed[i] = 1
					placesToPlantCount++
				}
				continue
			}

			if nextIndex == len(flowerbed)-1 && flowerbed[nextIndex] == isNotPlanted {
				if flowerbed[i] == isNotPlanted {
					flowerbed[i] = 1
					placesToPlantCount++
				}
				continue
			}

			if flowerbed[prevIndex] == isNotPlanted && flowerbed[nextIndex] == isNotPlanted {
				flowerbed[i] = 1
				placesToPlantCount++
			}
		}
	}

	fmt.Println(placesToPlantCount)
	return placesToPlantCount >= n
}
