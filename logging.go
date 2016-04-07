package main

import (
	"gitlab.com/buzz/user/model"
	"gitlab.fg/go/logger"
)

type userServiceLogginMiddleware struct {
	logger *logger.ServiceLogger
	UserService
}

func (mw userServiceLogginMiddleware) Create(newUser *model.CreateUser) (*model.User, error) {
	user, err := mw.UserService.Create(newUser)
	if err != nil {
		mw.logger.Info("Create", "Service Results", "success", "false", "error", err.Error())
		return user, err
	}
	mw.logger.Info("Create", "Service Results", "success", "true")
	return user, err
}

func (mw userServiceLogginMiddleware) GetAll() ([]model.User, error) {
	users, err := mw.UserService.GetAll()
	if err != nil {
		mw.logger.Info("GetAll", "Service Results", "success", "false", "error", err.Error())
		return users, err
	}
	mw.logger.Info("GetAll", "Service Results", "success", "true")
	return users, err
}

func (mw userServiceLogginMiddleware) GetByID(id string) (*model.User, error) {
	user, err := mw.UserService.GetByID(id)
	if err != nil {
		mw.logger.Info("GetByID", "Service Results", "success", "false", "error", err.Error())
		return user, err
	}
	mw.logger.Info("GetByID", "Service Results", "success", "true")
	return user, err
}

func (mw userServiceLogginMiddleware) GetByUsername(username string) (*model.User, error) {
	user, err := mw.UserService.GetByUsername(username)
	if err != nil {
		mw.logger.Info("GetByUsername", "Service Results", "success", "false", "error", err.Error())
		return user, err
	}
	mw.logger.Info("GetByUsername", "Service Results", "success", "true")
	return user, err
}

func (mw userServiceLogginMiddleware) Login(username, password, referer string) (model.JWTToken, error) {
	token, err := mw.UserService.Login(username, password, referer)
	if err != nil {
		mw.logger.Info("Login", "Service Results", "success", "false", "error", err.Error())
		return token, err
	}
	mw.logger.Info("Login", "Service Results", "success", "true")
	return token, err
}

func (mw userServiceLogginMiddleware) RefreshToken(userID, username, referer string) (model.JWTToken, error) {
	token, err := mw.UserService.RefreshToken(userID, username, referer)
	if err != nil {
		mw.logger.Info("RefreshToken", "Service Results", "success", "false", "error", err.Error())
		return token, err
	}
	mw.logger.Info("RefreshToken", "Service Results", "success", "true")
	return token, err
}

func (mw userServiceLogginMiddleware) Remove(id string) error {
	err := mw.UserService.Remove(id)
	if err != nil {
		mw.logger.Info("Remove", "Service Results", "success", "false", "error", err.Error())
		return err
	}
	return err
}

func (mw userServiceLogginMiddleware) Update(updatedUser *model.UpdateUser) (*model.User, error) {
	user, err := mw.UserService.Update(updatedUser)
	if err != nil {
		mw.logger.Info("Update", "Service Results", "success", "false", "error", err.Error())
		return user, err
	}
	mw.logger.Info("Update", "Service Results", "success", "true")
	return user, err
}
