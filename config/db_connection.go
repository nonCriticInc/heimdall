package config
import (
	"context"
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var DatabaseHostURL string
var DatabaseName string

var connectDB *bongo.Connection
var EntityCollection *bongo.Collection
var ApplicationCollection *bongo.Collection
var CertCollection *bongo.Collection
var ClientCollection *bongo.Collection
var ClientsOauthTokenCollection *bongo.Collection
var EtcdCollection *bongo.Collection
var OrganizationCollection *bongo.Collection
var PermissionCollection *bongo.Collection
var ResourceCollection *bongo.Collection
var RoleCollection *bongo.Collection
var UserCollection *bongo.Collection
var UsersOauthTokenCollection *bongo.Collection

func InitDBEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
	DatabaseHostURL = os.Getenv("DATABASE_MONGODB_HOST_URL")
	DatabaseName = os.Getenv("DATABASE_NAME")
}

// Connect Database
func InitDBConnection() {
	// DB Connect
	connection, err := CreateConnectionDB()
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
	connectDB = connection
}

// Initialize Database Collections
func InitDBCollections() {
	EntityCollection = connectDB.Collection("entityCollection")
	ApplicationCollection =connectDB.Collection("applicationCollection")
	CertCollection =connectDB.Collection("certCollection")
	ClientCollection =connectDB.Collection("clientCollection")
	ClientsOauthTokenCollection =connectDB.Collection("clientsOauthTokenCollection")
	EtcdCollection =connectDB.Collection("etcdCollection")
	OrganizationCollection =connectDB.Collection("organizationCollection")
	PermissionCollection =connectDB.Collection("permissionCollection")
	ResourceCollection =connectDB.Collection("resourceCollection")
	RoleCollection =connectDB.Collection("roleCollection")
	UserCollection =connectDB.Collection("userCollection")
	UsersOauthTokenCollection =connectDB.Collection("usersOauthTokenCollection")
}

func CreateConnectionDB() (*bongo.Connection, error) {
	config := &bongo.Config{
		ConnectionString: DatabaseHostURL,
		Database:         DatabaseName,
	}
	connection, err := bongo.Connect(config)
	return connection, err
}

func CloseConnectionDB(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}
