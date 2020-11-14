package main

import (
	"github.com/aicam/virtualbox-web-service/vbox"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) GetAllVMList() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := vbox.GetAllVMList()
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}

func (s *Server) StopVM() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := vbox.StopVM(context.Param("vm_name"))
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}

func (s *Server) RemoveVM() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := vbox.RemoveVM(context.Param("vm_name"))
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}

func (s *Server) CloneVM() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := vbox.CloneVM(context.Param("vm_name"))
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}

func (s *Server) StartVM() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := vbox.StartVM(context.Param("vm_name"))
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}

func (s *Server) ConfigVM() gin.HandlerFunc {
	return func(context *gin.Context) {
		ram, err := strconv.Atoi(context.Param("ram"))
		if err != nil {
			context.JSON(http.StatusOK, struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "Failed", Message: "Wrong RAM parameter"})
			return
		}
		cpu, err := strconv.Atoi(context.Param("cpu"))
		if err != nil {
			context.JSON(http.StatusOK, struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "Failed", Message: "Wrong CPU parameter"})
			return
		}
		result := vbox.ChangeVMProperties(context.Param("vm_name"),
			ram,
			cpu)
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}

func (s *Server) RunCommandInVM() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req struct {
			VMName  string `json:"vm_name"`
			Command string `json:"command"`
		}
		_ = context.BindJSON(&req)
		result := vbox.RunCommand(req.Command, req.VMName)
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Done.", Message: result})
	}
}
