package testmod

import (
	"fmt"
)

func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s! Have a nice day.\n", name), nil
	case "fr":
		return fmt.Sprintf("Salut, %s! Passe une bonne journée.\n", name), nil
	case "ru":
		return fmt.Sprintf("Здарова, %s! Пиздатого тебе дня.\n", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s! Que tengas un buen día.\n", name), nil
	case "kz":
		return fmt.Sprintf("Сәлем, %s! Бүгінгі күніңіз жақсы өтсін.\n", name), nil
	default:
		return "", fmt.Errorf("unsupported language: %s", lang)
	}
}
