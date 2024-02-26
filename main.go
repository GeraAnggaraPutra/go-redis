package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// User adalah struktur untuk data pengguna
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Inisialisasi Redis Client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Tes koneksi Redis
	if err := pingRedis(rdb); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	fmt.Println("Redis Connection: Successful")

	// Contoh data pengguna
	users := []User{
		{ID: "1", Name: "John Doe", Age: 30},
		{ID: "2", Name: "Jane Smith", Age: 24},
		{ID: "3", Name: "Lorem Ipsum", Age: 25},
		{ID: "4", Name: "Dolor Sit", Age: 21},
		{ID: "5", Name: "Amet", Age: 22},
	}

	// Simpan data pengguna ke Redis
	for _, user := range users {
		if err := setUser(rdb, user); err != nil {
			log.Fatalf("Error setting user: %v", err)
		}
	}

	// Dapatkan dan cetak data pengguna dari Redis
	printUsers(rdb)

	// Perbarui data pengguna di Redis
	users[0].Age = 35
	if err := updateUser(rdb, users[0]); err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	// Dapatkan dan cetak data pengguna yang diperbarui dari Redis
	printUsers(rdb)

	// Hapus data pengguna dari Redis
	if err := deleteUser(rdb, users[1].ID); err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}

	// Dapatkan dan cetak data pengguna yang tersisa dari Redis
	printUsers(rdb)
}

// pingRedis digunakan untuk menguji koneksi ke Redis
func pingRedis(rdb *redis.Client) error {
	_, err := rdb.Ping(ctx).Result()
	return err
}

// setUser digunakan untuk menyimpan data pengguna ke Redis
func setUser(rdb *redis.Client, user User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, user.ID, userJSON, 0).Err()
}

// getUser digunakan untuk mendapatkan data pengguna dari Redis berdasarkan ID
func getUser(rdb *redis.Client, id string) (User, error) {
	val, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal([]byte(val), &user)
	return user, err
}

// printUsers digunakan untuk mencetak semua data pengguna dari Redis
func printUsers(rdb *redis.Client) {
	users, err := getUsers(rdb)
	if err != nil {
		log.Fatalf("Error getting users: %v", err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

// getUsers digunakan untuk mendapatkan semua data pengguna dari Redis
func getUsers(rdb *redis.Client) ([]User, error) {
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	var users []User
	for _, key := range keys {
		user, err := getUser(rdb, key)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// updateUser digunakan untuk memperbarui data pengguna di Redis
func updateUser(rdb *redis.Client, user User) error {
	return setUser(rdb, user)
}

// deleteUser digunakan untuk menghapus data pengguna dari Redis berdasarkan ID
func deleteUser(rdb *redis.Client, id string) error {
	return rdb.Del(ctx, id).Err()
}
