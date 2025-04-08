## Use Open Router

```golang
package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	openrouter "github.com/shiroyaavish/open_router"
)

type Response struct {} // You reponse struct

func Create(c *fiber.Ctx) error {
	if err := c.BodyParser(&recipeRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	requestBody := map[string]interface{}{} // Your Request Body Details

	var response Response
    secret:= "Your Secret"
	err := openrouter.QuasarAlpha(requestBody, secret, &response)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Response:", response)
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Data Get Successfully",
		"data":    response,
	})

}
```
