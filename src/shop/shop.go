package shop

import "fmt"

func Shop() {
	var categories int
	fmt.Println("Choisissez une catégorie:")
	fmt.Println("1. Épées")
	fmt.Println("2. Armures")
	fmt.Println("3. Coéquipiers")
	fmt.Println("4. Autres")
	fmt.Scan(&categories)
	switch categories {
	case 1:
		fmt.Println("Stone Sword")
		fmt.Println("Iron Sword")
		fmt.Println("Golden Sword")
		fmt.Println("Diamond Sword")

	case 2:
		fmt.Println("Leather armor")
	case 3:
		fmt.Println("Ajouter un nouveau coéquipier")
	case 4:
		fmt.Println("Totem of undyng")
		fmt.Println("Crossbow")
	}
}
