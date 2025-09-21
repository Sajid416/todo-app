package otp

func (m *Manager) VerifyOTP(email,otp string) (bool,error){
	storedOTP,err:=m.RedisClient.Get(Ctx,email).Result()
	if err!=nil || storedOTP!=otp{
		return false,err
	}
	return true,m.RedisClient.Del(Ctx,email).Err()
}