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

// NewRedisClient initializes a new Redis client
func NewRedisClient(pass string) *RedisClient {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: pass, // no password set
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

func main() {
    redisClient := NewRedisClient()

    // Create a new document
    doc := &Document{
        ID:    "123",
        Name:  "example",
        Value: "value",
    }

    err := redisClient.CreateDocument(doc)
    if err != nil {
        log.Fatalf("Failed to create document: %v", err)
    }
    fmt.Println("Document created")

    // Read the document
    readDoc, err := redisClient.ReadDocument(doc.ID)
    if err != nil {
        log.Fatalf("Failed to read document: %v", err)
    }
    fmt.Printf("Read document: %+v\n", readDoc)

    // Update the document
    doc.Value = "new value"
    err = redisClient.UpdateDocument(doc)
    if err != nil {
        log.Fatalf("Failed to update document: %v", err)
    }
    fmt.Println("Document updated")

    // Delete the document
    err = redisClient.DeleteDocument(doc.ID)
    if err != nil {
        log.Fatalf("Failed to delete document: %v", err)
    }
    fmt.Println("Document deleted")
}
