package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/telemetry/config"
	"gorm.io/gorm"
)
type Book struct {
	Author string `json:"author"` 
	Title  string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) CreateBook(context *fiber.Ctx) error {
	Book := Book{}

	err := context.BodyParser(&Book)

	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "could not parse body"
		})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		 context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create book",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book created successfully",
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(&bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get books",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (r *Repository) setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_books/:id", r.GetBookByID)
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/books", r.GetBooks)
}


func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.setupRoutes(app)
	app.Listen(":8080")


}