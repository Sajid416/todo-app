package otp

import "time"

func (m *Manager) StoreOTP(email, otp string) error {
	return m.RedisClient.Set(Ctx, email, otp, 5*time.Minute).Err()
}
