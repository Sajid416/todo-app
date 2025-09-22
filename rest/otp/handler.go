package otp

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{
	Manager *Manager
}

func NewHandler(m *Manager) *Handler{
	return &Handler{Manager:m}
}

func SendEmail(email,otp string) error{
	log.Printf("Sending OTP to %s:%s\n",email,otp)
	return nil
}
func (h *Handler) SendOTP(w http.ResponseWriter, r *http.Request){
	email:=r.URL.Query().Get("email")
	if email==""{
		http.Error(w,"Email is required",http.StatusBadRequest)
		return
	}
	otp:=h.Manager.GenerateOTP()
	if err:=h.Manager.StoreOTP(email,otp);err!=nil{
		http.Error(w,"Failed to send email",http.StatusInternalServerError)
	    return
	}
	if err:=SendEmail(email,otp);err!=nil{
		http.Error(w, "Failed to send email",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP sent successfully"))
} 

func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request){
	email:=r.URL.Query().Get("email")
	inputOTP:=r.URL.Query().Get("otp")
	if email=="" || inputOTP==""{
		http.Error(w,"Email and OTP are required",http.StatusBadRequest)
		return
	}
	valid,err:=h.Manager.VerifyOTP(email,inputOTP)
	if err!=nil || !valid {
		http.Error(w,"Invalid OTP",http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status":"verified"})
}