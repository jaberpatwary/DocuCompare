package middleware

import (
	"app/src/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func LimiterConfig() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100, // Increased for better UX
		Expiration: 15 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).
				JSON(response.Common{
					Code:    fiber.StatusTooManyRequests,
					Status:  "error",
					Message: "খুব বেশি রিকোয়েস্ট পাঠানো হয়েছে, দয়া করে একটু পর চেষ্টা করুন।",
				})
		},
		SkipSuccessfulRequests: true,
	})
}
