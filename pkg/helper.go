package pkg

import (
	"database/sql"
	"finance/api/models"
	"math"
)

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}


func PrecentageCalculator(e models.ExpenseCalculator) (models.ExpenseCalculator,error) {
	totalExpense := math.Floor(float64(e.TotalExpense))
    e.Restaurant = math.Floor((float64(e.Restaurant) / totalExpense) * 100)
    e.Supermarkets = math.Floor((float64(e.Supermarkets) / totalExpense) * 100)
    e.BeautyMedecine = math.Floor((float64(e.BeautyMedecine) / totalExpense) * 100)
    e.EnerprainmentSport = math.Floor((float64(e.EnerprainmentSport) / totalExpense) * 100)
	
	return e,nil
}