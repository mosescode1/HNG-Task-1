package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/hng-task-1/types"
	utility "github.com/mosescode1/hng-task-1/utils"
)

func GetFacts(c *fiber.Ctx) error {
    numberStr := c.Query("number")
		// Convert the number query parameter to an integer
		number, err := strconv.Atoi(numberStr)
		var properties []string

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"number": "alphabet",
				"error": true,
			})
		}

		// Check if the number query parameter is empty
		if numberStr == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Number query parameter is required",
			})
		}

		// Check if the number is prime
		isPrime := utility.IsPrime(number)


		// Check if the number is perfect
		isPerfect := utility.IsPerfect(number)
		// Check if the number is Armstrong
		if utility.IsArmstrong(number) {
			properties = append(properties, "armstrong")
		}

		// check if the number is odd or even
		if number%2 == 0 {
			properties = append(properties, "even")
		} else {
			properties = append(properties, "odd")
		}

		// get digit sum
		digit_sum := utility.DigitSum(number)

		// Get FUN FACT
		resp, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d/math?json", number))
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		byt, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil
		}
		var responseData types.Response
 		_= json.Unmarshal(byt,&responseData)


		return c.JSON(fiber.Map{
			"number":     number,
			"properties": properties,
			"is_prime": isPrime,
			"is_perfect": isPerfect,
			"fun_fact":     responseData.Text,
			"digit_sum":  digit_sum,
		})
}