package main

import (
    "context"
	"encoding/json"
    "fmt"
    "github.com/go-redis/redis/v8"
    "log"
)

var ctx = context.Background()

// Document represents a basic structure for storing data in Redis
type Document struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Value string `json:"value"`
}

// RedisClient wraps the Redis client and provides CRUD operations
type RedisClient struct {
    client *redis.Client
}

type RedisConfig struct {
    Domain string
    Port int32
    Password string
}

// NewRedisClient initializes a new Redis client
func NewRedisClient(cnf *RedisConfig) *RedisClient {
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", cnf.Domain, cnf.Port),
        Password: cnf.Password, // no password set
        DB:       0,  // use default DB
    })

    return &RedisClient{client: rdb}
}

// CreateDocument saves a new document in Redis
func (r *RedisClient) CreateDocument(doc *Document) error {
    err := r.client.Set(ctx, doc.ID, doc, 0).Err()
    if err != nil {
        return err
    }
    return nil
}

// ReadDocument retrieves a document from Redis by ID
func (r *RedisClient) ReadDocument(id string) (*Document, error) {
    val, err := r.client.Get(ctx, id).Result()
    if err != nil {
        return nil, err
    }

    doc := &Document{}
    err = json.Unmarshal([]byte(val), doc)
    if err != nil {
        return nil, err
    }

    return doc, nil
}

// UpdateDocument updates an existing document in Redis
func (r *RedisClient) UpdateDocument(doc *Document) error {
    err := r.client.Set(ctx, doc.ID, doc, 0).Err()
    if err != nil {
        return err
    }
    return nil
}

// DeleteDocument deletes a document from Redis by ID
func (r *RedisClient) DeleteDocument(id string) error {
    err := r.client.Del(ctx, id).Err()
    if err != nil {
        return err
    }
    return nil
}

func main() {}
