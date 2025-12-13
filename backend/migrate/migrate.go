package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jmoiron/sqlx"
	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/models"
)

var schema = `
-- USERS
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    
    -- Authentication & Identity
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    
    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- TOPICS (relies on the users table due to creator_id)
CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,

    title VARCHAR(50) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    
    -- Foreign Key: Links to the Users table
    creator_id INT NOT NULL REFERENCES users(id) ON DELETE RESTRICT, 

    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE 
);

-- POSTS 
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,

    title VARCHAR(50) NOT NULL,
    body TEXT NOT NULL,
    
    -- Foreign Key: Links to the Users table
    creator_id INT NOT NULL REFERENCES users(id) ON DELETE RESTRICT, 

    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE 
);

-- COMMENTS 
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,

    body TEXT NOT NULL,
    
    -- Foreign Key: Links to the Users table
	post_id INT NOT NULL REFERENCES posts(id) ON DELETE RESTRICT,
    creator_id INT NOT NULL REFERENCES users(id) ON DELETE RESTRICT, 

    -- Timestamps
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE 
);`

var insertSampleQuery = `
-- Insert a random user and topic
insert into users (id, username, email, password_hash)
values (1, 'rahul', 'rahul@gmail.com', '123hash');
insert into topics (name, description, creator_id)
values ('charity', 'this is to talk about charity', 1);
`

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.MustExec(schema)
	//initializers.DB.MustExec(insertSampleQuery)

	topics := []models.Topic{}

	err := initializers.DB.Select(&topics, "select * from topics order by id asc")
	if err != nil {
		log.Println("Error selecting posts: ", err)
	}
	fmt.Println(topics)

	fmt.Println("Connected to DB")

}
