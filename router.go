package main

func (s *Server) Route() {
	s.Router.GET("/get_vms_info", s.GetAllVMList())
	s.Router.GET("/stop_vm/:vm_name", s.StopVM())
	s.Router.GET("/remove_vm/:vm_name", s.RemoveVM())
	s.Router.GET("/config_vm/:vm_name/:ram/:cpu", s.ConfigVM())
	s.Router.GET("/start_vm/:vm_name", s.StartVM())
	s.Router.POST("/run_command", s.RunCommandInVM())
	s.Router.GET("/clone_vm/:vm_name", s.CloneVM())
}
