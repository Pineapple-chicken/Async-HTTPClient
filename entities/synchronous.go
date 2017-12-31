package entities

func Naive() [10]traveAgency.Info {

	var ifs [10]traveAgency.Info
	customers := traveAgency.GetCustomerDetails()
	destinations := traveAgency.GetRecommendedDestinations(customers)

	for index, dest := range destinations {
		q := traveAgency.GetQuote(dest)
		w := traveAgency.GetWeatherForecast(dest)
		ifs[index] = traveAgency.Info{Destinations:dest, Quote:q, Weather:w}
	}

	return ifs
}
