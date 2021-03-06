package histories

import (
	"daily-tracker-calories/bussiness/calories"
	"daily-tracker-calories/bussiness/foods"
	"daily-tracker-calories/bussiness/users"
	"log"
	"strconv"
	"time"
)

type serviceHistories struct {
	historiesRepository Repository
	foodsRepository     foods.Repository
	usersService        users.Service
	caloriesService     calories.Service
	foodsService        foods.Service
}

func NewService(repositoryHistories Repository, repositoryFoods foods.Repository, serviceUser users.Service, serviceCalorie calories.Service, serviceFood foods.Service) Service {
	return &serviceHistories{
		historiesRepository: repositoryHistories,
		foodsRepository:     repositoryFoods,
		usersService:        serviceUser,
		caloriesService:     serviceCalorie,
		foodsService:        serviceFood,
	}
}

func (service *serviceHistories) CreateHistories(histories *Domain) (*Domain, error) {
	user, err := service.usersService.FindByID(histories.UserID)
	if err != nil {
		return &Domain{}, err
	}
	histories.UserID = user.ID
	histories.NameUser = user.Name
	food, err := service.foodsService.GetFoodByName(histories.FoodName)
	if err != nil {
		return &Domain{}, err
	}
	histories.FoodID = food.ID
	histories.FoodName = food.Name
	histories.Calorie = food.Calorie
	dateTime := time.Now().Format("2012006")
	histories.Date = dateTime
	log.Println(histories)
	result, err := service.historiesRepository.Insert(histories)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceHistories) GetAllHistoriesByUserID(userid int) (*[]Domain, error) {
	user, err := service.historiesRepository.GetAllHistoriesByUserID(userid)
	if err != nil {
		return &[]Domain{}, err
	}
	return user, nil
}

func (service *serviceHistories) UserStat(userid int) (float64, float64, string, string, error) {
	currentCalorie, err := service.historiesRepository.SumCalorieByUserID(userid)
	if err != nil {
		return 0.0, 0.0, "", "", err
	}
	needCalorie, err := service.caloriesService.GetCalorieFloat(userid)
	if err != nil {
		return 0.0, 0.0, "", "", err
	}
	var status string
	divCalorie := currentCalorie / needCalorie
	result := divCalorie * 100
	convResult := strconv.Itoa(int(result))
	strPercentage := convResult + " %"
	if result < 80 {
		status = "Kurang Makan (<80%)"
	} else if result >= 80 && result <= 100 {
		status = "Cukup Makan (80 - 100%)"
	} else {
		status = "Kelebihan Makan (>100%)"
	}
	return currentCalorie, needCalorie, strPercentage, status, nil
}
