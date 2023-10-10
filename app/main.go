package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	tmp := []byte(" ")
	i := 0

	for {
		if i == 0{
			// Створюємо об'єкт гешу
			hasher := sha3.New256()
			hasher.Write(tmp)
			hashBytes := hasher.Sum(nil)
			hashString := fmt.Sprintf("%x", hashBytes)
			fmt.Printf("%s\n", tmp)

		if hashString == "a03ab19b866fc585b5cb1812a2f63ca861e7e7643ee5d43fd7106b623725fd67" ||
			hashString == "d182aed568b01fee105557a1d173791c798030db267cf94e17102b94dcbbda3c" ||hashString == "7b6a784b05c64d2e669e026fc61296eca2ee8acd5112eb8ae5f16023809e203b" {
			fmt.Println("Знайдено:", hashString)
			fmt.Scanln() // Блокуваня до моменту натиснення кнопки
		}
	}

		tmp[i]++

		// додаємо символ або оновлюємо значення пароля
		if tmp[i] == 127 {
			tmp[i] = 32
			i++

			if i == len(tmp) {
				tmp = append(tmp, 32)
			}
		} else{
			i = 0
		}
	}
}

