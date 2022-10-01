package whatsapp

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var (
	WhatsAppDatastore *sqlstore.Container
	WhatsAppClient    = make(map[string]*whatsmeow.Client)
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

// func init() {
// 	dbType := os.Getenv("WHATSAPP_DATASTORE_TYPE")

// 	dbURI := os.Getenv("WHATSAPP_DATASTORE_URI")

// 	datastore, err := sqlstore.New(dbType, dbURI, nil)
// 	if err != nil {
// 		log.Fatal("Error Connect WhatsApp Client Datastore")
// 	}

// 	WhatsAppDatastore = datastore
// }

func InitWhatsApp() *whatsmeow.Client {
	dbLog := waLog.Stdout("Database", "DEBUG", true)

	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	return client
}

func InitWhatsAppV2(device *store.Device, jid string) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)

	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	if WhatsAppClient[jid] == nil {
		if device == nil {
			// Initialize New WhatsApp Client Device in Datastore
			device = container.NewDevice()
		}

		// Set Client Properties
		store.DeviceProps.Os = proto.String("Go WhatsApp Multi-Device REST")
		store.DeviceProps.PlatformType = waproto.DeviceProps_DESKTOP.Enum()
		store.DeviceProps.RequireFullSync = proto.Bool(false)

		// Set Client Versions
		store.DeviceProps.Version.Primary = proto.Uint32(uint32(version.Major))
		store.DeviceProps.Version.Secondary = proto.Uint32(uint32(version.Minor))
		store.DeviceProps.Version.Tertiary = proto.Uint32(uint32(version.Patch))

		// Initialize New WhatsApp Client
		// And Save it to The Map
		WhatsAppClient[jid] = whatsmeow.NewClient(device, nil)

		// Set WhatsApp Client Auto Reconnect
		WhatsAppClient[jid].EnableAutoReconnect = true

		// Set WhatsApp Client Auto Trust Identity
		WhatsAppClient[jid].AutoTrustIdentity = true
	}
}
