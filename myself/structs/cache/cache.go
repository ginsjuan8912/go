package cache

import (
	"errors"
	"fmt"
)

var (
	ErrExistingKey     = errors.New("The key exists already in the cache")
	ErrCapacityReached = errors.New("The capacity has been reached")
	ErrInvalidKey      = errors.New("The key is invalid")
	ErrKeyNotExists    = errors.New("The key doesn't exists")
)

type Cache struct {
	capacity int
	elements map[string]interface{}
}

func Create(_capacity int) Cache {
	return Cache{
		capacity: _capacity,
		elements: map[string]interface{}{},
	}
}

// Set *
//*/
func (c Cache) Set(key string, value interface{}) error {

	//The key cannot be empty
	if key == "" {
		return ErrInvalidKey
	}

	//The number of elements cannot reach it's capacity
	if len(c.elements) >= c.capacity {
		return ErrCapacityReached
	}

	//The element cannot exists in an element
	if _, ok := c.elements[key]; ok {
		return ErrExistingKey
	}

	//Perform the save operation
	c.elements[key] = value
	return nil
}

func (c Cache) Get(key string) (interface{}, error) {
	if _, ok := c.elements[key]; ok {
		return c.elements[key], nil
	}

	return nil, ErrKeyNotExists
}

func (c Cache) HandleError(err error, key string) {
	switch err {
	case ErrExistingKey:
		fmt.Printf("The %v already exists in the cache", key)
	case ErrCapacityReached:
		fmt.Printf("The capacity has been reached cannot insert %v", key)
	case ErrInvalidKey:
		fmt.Printf("The key %v cannot be empty", key)
	case ErrKeyNotExists:
		fmt.Printf("The %v key doesn't exists in the cache", key)
	case nil:
		fmt.Printf("The %v key was added sucessfully", key)
	default:
		fmt.Printf("Unkown error")
	}

}

// Dispose /*
func (c Cache) Dispose() {
	c.elements = nil
	c.capacity = 0
}
