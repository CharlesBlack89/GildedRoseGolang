/*

- Christopher Black
- Gilded Rose Kata

*/

package main

import (
	"fmt"
	"strings"
)

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("Ave, mundus! Welcome to the Gilded Rose Kata!")
	GildedRose()
}

func GildedRose() {
	for _, item := range items {
		// Sulfuras doesn't need to be updated
		if item.name != "Sulfuras, Hand of Ragnaros" {
			updateItem(&item)
		}
	}
}

func updateItem(i *Item) {
	updateSellIn(i)

	if i.name == "Backstage passes to a TAFKAL80ETC concert" {
		// update backstage pass
	} else {
		updateQuality(i)
	}

}

func updateQuality(i *Item) {
	entropy := 2
	if i.quality > 1 && i.quality <= 50 {
		if i.sellIn < 1 || isConjuredItem(i) {
			// Aged Brie is the only item whose
			// quality improves with age. Like fine wine.
			if i.name == "Aged Brie" {
				i.quality++
			} else {
				i.quality = i.quality - entropy
			}
		} else {
			// in all other cases, we decrement quality by 1
			if i.name == "Aged Brie" {
				i.quality++
			} else {
				if isConjuredItem(i) {
					i.quality = i.quality - 4
				} else {
					i.quality--
				}
			}
		}
	}
}

func updateSellIn(i *Item) {
	i.sellIn--
}

// A helper function to determine if an item is conjured.
func isConjuredItem(i *Item) bool {
	return strings.Contains(strings.ToLower(i.name), "conjured")
}

// This item has unique properties which make me think a seperate
// helper function to update quality is logical
func updateBackStagePass(i *Item) {

}
