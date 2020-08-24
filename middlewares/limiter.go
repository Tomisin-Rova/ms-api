package middlewares

//var limiterCache = cache.NewCache(secrets.GetSecrets().RedisAddress, secrets.GetSecrets().RedisPassword, 10, 1800, "RateLimiter_") // 30 minutes.

/**
 * Limits the request to rate per duration.
 * Using this article https://blog.cloudflare.com/counting-things-a-lot-of-different-things/
 * And found an already existing lib: github.com/shareed2k/fiber_limiter
 *
 * param: int           rate
 * param: time.Duration duration
 * return: func(*fiber.Ctx)
 */
//func RateLimiter(rate int, duration time.Duration) func(*fiber.Ctx) {
//	client := limiterCache.Client().(*redis.Client)
//	cfg := limiter.Config{
//		Rediser:   client,
//		Max:       rate,
//		Burst:     rate,
//		Period:    duration,
//		Prefix:    "helpnow_limit",
//		Algorithm: limiter.SlidingWindowAlgorithm,
//	}
//
//	return limiter.New(cfg)
//}
