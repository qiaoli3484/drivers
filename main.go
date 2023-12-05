package main

import (
	"drivers/ssd1306"
	"machine"
	"time"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		SDA: machine.GP4,
		SCL: machine.GP5,
	})
	oled := ssd1306.New(machine.I2C0, 0x3c)

	oled.Init()
	oled.Clear()
	oled.ShowChar(0, 0, 'F', 16)
	oled.ShowChar(8, 0, 'O', 16)
	oled.ShowChar(16, 0, 'x', 16)
	oled.ShowString(0, 2, []byte("hllowbbbbhllowbbbbhxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"), 16)
	for {
		time.Sleep(1 * time.Second)
	}
}

//tinygo flash -target pico
