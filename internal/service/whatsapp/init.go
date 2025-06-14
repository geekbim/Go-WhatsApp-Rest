package whatsapp

import (
	"context"
	"fmt"
	"go_wa_rest/domain/service"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waCompanionReg"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var WhatsAppClient = make(map[string]*whatsmeow.Client)

type whatsAppService struct {
}

func NewWhatsAppService() service.WhatsAppService {
	return &whatsAppService{}
}

type MessageStatus struct {
	sync.RWMutex
	StatusMap map[string]string
}

var MessageStatuses = &MessageStatus{
	StatusMap: make(map[string]string),
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	case *events.Receipt:
		for _, receipt := range v.MessageIDs {
			messageID := receipt
			status := ""

			switch v.Type {
			case types.ReceiptTypeDelivered:
				status = "delivered"
			case types.ReceiptTypeRead:
				status = "read"
			case types.ReceiptTypePlayed:
				status = "opened"
			default:
				status = "sent"
			}

			MessageStatuses.Lock()
			MessageStatuses.StatusMap[messageID] = status
			MessageStatuses.Unlock()
		}
	}
}

func InitWhatsApp() *whatsmeow.Client {
	ctx := context.Background()
	dbLog := waLog.Stdout("Database", "DEBUG", true)

	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New(ctx, "sqlite3", "file:session.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	return client
}

func InitWhatsAppV2(device *store.Device, jid string) {
	ctx := context.Background()
	dbLog := waLog.Stdout("Database", "DEBUG", true)

	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New(ctx, "sqlite3", "file:session_v2.db?_foreign_keys=on&cache=shared&mode=rw", dbLog)
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
		store.DeviceProps.PlatformType = waCompanionReg.DeviceProps_DESKTOP.Enum()
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
