package redis

/********************************************************************************
* Temancode Redis Client Package                                                *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2022-12-28                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import "github.com/redis/go-redis/v9"

type serviceRedis struct {
	Client *redis.Client
}

type ServiceRedis interface {
	SaveAuthorizationUser() error
	GetAuthorizationUser() error
}

func NewServiceRedis() ServiceRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &serviceRedis{Client: rdb}
}

func (h *serviceRedis) SaveAuthorizationUser() error {
	return nil
}

func (h *serviceRedis) GetAuthorizationUser() error {
	return nil
}
