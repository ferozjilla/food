package food_test

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ferozjilla/home/food"
)

var _ = Describe("Food", func() {
	It("Can create and marshal an ingredient", func() {
		ingredient := Ingredient{
			Name:     "Milk",
			Metadata: "2L, green cap",
			Expiry:   time.Now(),
		}

		_, err := json.Marshal(ingredient)
		Expect(err).NotTo(HaveOccurred())
	})
})
