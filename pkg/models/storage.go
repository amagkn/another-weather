package models

import "sync"

type WeatherStorage struct {
	mu       sync.RWMutex
	weathers map[string]*Weather
}

func NewWeatherStorage() *WeatherStorage {
	return &WeatherStorage{
		weathers: make(map[string]*Weather),
	}
}

func (s *WeatherStorage) GetWeather(city string) *Weather {
	s.mu.RLock()
	defer s.mu.RUnlock()

	weather, ok := s.weathers[city]
	if !ok {
		return nil
	}

	return weather
}

func (s *WeatherStorage) UpdateWeather(weather *Weather) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.weathers[weather.City] = weather
}
