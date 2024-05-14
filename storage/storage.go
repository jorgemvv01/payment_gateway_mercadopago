package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"payment_gateway_mercadopago/domain/models/business_model"
	"payment_gateway_mercadopago/domain/models/log_model"
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
	"payment_gateway_mercadopago/domain/models/order_model"
	"payment_gateway_mercadopago/domain/models/product_model"
	"payment_gateway_mercadopago/domain/models/user_model"
)

var DB *gorm.DB
var isOpenTestDB bool

func InitializeDB() {
	if DB == nil {
		var err error
		dsn := os.Getenv("DATABASE_URL")
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
}

func MigrateModels() {
	if err := DB.AutoMigrate(
		&user_model.User{},
		&business_model.Business{},
		&product_model.Product{},
		&order_model.Order{},
		&order_model.OrderDetail{},
		&mp_payment_model.MPPayment{},
		&log_model.Log{},
	); err != nil {
		panic("failed to migrate models")
	}
}

func SeedData() {
	userData := &user_model.User{
		Email:    "johndoe@gmail.com",
		Name:     "John",
		LastName: "Doe",
	}
	DB.Find(&userData)
	if userData.ID == 1 {
		return
	}

	if err := DB.Create(
		userData,
	).Error; err != nil {
		panic("failed seeding User data")
	}

	businessData := []business_model.Business{
		{
			Name:        "Nexus Store",
			Description: "The BEST shoe Store",
			MpToken:     os.Getenv("MP_TOKEN"),
		},
		{
			Name:        "TechTrend Haven",
			Description: "Explore the latest in technology and innovative gadgets. From smart devices to cutting-edge equipment, we're at the forefront of innovation.",
			MpToken:     os.Getenv("MP_TOKEN"),
		},
		{
			Name:        "Vintage Charm Bazaar",
			Description: "Immerse yourself in the charm of the past with our unique collection of vintage items. From retro fashion to classic decor, discover treasures from yesteryear's.",
			MpToken:     os.Getenv("MP_TOKEN"),
		},
	}
	if err := DB.CreateInBatches(businessData, len(businessData)).Error; err != nil {
		panic("failed seeding Business data")
	}

	products := []product_model.Product{
		{
			BusinessID:  1,
			Name:        "Air Jordan 1 Retro High OG",
			Description: "The Air Jordan 1 Retro High remakes the classic sneaker, giving you a fresh look with a familiar feel. Premium materials with new colors and textures give modern expression to an all-time favorite.",
			Code:        "DZ5485-160",
			Price:       703778,
			Discount:    12,
			Tax:         19,
			Image:       "",
		},
		{
			BusinessID:  1,
			Name:        "Nike Dunk Low Retro",
			Description: "Created for the hardwood but taken to the streets, the basketball icon returns with classic details and throwback hoops flair. Channeling '80s vibes, its padded, low-cut collar lets you take your game anywhere—in comfort.",
			Code:        "DV0831-108",
			Price:       449636,
			Discount:    0,
			Tax:         19,
			Image:       "",
		},
		{
			BusinessID:  1,
			Name:        "Nike P-6000",
			Description: "The Nike P-6000 draws on the 2006 Nike Air Pegasus, bringing you a mash-up of iconic style that's breathable, comfortable and evocative of that early-2000s vibe.",
			Code:        "CD6404-201",
			Price:       427845,
			Discount:    25,
			Tax:         19,
			Image:       "",
		},
		{
			BusinessID:  1,
			Name:        "Nike Zegama 2",
			Description: "Up the mountain, through the woods, to the top of the trail you can go. Equipped with an ultra-responsive ZoomX foam midsole, the Zegama 2 is designed to conquer steep ridges, jagged rocks and races from trailhead to tip. Optimal cushioning complements a rugged outsole made for your trail running journey.",
			Code:        "CD6404-201",
			Price:       700110,
			Discount:    11,
			Tax:         19,
			Image:       "",
		},
		{
			BusinessID:  1,
			Name:        "Nike Air Max 270",
			Description: "Nike's first lifestyle Air Max brings you style, comfort and big attitude in the Nike Air Max 270. The design draws inspiration from Air Max icons, showcasing Nike's greatest innovation with its large window and fresh array of colors.",
			Code:        "AH8050-100",
			Price:       622320,
			Discount:    32,
			Tax:         19,
			Image:       "",
		},
	}
	if err := DB.CreateInBatches(products, len(products)).Error; err != nil {
		panic("failed seeding Products data")
	}
}

func InitializeTestDB() {
	if DB == nil && !isOpenTestDB {
		isOpenTestDB = true
		var err error
		dsn := os.Getenv("DATABASE_TEST_URL")
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		MigrateModels()
		SeedData()
		return
	}
}
