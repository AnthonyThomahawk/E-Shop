package database

import (
	prd "github.com/AnthonyThomahawk/E-Shop/server/product"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) error {
	err := db.First(&prd.Category{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err = seedCategories(db)
		if err != nil {
			return errors.Errorf("could not seed categories, error; %v", err)
		}
	}

	err = db.First(&prd.Product{}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err := seedProducts(db)
		if err != nil {
			return errors.Errorf("could not seed products, error; %v", err)
		}
	}

	return nil
}

func seedCategories(db *gorm.DB) error {
	categories := []prd.Category{
		{
			Description: "Ανακαλύψτε μοναδικές ποικιλίες μαύρου & μαύρου αρωματικού τσαγιού από την Κεϋλάνη, την Κίνα και την Ινδία.",
			Label:       "Μαύρο",
		},
		{
			Description: "Οι καλύτερες ποικιλίες πράσινου και πράσινου αρωματικού τσαγιού από την Κεϋλάνη, την Κίνα, την Ιαπωνία και την Ταϊβάν.",
			Label:       "Πράσινο",
		},
		{
			Description: "Ελαφρύ και γεμάτο γεύση, σκέτο ή αρωματικό, το λευκό τσάι αποτελεί μια ιδανική επιλογή για όλες τις ώρες της ημέρας.",
			Label:       "Λευκό",
		},
		{
			Description: "Ισορροπεί μεταξύ πράσινου και μαύρου τσαγιού. Είτε χαμηλής, είτε υψηλής ζύμωσης αποτελεί μια μοναδική γευστική επιλογή.",
			Label:       "Oolong",
		},
		{
			Description: "Το ρόφημα από την Νότια Αφρική, πλούσιο σε βιταμίνη C, με ξεχωριστή γεύση και χωρίς καφεΐνη. ",
			Label:       "Rooibos",
		},
	}

	return db.Create(categories).Error
}

func seedProducts(db *gorm.DB) error {
	var categories []prd.Category
	err := db.Find(&categories).Error
	if err != nil {
		return err
	}

	blackID, err := getCategoryID(categories, "Μαύρο")
	if err != nil {
		return err
	}

	greenID, err := getCategoryID(categories, "Πράσινο")
	if err != nil {
		return err
	}

	whiteID, err := getCategoryID(categories, "Λευκό")
	if err != nil {
		return err
	}

	rooibosID, err := getCategoryID(categories, "Rooibos")
	if err != nil {
		return err
	}

	products := []prd.Product{
		// Black tea
		{
			CategoryID:      blackID,
			SKU:             "0000592",
			Label:           "Dimbula",
			Description:     "Γευστικό και αρωματικό τσάι",
			Characteristics: "Η κοιλάδα Dimbula παράγει γευστικά τσάγια τους 4 πρώτους μήνες του χρόνου όταν ο καιρός είναι κρύος και ξηρός. Ένα πλούσιο σε άρωμα τσάι με κεχριμπαρένιο χρώμα που συνδυάζεται όμορφα με γάλα. Είναι μοναδικής προέλευσης από την Σρι Λανκα (περιοχή Dimbula).",
			Price:           2.10,
			Stock:           50000,
		},
		{
			CategoryID:      blackID,
			SKU:             "0000593",
			Label:           "Deniyaya",
			Description:     "Τσάι με χάλκινο χρώμα και έντονη γεύση",
			Characteristics: "Από φυτά μεγαλωμένα στις χαμηλές δυτικές πλαγιές του νησιού, που συνορεύει με το τροπικό δάσος Sinharaja. Τσάι με χάλκινο χρώμα, τολμηρή και έντονη γεύση. Συνδυάζεται ιδανικά με μέλι. Ένα τσάι μοναδικής προέλευσης από την Σρί Λάνκα (περιοχή Deniyaya).",
			Price:           2.80,
			Stock:           400,
		},
		{
			CategoryID:      blackID,
			SKU:             "0000594",
			Label:           "Ruhuhu",
			Description:     "Τσάι με ένταση και χαρακτήρα στην γεύση",
			Characteristics: "H μοναδικότητα του χώματος στη νότια Σρι Λάνκα προσφέρει ένα τσάι με σκούρο φύλλο, ένταση και χρώμα στην κούπα.",
			Price:           2.10,
			Stock:           2000,
		},
		{
			CategoryID:      blackID,
			SKU:             "0000595",
			Label:           "Loolecondera",
			Description:     "Μαλακό τσάι με γεμάτη γεύση",
			Characteristics: "Επιλεγμένα φύλλα τσαγιού, από την φυτεία Loolecondera, την 1η της Κευλάνης από το 1867. Ένα τσάι με σκούρο και γεμάτο σώμα, γεύση βύνης και φρούτων και γλυκό άρωμα. Προτείνεται για πρωινό τσάι με μία σταγόνα γάλα. Μοναδικής προέλευσης από την Σρί Λάνκα!",
			Price:           2.10,
			Stock:           150,
		},
		{
			CategoryID:      blackID,
			SKU:             "0000596",
			Label:           "English Breakfast",
			Description:     "Το κλασικό μαύρο τσάι του πρωϊνού",
			Characteristics: "Η πλούσια γεύση του τσαγιού από την περιοχή Dimbula συνδυασμένη, από τους γευσιγνώστες της Mlesna, με την φωτεινότητα του τσαγιού από την Nuwara Eliya, προσφέρουν το ιδανικό English Breakfast. Ανώτερης ποιότητας έντονο τσάι για το πρώι που μπορεί να συνδυαστεί με γάλα.",
			Price:           2.10,
			Stock:           300,
		},
		// Green tea
		{
			CategoryID:      greenID,
			SKU:             "0000597",
			Label:           "Young Hyson",
			Description:     "Τσάι παρασκευασμένο με αρχαία Κινέζικη συνταγή",
			Characteristics: "Η ποικιλία πράσινου τσαγιού Young Hyson παρασκευάζεται σύμφωνα με αρχαίες Κινέζικες μεθόδους από φυτά που μεγαλώνουν σε μεγάλα υψόμετρα. Είναι 100% τσάι Κεϋλάνης. Διαθέσιμο και σε μεταλλικό κουτί 100γρ.",
			Price:           2.20,
			Stock:           4000,
		},
		{
			CategoryID:      greenID,
			SKU:             "0000598",
			Label:           "Sencha Mlesna",
			Description:     "Το καθημερινό τσάι της Ιαπωνίας",
			Characteristics: "Το πιο γνωστό γιαπωνέζικο τσάι με την ελαφρώς στυφή γεύση. Ιδανικό για να συνοδεύει ελαφριά γεύματα. Διαθέσιμο και σε μεταλλικό κουτί 100γρ.",
			Price:           2.20,
			Stock:           50,
		},
		{
			CategoryID:      greenID,
			SKU:             "0000599",
			Label:           "Royal Gunpowder",
			Description:     "Δυνατό πράσινο τσάι Κεϋλάνης",
			Characteristics: "Το φύλλο στρίβεται με υψηλή πίεση σε μικρά ασημοπράσινα μπαλάκια. Έχει δυνατή, ελαφρώς καπνιστή γεύση. Διαθέσιμο και σε μεταλλικό κουτί 100γρ.",
			Price:           2.20,
			Stock:           300,
		},
		{
			CategoryID:      greenID,
			SKU:             "0000600",
			Label:           "Κοιλάδα Dumbara",
			Description:     "Μαλακό τσάι για όλη την ημέρα",
			Characteristics: "Τα φύλλα του τσαγιού μαζεύονται νωρίς το πρωί στην κοιλάδα Dumbara προσφέροντας ως αποτέλεσμα ένα μαλακό και γλυκό τσάι, ιδανικό για όλη την ημέρα. Διαθέσιμο και σε μεταλλικό κουτί 100γρ.",
			Price:           3.40,
			Stock:           600,
		},
		{
			CategoryID:      greenID,
			SKU:             "0000601",
			Label:           "PI LO CHUN - Σαλιγκάρια Νεφρίτη",
			Description:     "Αρωματική και φρουτώδης γεύση",
			Characteristics: "Είναι γνωστό και ως «Σαλιγκάρια από Νεφρίτη». Με απαλό, φωτεινό πράσινο χρώμα στην κούπα και γεύση που βγάζει αρώματα, φρούτα και μία ελαφριά γλυκιά επίγευση. Διαθέσιμο και σε μεταλλικό κουτί 50γρ.",
			Price:           6.10,
			Stock:           1200,
		},
		// White tea
		{
			CategoryID:      whiteID,
			SKU:             "0000602",
			Label:           "Silver Tip",
			Description:     "Οι κλειστές κορυφές του φυτού μαζεμένες μία μία",
			Characteristics: "Οι ασημένιες κλειστές κορυφές του φυτού προσφέρουν μία εμπειρία που αλλάζει ανάλογα με τον χρόνο παραμονής στο νερό. Υψηλής ποιότητας μπουμπούκια φύλλων, από φυτείες τσαγιού στην Σρι Λάνκα, μαζεύονται με το χέρι μόλις βγούν.",
			Price:           13.60,
			Stock:           3000,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000603",
			Label:           "Silver Needle",
			Description:     "Ποιοτικό λευκό τσάι από την κορυφή του φυτού",
			Characteristics: "Οι ασημένιες βελονίτσες του Camelia Sinensis, μαζεμένες νωρίς το πρωί, προσφέρουν μία γευστική κούπα τσαγιού. Μπουμπούκια φύλλών τσαγιού από φυτά που μεγαλώνουν στην Κίνα, μαζεμένα ένα-ένα με το χέρι.",
			Price:           9.00,
			Stock:           550,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000604",
			Label:           "Δάκρυα του Δράκου",
			Description:     "Χειροποίητες μπαλίτσες λευκού τσαγιού με άνθη γιασεμιού",
			Characteristics: "Χειροποίητα στριμμένα φυλλαράκια τσαγιού από την Σρι Λάνκα με ένα λουλούδι γιασεμιού στο εσωτερικό. Παράγεται κάθε Άνοιξη όταν ανθίσουν τα γιασεμιά. Απλά μοναδικό.",
			Price:           12.00,
			Stock:           1300,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000605",
			Label:           "Λευκό Λεμόνι",
			Description:     "Ελαφρύ και ακατέργαστο λευκό τσάι με φυσικό αιθέριο έλαιο λεμονιού",
			Characteristics: "Ακατέργαστο λευκό τσάι από την Σρι Λάνκα με φυσικό αιθέριο έλαιο λεμονιού. Με υπόξινη επίγευση.",
			Price:           4.00,
			Stock:           600,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000606",
			Label:           "Λευκό Ροδάκινο",
			Description:     "Ελαφρύ και ακατέργαστο λευκό τσάι με φυσικό αιθέριο έλαιο λευκού ροδάκινου",
			Characteristics: "Ακατέργαστο λευκό τσάι από την Σρι Λάνκα με φυσικό αιθέριο έλαιο λευκού ροδάκινου.. Με γλυκιά επίγευση.",
			Price:           4.00,
			Stock:           950,
		},
		// Oolong tea
		{
			CategoryID:      whiteID,
			SKU:             "0000607",
			Label:           "Diamond Oolong",
			Description:     "Τσάι Oolong συνδυασμένο με λουλούδια",
			Characteristics: "Τσάι oolong συνδυασμένο με λουλούδια καλέντουλας και καλαμποκιού. Ιδανικό για όλη την ημέρα.",
			Price:           3.00,
			Stock:           3500,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000608",
			Label:           "Oolong Se Chung",
			Description:     "Τσάι υψηλής ζύμωσης με γήινη γεύση",
			Characteristics: "Τσάι υψηλής ζύμωσης, που αιωρείται μεταξύ πράσινου και μαύρου. Έχει γήινη γεύση και έντονα χαρακτηριστικά.",
			Price:           2.70,
			Stock:           450,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000609",
			Label:           "Milky Oolong",
			Description:     "Ποιοτικό τσάι επεξεργασμένο με γάλα",
			Characteristics: "Μία ξεχωριστή γεύση. Τσάι oolong ειδικά επεξεργασμένο με γάλα για να αποκτήσει κρεμώδη και πλούσια γεύση. Απλά μοναδικό. (περιέχει λακτόζη).",
			Price:           5.90,
			Stock:           1700,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000610",
			Label:           "Oolong Mlesna",
			Description:     "Ένα ποιοτικό τσάι με την σφραγίδα της Mlesna",
			Characteristics: "Ποιοτικό τσάι χαμηλής ζύμωσης, ελαφρύ. Διαθέτει πλούσια αρώματα και αποτελεί δείγμα της ποιότητας των τσαγιών της Mlesna.",
			Price:           3.50,
			Stock:           800,
		},
		{
			CategoryID:      whiteID,
			SKU:             "0000611",
			Label:           "Σιδερένια Θέα - Ti Kuan Yin",
			Description:     "Oolong χαμηλής ζύμωσης",
			Characteristics: "Ποιοτικό χαμηλής ζύμωσης oolong της Ταϊβάν. Γευστικό και έντονα αρωματικό με ιδιαίτερο πράσινο χρώμα.",
			Price:           4.90,
			Stock:           950,
		},
		// Rooibos tea
		{
			CategoryID:      rooibosID,
			SKU:             "0000612",
			Label:           "Organic Rooibos",
			Description:     "Το γνωστό ρόφημα της Νότιας Αφρικής",
			Characteristics: "Βιολογικό κόκκινο τσάι από την Ν. Αφρική. Πίνεται ευχάριστα με μέλι και γάλα. Είναι πλούσιο σε Βιταμίνη C και χωρίς καφεΐνη.",
			Price:           2.50,
			Stock:           3900,
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000613",
			Label:           "Rooibos Σοκολάτα & Καραμέλα",
			Description:     "Rooibos συνδυασμένο με σοκολάτα και καραμέλα",
			Characteristics: "Rooibos από την Νότια Αφρική από την Νότια Αφρική με κομμάτια σοκολάτας (ζάχαρη, κακαόμαζα, βούτυρο κακάο, σκόνη κακάο, λεκιθίνη Ε322, βανίλια), μήλο και flakes καρύδας. Χωρίς καφεΐνη.",
			Price:           2.70,
			Stock:           7450,
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000614",
			Label:           "Rooibos Πορτοκάλι & Τσίλι",
			Description:     "Φλούδες πορτοκαλιού και πιπεριές τσίλι",
			Characteristics: "Rooibos Νότιας Αφρικής, λεμονόχορτο, τσιπς μπανάνας (μπανάνα, λάδι καρύδας, ζάχαρη, άρωμα) τσίλι, φλούδες και λουλούδια πορτοκαλιού, άρωμα. Χωρίς καφεΐνη.",
			Price:           2.70,
			Stock:           0,
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000615",
			Label:           "Rooibos Μήλο & Τριαντάφυλλο",
			Description:     "Φρούτα και λουλούδια με κόκκινο τσάι",
			Characteristics: "Κόκκινο τσάι με κομμάτια μήλου, πέταλα τριαντάφυλλου και βανίλια. Χωρίς καφεΐνη.",
			Price:           2.70,
			Stock:           1300,
		},
		{
			CategoryID:      rooibosID,
			SKU:             "0000616",
			Label:           "Rooibos Ρούμι & Πικραμύγδαλο",
			Description:     "Ένας χειμερινός γευστικός συνδυασμός",
			Characteristics: "Κόκκινο τσάι με κομμάτια πικραμύγδαλου, σταφίδες και άρωμα ρούμι. Χωρίς καφεΐνη.",
			Price:           2.70,
			Stock:           16000,
		},
	}

	return db.Create(products).Error
}

func getCategoryID(categories []prd.Category, label string) (uint, error) {
	for _, cat := range categories {
		if cat.Label == label {
			return cat.ID, nil
		}
	}
	return 0, errors.Errorf("Category with label %v not in slice", label)
}
