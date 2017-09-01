package main

import (
	"fmt"
	"testing"
)

func TestIsConjuredItem(t *testing.T) {
	// we expect true
	if !isConjuredItem(&Item{"Conjured Elven Blade", 10, 10}) {
		t.Error("Function \"isConjuredItem()\" failed...")
	}
	// we expect false
	if isConjuredItem(&Item{"Elven Blade", 10, 10}) {
		t.Error("Function \"isConjuredItem()\" failed...")
	}
	conjuredItem := Item{"Conjured daedric great ax", 30, 10}
	updateQuality(&conjuredItem)
	if conjuredItem.quality != 8 {
		t.Error("Conjured items should degrade twice as fast...")
	}

	// a conjured item with a sellIn of < 1 should degrade by a factor of 4
	reallyOldConjuredItem := Item{"Conjured dwarven curass", -10, 5}
	if updateQuality(&reallyOldConjuredItem); reallyOldConjuredItem.quality != 1 {
		t.Error("Conjured items with a sellIn of < 1 should degrade by a factor of 4")
	}
}

func TestUpdateSellInBasic(t *testing.T) {
	i := Item{"Great Blade", 5, 1}
	if updateSellIn(&i); i.sellIn != 4 {
		t.Error("Function \"updateSellIn\" failed...")
	}
}

func TestUpdateQuality(t *testing.T) {
	lowQualityItem := Item{"Iron dagger", 5, 1}

	// the quality of an item can never be < 1
	if updateQuality(&lowQualityItem); lowQualityItem.quality < 1 {
		t.Error("Function \"updateQuality\" failed...Items cannot have a quality < 1")
	}

	// normally, item quality will decrease by 1
	midQualityItem := Item{"Glass bow", 10, 10}
	if updateQuality(&midQualityItem); midQualityItem.quality != 9 {
		t.Error("Function \"updateQuality\" failed...Items should normally decrement quality by 1")
	}

	// conjured items should decrease in quality twice as fast
	conjuredItem := Item{"Conjured daedric spear", 10, 25}
	if updateQuality(&conjuredItem); conjuredItem.quality != 23 {
		t.Error("Function \"updateQuality\" failed...Conjured items should decrease in quality by 2")
	}

	// lastly, test an item whose sellIn date has passed ~ should also decrease in quality 2x
	dustyOldItem := Item{"Rusty copper dagger", 0, 3}
	if updateQuality(&dustyOldItem); dustyOldItem.quality != 1 {
		t.Error("Function \"updateQuality\" failed...Items whose sellIn date has passed should decrease in quality by 2")
	}
	oldConjuredItem := Item{"Conjured elven bow", -50, 5}
	if updateQuality(&oldConjuredItem); oldConjuredItem.quality != 1 {
		fmt.Println(oldConjuredItem.quality)
		t.Error("Old conjured items should deteriorate twice as fast as normal items")
	}
}

func TestUpdateQualityTooHigh(t *testing.T) {
	i := Item{"Elven glass blade", 10, 60}
	updateQuality(&i)
	if i.quality != 60 {
		t.Error("Update quality is broken!")
	}
}

func TestQualityPastSellIn(t *testing.T) {
	i := Item{"Iron coat", 0, 3}
	updateQuality(&i)
	if i.quality != 1 {
		t.Error("Update quality doesn't work as expected when an item's sellIn is < 1")
	}
}

func TestGetQualityPast50(t *testing.T) {
	i := Item{"Aged Brie", 0, 49}
	updateQuality(&i)
	if i.quality > 50 {
		t.Error("Items' quality cannot surpass 50...")
	}
}

func TestAgedBrie(t *testing.T) {
	i := Item{"Aged Brie", 0, 45}
	updateQuality(&i)
	if i.quality != 46 {
		t.Error("Aged Brie's quality should improve with time...")
	}
	i1 := Item{"Aged Brie", 10, 40}
	updateQuality(&i1)
	if i1.quality != 41 {
		t.Error("Aged Brie's quality should improve with time...")
	}

}

func TestGuildedRoseSainity(t *testing.T) {
	i := Item{"+5 Dexterity Vest", 10, 20}
	updateItem(&i)
	// this is a 'normal' item ~ sellIn && quality should only decrement by 1
	if i.sellIn != 9 && i.quality != 19 {
		t.Error("Missing basic functionality")
	}
}
