//go:build pdpi

package machine

const (
	LED              Pin = GPIO24
	LightSensor      Pin = ADC0
	VBUSDischarge    Pin = GPIO6
	PowerGoodDisable Pin = GPIO19

	xoscFreq = 12 // MHz

	I2C0_SDA_PIN = GPIO20
	I2C0_SCL_PIN = GPIO21

	I2C1_SDA_PIN = GPIO2
	I2C1_SCL_PIN = GPIO3

	usb_STRING_PRODUCT      = "PDPi"
	usb_STRING_MANUFACTURER = "oxplot"

	UART_TX_PIN = GPIO0
	UART_RX_PIN = GPIO1

	SPI0_SCK_PIN = GPIO18
	SPI0_SDO_PIN = GPIO19 // Tx
	SPI0_SDI_PIN = GPIO16 // Rx

	SPI1_SCK_PIN = GPIO10
	SPI1_SDO_PIN = GPIO11 // Tx
	SPI1_SDI_PIN = GPIO12 // Rx
)

var (
	usb_VID uint16 = 0x2E8A
	usb_PID uint16 = 0x000A
)
