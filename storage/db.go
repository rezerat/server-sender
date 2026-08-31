package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool

func Сonnect() error {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Env not found \n%v", err)
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

func SaveDevice(nickname string, PubKey []byte) error{
    query := `INSERT INTO devices (client_id, public_key) VALUES ($1, $2)`
    _, err := pool.Exec(context.Background(), query, nickname, PubKey)
    return err

}

func GetPublicKey(clientId string) ([]byte, error) {
    pubKey := []byte{}
    query := `SELECT slice_pubKey WHERE client_id = $1`
    err := pool.QueryRow(context.Background(), query, clientId).Scan(&pubKey)
    if err != nil {
        return nil, err
    }
    return pubKey, nil
}
