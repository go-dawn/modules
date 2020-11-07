// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	redis "github.com/go-redis/redis/v8"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Cmdable is an autogenerated mock type for the Cmdable type
type Cmdable struct {
	mock.Mock
}

// Del provides a mock function with given fields: ctx, keys
func (_m *Cmdable) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *redis.IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *redis.IntCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.IntCmd)
		}
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, keys
func (_m *Cmdable) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *redis.IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *redis.IntCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.IntCmd)
		}
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key
func (_m *Cmdable) Get(ctx context.Context, key string) *redis.StringCmd {
	ret := _m.Called(ctx, key)

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *redis.StringCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

// MGet provides a mock function with given fields: ctx, keys
func (_m *Cmdable) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *redis.SliceCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *redis.SliceCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.SliceCmd)
		}
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, cursor, match, count
func (_m *Cmdable) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	ret := _m.Called(ctx, cursor, match, count)

	var r0 *redis.ScanCmd
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string, int64) *redis.ScanCmd); ok {
		r0 = rf(ctx, cursor, match, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.ScanCmd)
		}
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, value, expiration
func (_m *Cmdable) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 *redis.StatusCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}

	return r0
}
