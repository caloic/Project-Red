package shop

import "fmt"

func Shop() {
	var categories int
	fmt.Println("Choisissez une catégorie:")
	fmt.Println("1. Épées")
	fmt.Println("2. Armures")
	fmt.Println("3. Coéquipiers")
	fmt.Println("4. Autres")
	fmt.Println("5. Crafter un item")
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
		fmt.Println("Totem of undying")
		fmt.Println("Crossbow")
	case 5:
		fmt.Println("Stone Sword")
		fmt.Println("Iron Sword")
		fmt.Println("Golden Sword")
		fmt.Println("Diamond Sword")
		fmt.Println("Iron Armors")
		fmt.Println("Golden Armors")
		fmt.Println("Diamond Armors")
		fmt.Println("Stone Axe")
		fmt.Println("Iron Axe")
		fmt.Println("Golden Axe")
		fmt.Println("Diamond Axe")
		fmt.Println("Bow")
		fmt.Println("Crossbow")
		fmt.Println("Shield")
	}
}
