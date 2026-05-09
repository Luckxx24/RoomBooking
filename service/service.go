package service

import "github.com/luckxx24/RoomBooking/store"

type Service struct {
	Store store.Storage
}

func Helperrole(r string) bool {
	if r == "admin" || r == "user" {
		return true
	}

	return false
}

func HelperPage(Page, Pagesize int) (int, int) {
	offset := (Page - 1) * Pagesize

	return offset, Pagesize
}
