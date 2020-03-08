package store

import (
	"github.com/ExploratoryEngineering/air-quality-sensor-node/server/pkg/model"
)

// Store defines the persistence interface.
type Store interface {

	// ############################################################
	//                     Device
	// ############################################################

	// PutDevice adds a new device to the database.  Note that the ID
	// field is not assigned by the underlying database, but is
	// assumed to have a value.
	PutDevice(d *model.Device) error

	// GetDevice gets device by ID.
	GetDevice(id string) (*model.Device, error)

	// UpdateDevice updates device in database
	UpdateDevice(d *model.Device) error

	// DeleteDevice deletes device from database.
	DeleteDevice(id string) error

	// ListDevices lists devices
	ListDevices(offset int, limit int) ([]model.Device, error)

	// ############################################################
	//                     Calibration
	// ############################################################

	// PutCal adds a new calibration entry.  Note that the device
	// referred to must already exist in the database.
	PutCal(c *model.Cal) (int64, error)

	// GetCal gets a calibration entry by id
	GetCal(id int64) (*model.Cal, error)

	// UpdateCal updates calibration entry
	UpdateCal(c *model.Cal) error

	// DeleteCal deletes calibration entry.
	DeleteCal(id int64) error

	// ListCal lists calibration data ordered by DeviceID and ValidFrom in ascending order.
	ListCals(offset int, limit int) ([]model.Cal, error)

	// ListCalForDevice lists calibration data for device ordered by
	// validFrom date in descending order.
	ListCalsForDevice(deviceID string) ([]model.Cal, error)

	// ############################################################
	//                     Message
	// ############################################################
	PutMessage(m *model.Message) error
	GetMessage(id int64) (*model.Message, error)
	ListMessages(offset int, limit int) ([]model.Message, error)

	// Close the database
	Close() error
}
