package entities

func Optimized() [10]TraveAgency.Info {

    var ifs [10]TraveAgency.Info
    weathers := [10]chan TraveAgency.Weather{}
    quotes := [10]chan TraveAgency.Quoting{}
	customers := TraveAgency.GetCustomerDetails()
	destinations := TraveAgency.GetRecommendedDestinations(customers)

    for i := range weathers {
		weathers[i] = make(chan TraveAgency.Weather)
	}

	for i := range quotes {
		quotes[i] = make(chan TraveAgency.Quoting)
	}

	for i, d := range destinations {
		go func() {
			quotes[i] <- TraveAgency.GetQuote(d)
		}()

		go func() {
			weathers[i] <- TraveAgency.GetWeatherForecast(d)
		}()
	}

	for i, d := range destinations {
		ifs[i] = TraveAgency.Info{Destinations:d, Quote:<-quotes[i], Weather:<-weathers[i]}
	}

	return ifs
}
