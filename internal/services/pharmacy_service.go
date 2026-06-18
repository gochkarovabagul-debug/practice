package services

import (
	"context"

	"github.com/gochkarovabagul-debug/practice/internal/models"
	"github.com/gochkarovabagul-debug/practice/internal/repositories"
)

func PharmacyListService(c context.Context, filter models.PharmacyFilter) (any, error) {
	return repositories.PharmacyList(c, filter)
}
func CreatePharmacyService(c context.Context, name string, address string, hours int, adminuserid int) error {
	return repositories.CreatePharmacy(c, name, address, hours, adminuserid)
}
func DeletePharmacyService(c context.Context, categoryid int) error {
	return repositories.DeletePharmacy(c, categoryid)
}
func UpdatePharmacyService(c context.Context, id int, req models.PharmacyCreateRequest) error {
	return repositories.UpdatePharmacy(c, id, req)
}
func GetPharmacyService(c context.Context, id int) (models.PharmacyResponse, error) {
	return repositories.GetPharmacy(c, id)
}
func FindNearbyPharmaciesService(c context.Context, lon float64, lat float64) (models.NearPharmacies, error) {
	return repositories.FindNearbyPharmacies(c, lon, lat)
}
