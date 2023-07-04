// PDPI Stage 2 Bootloader

#define BOARD_PICO_FLASH_SPI_CLKDIV   4
#define BOARD_CMD_READ                0xeb
#define BOARD_QUAD_OK                 1
#define BOARD_QUAD_ENABLE_STATUS_BYTE 1
#define BOARD_QUAD_ENABLE_BIT_MASK    0x40
#define BOARD_SPLIT_STATUS_WRITE      0
#define BOARD_WAIT_CYCLES             4

#include "rp2040-boot-stage2.S"
