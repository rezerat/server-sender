package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool

func Сonnect() error {
    err := godotenv.Load(".env")
    if err != nil {
		log.Println(".env doesn't find")
	}
    pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        return fmt.Errorf("Pool creation error: %w\n", err)
    }
   if err = pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("Database is not reachable: %w\n", err)
	}

	return nil
}

func SaveDevice(clientId string, nickname string, pubKey []byte) error{
    query := `INSERT INTO devices (client_id, nickname, public_key) VALUES ($1, $2, $3)`
    println(nickname)
    _, err := pool.Exec(context.Background(), query, clientId, nickname, pubKey)
    return err

}

func GetPublicKey(clientId string) ([]byte, error) {
    pubKey := []byte{}
    query := `SELECT public_key FROM devices WHERE client_id = $1`
    err := pool.QueryRow(context.Background(), query, clientId).Scan(&pubKey)
    if err != nil {
        return nil, err
    }
    return pubKey, nil
}
