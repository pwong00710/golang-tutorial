package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"
	"flag"
	"fmt"
	"log"
	"tutorial/golang/test_orm/models"
)

func main() {
	showSQL := flag.Bool("showSQL", false, "a bool")

	flag.Parse()

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	//fmt.Printf("showSQL: %v\n", *showSQL)

	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=testdb sslmode=disable password=Md10112008")
	db.LogMode(*showSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//db.AutoMigrate(&models.User{}, &models.CreditCard{}, &models.Address{}, &models.Email{}, &models.Language{})

	tx := db.Begin()
	/*
		defer func() {
			if tx != nil && tx.Error != nil {
				fmt.Printf("Rollback...\n")
				tx.Rollback()
			}
		}()
	*/

	user := models.User{
		FirstName: "Arthur",
		LastName:  "Dent",
		Username:  "adent",
		Salary:    5000,
		Birthday:  time.Now(),
	}

	var emails []models.Email
	email1 := fmt.Sprintf("%s.%s.%s", "arthur.dent", srand(5, 5, false), "@gmail.com")
	email2 := fmt.Sprintf("%s.%s.%s", "arthur.dent", srand(5, 5, false), "@yahoo.com")

	emails = append(emails, models.Email{Email: email1, Subscribed: true})
	emails = append(emails, models.Email{Email: email2, Subscribed: false})
	//emails = append(emails, models.Email{Email: "arthur.dent@yahoo.com", Subscribed: false})

	user.Emails = emails
	user.CreditCard = models.CreditCard{Number: "017-084-984-333"}
	user.BillingAddress = models.Address{Address1: randString(10), Address2: randString(10), Post: sql.NullString{String: randString(10), Valid: true}}
	user.ShippingAddress = models.Address{Address1: randString(10), Address2: randString(10), Post: sql.NullString{String: randString(10), Valid: true}}

	langs := []models.Language{}
	tx.Find(&langs)
	//fmt.Printf("langs:%v\n", langs)
	user.Languages = []models.Language{langs[randInt(0, len(langs)-1)]}

	if tx.NewRecord(user) {
		//tx.Create(&user)
		err := tx.Create(&user).Error
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	tx.Commit()

	tx = db.Begin()
	fmt.Printf("User:%v\n", user)
	tx.Save(&user)
	tx.Commit()

	tx = db.Begin()

	user = models.User{}
	tx.Table("users").Joins("join emails on emails.user_id = users.id").
		Where("emails.email = ?", email1).
		//Preload("CreditCard").
		//Preload("Emails").
		Find(&user)
	fmt.Printf("User:%v\n", user)
	tx.Model(&user).Update(models.User{Salary: 8000, FirstName: "Peter"})
	tx.Model(&user).Set("gorm:save_associations", false).Update(models.User{Salary: 9000, FirstName: "Mary"})

	//creditCard := models.CreditCard{}
	tx.Model(&user).Related(&user.CreditCard, "CreditCard")
	creditCard := &(user.CreditCard)
	//fmt.Printf("Address: %v/%v\n", &creditCard, &(user.CreditCard))
	fmt.Printf("creditCard.Number: %v\n", (*creditCard).Number)
	tx.Model(creditCard).Update(models.CreditCard{Number: "017-095-450-133"})
	tx.Model(&user).Update(models.User{Salary: 9000, FirstName: "John"})

	user.ID = 0
	user.UserID = 0
	db.First(&user, 1)
	fmt.Printf("User:%v\n", user)
	user.CreditCard.ID = 0
	tx.Model(&user).Related(&user.CreditCard, "CreditCard")
	tx.Model(&user).Update(models.User{Salary: 5000, FirstName: "John"})

	/*
		user = models.User{}
		tx.Table("users").Joins("join emails on emails.user_id = users.id").
			Where("emails.email = ?", email2).
			//Preload("CreditCard").
			//Preload("Emails").
			Find(&user)
		fmt.Printf("User:%v\n", user)
		creditCard = user.CreditCard
		fmt.Printf("creditCard.Number: %v\n", creditCard.Number)
	*/

	tx.Commit()
}

func randInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func randString(len int) string {
	return fmt.Sprintf("%s", srand(len, len, false))
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// generates a random string
func srand(min, max int, readable bool) string {

	var length int
	var char string

	if min < max {
		length = min + rand.Intn(max-min)
	} else {
		length = min
	}

	if readable == false {
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	} else {
		char = "ABCDEFHJLMNQRTUVWXYZabcefghijkmnopqrtuvwxyz23479"
	}

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}
