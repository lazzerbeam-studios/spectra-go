package cache

import "context"

func SetKey(key string, value string, seconds int64) error {
	ctx := context.Background()
	err := ValkeyClient.Do(ctx, ValkeyClient.B().Set().Key(key).Value(value).ExSeconds(seconds).Build()).Error()
	if err != nil {
		return err
	}
	return nil
}

func GetKey(key string) (string, error) {
	ctx := context.Background()
	value, err := ValkeyClient.Do(ctx, ValkeyClient.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return "", err
	}
	return value, nil
}

func DelKey(key string) error {
	ctx := context.Background()
	err := ValkeyClient.Do(ctx, ValkeyClient.B().Del().Key(key).Build()).Error()
	if err != nil {
		return err
	}
	return nil
}
