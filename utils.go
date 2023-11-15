package main

// function to check is db is connected, otherwise shut down server
func checkDBConnection() error {
	_, redisErr := db.Ping(ctx).Result()
	if redisErr != nil {
		return redisErr
	}
	_, redisErr = sessionDB.Ping(ctx).Result()
	if redisErr != nil {
		return redisErr
	}
	return nil
}
