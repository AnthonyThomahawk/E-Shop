package database

import (
	"github.com/AnthonyThomahawk/E-Shop/server/product"
	"github.com/AnthonyThomahawk/E-Shop/server/user"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) error {
	err := db.First(&product.Category{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err = seedCategories(db)
		if err != nil {
			return errors.Errorf("could not seed categories, error; %v", err)
		}
	}

	err = db.First(&product.Product{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err = seedProducts(db)
		if err != nil {
			return errors.Errorf("could not seed products, error; %v", err)
		}
	}

	err = db.First(&user.Role{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err = seedRoles(db)
		if err != nil {
			return errors.Errorf("could not seed roles, error; %v", err)
		}
	}

	err = db.First(&user.User{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err := seedUsers(db)
		if err != nil {
			return errors.Errorf("could not seed users, error; %v", err)
		}
	}

	return nil
}

func seedCategories(db *gorm.DB) error {
	categories := []product.Category{
		{
			Description: "Discover unique black and black flavored tea varieties from Sri Lanka (Ceylon), China and India.",
			Label:       "Black",
			Color:       "#000000",
		},
		{
			Description: "The best green and green flavored teas from Ceylon, China, Japan and Taiwan.",
			Label:       "Green",
			Color:       "#30bd06",
		},
		{
			Description: "Light and full flavored, low in caffeine, white and white flavored tea is the best choice for anytime of the day.",
			Label:       "White",
			Color:       "#a8b3b2",
		},
		{
			Description: "It balances between black and green tea. Either low or high fermented it is always a tasteful choice.",
			Label:       "Oolong",
			Color:       "#66443e",
		},
		{
			Description: "The South African infusion. Full in Vitamin C, without any caffeine.",
			Label:       "Rooibos",
			Color:       "#943625",
		},
	}

	return db.Create(categories).Error
}

func seedProducts(db *gorm.DB) error {
	var categories []product.Category
	err := db.Find(&categories).Error
	if err != nil {
		return err
	}

	blackID, err := getCategoryID(categories, "Black")
	if err != nil {
		return err
	}

	greenID, err := getCategoryID(categories, "Green")
	if err != nil {
		return err
	}

	whiteID, err := getCategoryID(categories, "White")
	if err != nil {
		return err
	}

	oolongID, err := getCategoryID(categories, "Oolong")
	if err != nil {
		return err
	}

	rooibosID, err := getCategoryID(categories, "Rooibos")
	if err != nil {
		return err
	}

	products := []product.Product{
		// Black tea
		{
			CategoryID:      blackID,
			SKU:             "0000592",
			Label:           "Dimbula",
			Description:     "A flavorful and tasty tea",
			Characteristics: "Dimbula Valley produces flavorful teas in the first 4 months of the year, when the weather is cold and crisp. A highly aromatic tea with salty and rich taste which combines perfect with milk.",
			Price:           2.10,
			Stock:           50000,
			ImageURL:        "https://tea.gr/uploads/gallery/BT01_3.jpg",
		},
		{
			CategoryID:      blackID,
			SKU:             "0000593",
			Label:           "Deniyaya",
			Description:     "Tea with bronze color and intense aroma",
			Characteristics: "From tea plants that grow in the lower slopes of the island, next to the Sinharaja rainforest, this tea with the bold and intense taste is best served with honey.",
			Price:           2.80,
			Stock:           400,
			ImageURL:        "https://tea.gr/uploads/gallery/BT02_3.jpg",
		},
		{
			CategoryID:      blackID,
			SKU:             "0000594",
			Label:           "Ruhunu",
			Description:     "Strong tea with character",
			Characteristics: "The southern part of Sri Lanka has exclusive character in the soil which gives the blackness to this leaf and the strength and color to the cup.",
			Price:           2.10,
			Stock:           2000,
			ImageURL:        "https://tea.gr/uploads/gallery/BT05_2.jpg",
		},
		{
			CategoryID:      blackID,
			SKU:             "0000595",
			Label:           "Loolecondera",
			Description:     "A strong flavorful tea",
			Characteristics: "Carefully selected tea leaves, from Loolecondera estate, Ceylon's first estate, est.1867, offer a rich cup of mellow tea. Great at morning also with a milk drop. Ceylon single-origin Tea.",
			Price:           2.10,
			Stock:           150,
			ImageURL:        "https://tea.gr/uploads/gallery/BT07_3.jpg",
		},
		{
			CategoryID:      blackID,
			SKU:             "0000596",
			Label:           "English Breakfast",
			Description:     "The classic breakfast tea",
			Characteristics: "The rich flavor of quality Dimbula teas have been sharpened with the brightness of selected delicate Nuwara-Eliya teas and skillfully blend into this fine English Breakfast tea. Brew strong and enjoy with milk.",
			Price:           2.10,
			Stock:           300,
			ImageURL:        "https://tea.gr/uploads/gallery/BT09_3.jpg",
		},
		// Green tea
		{
			CategoryID:      greenID,
			SKU:             "0000597",
			Label:           "Young Hyson",
			Description:     "Prepared according to an ancient recipe",
			Characteristics: "Young Hyson variety of green tea manufactured according to Chinese manufacturing methods and is made from high grown teas. This is a pure Ceylon green tea.",
			Price:           2.20,
			Stock:           4000,
			ImageURL:        "https://tea.gr/uploads/gallery/GT01_3.jpg",
		},
		{
			CategoryID:      greenID,
			SKU:             "0000598",
			Label:           "Sencha Mlesna",
			Description:     "Japan's everyday tea",
			Characteristics: "Τhe most common Japanese tea. It has a slightly bitter aftertaste. Ideal for light meals. Also available in a 100gr metal caddie.",
			Price:           2.20,
			Stock:           50,
			ImageURL:        "https://tea.gr/uploads/gallery/GT02_3.jpg",
		},
		{
			CategoryID:      greenID,
			SKU:             "0000599",
			Label:           "Royal Gunpowder",
			Description:     "Strong Ceylon green tea",
			Characteristics: "Τhe leaf is rolled in high pressure in to small silvergreen balls. It has a strong slightly smoky flavor.",
			Price:           2.20,
			Stock:           300,
			ImageURL:        "https://tea.gr/uploads/gallery/GT03_3.jpg",
		},
		{
			CategoryID:      greenID,
			SKU:             "0000600",
			Label:           "Dumbara Valley",
			Description:     "Soft all day green tea",
			Characteristics: "In Dumbara Valley the tea is plucked early in the morning. This results to an all day smooth and mellow tea.",
			Price:           3.40,
			Stock:           600,
			ImageURL:        "https://tea.gr/uploads/gallery/GT04_3.jpg",
		},
		{
			CategoryID:      greenID,
			SKU:             "0000601",
			Label:           "PI LO CHUN - JADE SNAILS",
			Description:     "Flavorful and fruity taste",
			Characteristics: "Also known as «Jade Snails of the Spring». It brews a light green shining cup and its flavor is finely aromatic and fruity, with a light touch of sweetness.",
			Price:           6.10,
			Stock:           1200,
			ImageURL:        "https://tea.gr/uploads/gallery/GT06_3.jpg",
		},
		// White tea
		{
			CategoryID:      whiteID,
			SKU:             "0000602",
			Label:           "Silver Tip",
			Description:     "The silver buds of the tea tree",
			Characteristics: "The silver buds of the plant, carefully plucked offer a taste experience that changes according to the brew time. High quality buds from Sri Lanka, hand plucked from the top of the tea plant.",
			Price:           13.60,
			Stock:           3000,
			ImageURL:        "https://tea.gr/uploads/gallery/WT02_2.jpg",
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000603",
			Label:           "Silver Needle",
			Description:     "Quality white tea buds from the top of the plant",
			Characteristics: "The Silver buds of Camelia Sinensis, hand plucked early in the morning, offer a tasty and light cup. Hand plucked buds from the top of teaplants that grow in China, as soon as they blossom. ",
			Price:           9.00,
			Stock:           550,
			ImageURL:        "https://tea.gr/uploads/gallery/WT03_2.jpg",
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000604",
			Label:           "JASMINE DRAGON TEARS",
			Description:     "Handmade teaballs with jasmine flowers",
			Characteristics: "Hand rolled tea leaves from Sri Lanka with a jasmine flower inside.",
			Price:           12.00,
			Stock:           1300,
			ImageURL:        "https://tea.gr/uploads/gallery/WF05_2.jpg",
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000605",
			Label:           "WHITE ANGEL",
			Description:     "Cake flavored white tea",
			Characteristics: "White and green tea from China, cinnamon, black tea, orange peel, almonds, lemongrass, orange oil & blossoms, rose petals, flavour.",
			Price:           4.00,
			Stock:           600,
			ImageURL:        "https://tea.gr/uploads/gallery/WF02_2.jpg",
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000606",
			Label:           "WHITE PEACH",
			Description:     "Light and unprocessed white tea with white peach flavor",
			Characteristics: "Pai mu tan tea from Sri Lanka with white peach natural oil extracts.",
			Price:           4.00,
			Stock:           950,
			ImageURL:        "https://tea.gr/uploads/gallery/90404.jpg",
		},
		// Oolong tea
		{
			CategoryID:      oolongID,
			SKU:             "0000607",
			Label:           "DIAMOND OOLONG",
			Description:     "Oolong tea combined with flowers",
			Characteristics: "Oolong tea from China, marigold flowers 2%, cornflowers 1,2%, flavor. Ideal for all day long. ",
			Price:           3.00,
			Stock:           3500,
			ImageURL:        "https://tea.gr/uploads/gallery/OL01_2.jpg",
		},
		{
			CategoryID:      oolongID,
			SKU:             "0000608",
			Label:           "Oolong Se Chung",
			Description:     "High fermented oolong with earthy aftertaste",
			Characteristics: "High fermented tea, between green and black. It has an earthy flavor and intense characteristics in the cup.",
			Price:           2.70,
			Stock:           450,
			ImageURL:        "https://tea.gr/uploads/gallery/OL02_2.jpg",
		},
		{
			CategoryID:      oolongID,
			SKU:             "0000609",
			Label:           "Milky Oolong",
			Description:     "Quality milk processed oolong tea",
			Characteristics: "Low fermented Oolong specialy proccessed with milk for a creamy rich taste. A unique flavor.",
			Price:           5.90,
			Stock:           1700,
			ImageURL:        "https://tea.gr/uploads/gallery/OL04_2.jpg",
		},
		{
			CategoryID:      oolongID,
			SKU:             "0000610",
			Label:           "Oolong Mlesna",
			Description:     "A quality oolong with the stamp of Mlesna",
			Characteristics: "A quality tea. It is low fermented. Light with rich aromas and is a specimen of the quality of Mlesna.",
			Price:           3.50,
			Stock:           800,
			ImageURL:        "https://tea.gr/uploads/gallery/OL05_2.jpg",
		},
		{
			CategoryID:      oolongID,
			SKU:             "0000611",
			Label:           "IRON GODDESS - TI KUAN YIN",
			Description:     "Low fermented Oolong",
			Characteristics: "Quality low fermented oolong from Taiwan. Tasty and flavorful with a bright green color.",
			Price:           4.90,
			Stock:           950,
			ImageURL:        "https://tea.gr/uploads/gallery/60686.jpg",
		},
		// Rooibos tea
		{
			CategoryID:      rooibosID,
			SKU:             "0000612",
			Label:           "Organic Rooibos",
			Description:     "The famous South African infusion",
			Characteristics: "Organic red tea from South Africa. Enjoy it with honey.",
			Price:           2.50,
			Stock:           3900,
			ImageURL:        "https://tea.gr/uploads/gallery/RO01_2.jpg",
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000613",
			Label:           "Rooibos chocolate & caramel",
			Description:     "Rooibos mixed with chocolate and caramel",
			Characteristics: "Organic red tea from South Africa, combined with chunks of chocolate and caramel. Caffeine free.",
			Price:           2.70,
			Stock:           7450,
			ImageURL:        "https://tea.gr/uploads/gallery/RO02_2.jpg",
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000614",
			Label:           "Rooibos Chilli & Orange",
			Description:     "Orange peels and chilli peppers",
			Characteristics: "Organic red tea from South Africa, combined with chunks of chilli peppers and orange peels.",
			Price:           2.70,
			Stock:           0,
			ImageURL:        "https://tea.gr/uploads/gallery/RO04_2.jpg",
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000615",
			Label:           "Rooibos Apple & Rose",
			Description:     "Fruits and flowers with red tea",
			Characteristics: "Red tea with apple chunks, rose petals and vanilla. Caffeine free.",
			Price:           2.70,
			Stock:           1300,
			ImageURL:        "https://tea.gr/uploads/gallery/RO05_2.jpg",
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000616",
			Label:           "Rooibos Rum & Bitter almond",
			Description:     "A winter delight",
			Characteristics: "Red tea with bitter almond chunks, raisins and rum aroma. Caffeine free.",
			Price:           2.70,
			Stock:           16000,
			ImageURL:        "https://tea.gr/uploads/gallery/RO06_2.jpg",
		},
	}

	return db.Create(products).Error
}

func seedRoles(db *gorm.DB) error {

	roles := []user.Role{
		{
			Name: "Admin",
		},
		{
			Name: "Customer",
		},
	}

	return db.Create(roles).Error

}

func seedUsers(db *gorm.DB) error {

	var roles []user.Role
	err := db.Find(&roles).Error
	if err != nil {
		return err
	}

	adminID, err := getRoleID(roles, "Admin")
	if err != nil {
		return err
	}

	customerID, err := getRoleID(roles, "Customer")
	if err != nil {
		return err
	}

	password := "12345"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []user.User{
		{
			Email:        "panos@panopoulos.com",
			PasswordHash: string(hash[:]),
			RoleID:       customerID,
		},
		{
			Email:        "antonis@antonopoulos.com",
			PasswordHash: string(hash[:]),
			RoleID:       adminID,
		},
	}

	return db.Create(users).Error
}

func getCategoryID(categories []product.Category, label string) (uint, error) {

	for _, cat := range categories {
		if cat.Label == label {
			return cat.ID, nil
		}
	}
	return 0, errors.Errorf("Category with label %v not in slice", label)
}

func getRoleID(roles []user.Role, name string) (uint, error) {
	for _, role := range roles {
		if role.Name == name {
			return role.ID, nil
		}
	}
	return 0, errors.Errorf("Role with name %v not in slice", name)
}
