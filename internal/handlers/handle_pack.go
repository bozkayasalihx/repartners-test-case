package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	calculatePack "github.com/bozkayasalihx/repartners-test-case/cmd"
)

type JsonBody struct {
	PackSizes    []int `json:"pack_sizes"`
	OrderedPacks []int `json:"ordered_packs"`
}

func prettyPrint(values map[int]int) string {
	var str string
	for k, v := range values {
		str += fmt.Sprintf("%d x %d", k, v)
	}
	return str
}

func HandlePacks(c *fiber.Ctx) error {
	c.Accepts("application/json")
	jsonBody := new(JsonBody)
	c.BodyParser(jsonBody)

	if len(jsonBody.OrderedPacks) == 0 || len(jsonBody.PackSizes) == 0 {
		return c.Send([]byte("please add inputs to body!!"))
	}

	packCalculator := calculatePack.NewPackCalculator(
		jsonBody.PackSizes,
		calculatePack.NewBinarySearcher(jsonBody.PackSizes),
	)
	pretty := map[int]string{}
	for _, order := range jsonBody.OrderedPacks {
		result := packCalculator.CalculatePacks(order)
		pretty[order] = prettyPrint(result)
	}

	return c.JSON(pretty)
}
