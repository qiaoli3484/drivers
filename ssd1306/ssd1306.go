package ssd1306

type VccMode uint8

type Bus interface {
	Tx(addr uint16, w []byte, r []byte) error
}

type ssd1306 struct {
	address uint16
	bus     Bus
}

func Init(bus Bus, address uint16) *ssd1306 {
	return &ssd1306{
		bus:     bus,
		address: address,
	}
}

func (s *ssd1306) write_Cmd(cmd uint8) {
	/*buf := make([]uint8, len(data)+1)
	buf[0] = reg
	copy(buf[1:], data)*/
	s.bus.Tx(s.address, []byte{0x00, cmd}, nil)
}

func (s *ssd1306) write_Data(cmd uint8) {
	s.bus.Tx(s.address, []byte{0x40, cmd}, nil)
}

//清屏函数,清完屏,整个屏幕是黑色的!和没点亮一样
func (s *ssd1306) Clear() {

	for i := 0; i < 8; i++ {
		s.write_Cmd(0xb0 + uint8(i)) //设置页地址（0~7）
		s.write_Cmd(0x00)            //设置显示位置—列低地址
		s.write_Cmd(0x10)            //设置显示位置—列高地址
		for n := 0; n < 128; n++ {
			s.write_Data(0)
		}
	} //更新显示
}

func (s *ssd1306) Run() {

	//SSD1306复位之后，默认的是页寻址方式

	s.write_Cmd(0xAE) //关闭显示

	s.write_Cmd(0x20) //设置内存地址模式
	s.write_Cmd(0x10) //00水平地址模式，01垂直地址模式 ，02页显示模式

	//s.write_Cmd(0x00); // ---set low column address,初始化设置了没用,因为OLED_SetPos函数中会重设
	//s.write_Cmd(0x10); // ---set high column address,初始化设置了没用,因为OLED_SetPos函数中会重设

	s.write_Cmd(0x40) //--set start line address,从RAM中哪一行起读取显示内容
	s.write_Cmd(0xb0) // 设置起始页的地址模式 b0-b7

	s.write_Cmd(0x81) //设置对比度
	s.write_Cmd(0xff) //亮度调节 0x00~0xff 256级

	s.write_Cmd(0xa1) // 0xa0左右反置 0xa1正常
	s.write_Cmd(0xc8) // 0xc0上下反置 0xc8正常

	s.write_Cmd(0xa6) //设置显示方式;正常显示 a7反相显示;

	s.write_Cmd(0xa8) //--set multiplex ratio(1 to 64)复用率为 1~64
	s.write_Cmd(0x3F) //00h-3fh

	s.write_Cmd(0xa4) //全局显示开启;0xa4正常,0xa5无视命令点亮全屏

	s.write_Cmd(0xd3) //-设置屏幕上下偏移 offset
	s.write_Cmd(0x00) //-00h-ffh

	s.write_Cmd(0xd5) //设置时钟分频因子,震荡频率
	s.write_Cmd(0x80) //[3:0],分频因子;[7:4],震荡频率10h-f0h 越大越快

	s.write_Cmd(0xd9) //--set pre-charge period
	s.write_Cmd(0xF1) //

	s.write_Cmd(0xda) //--set com pins hardware configuration
	s.write_Cmd(0x12) //12864 0x12，12832  0x02

	s.write_Cmd(0xdb) //--set vcomh
	s.write_Cmd(0x30) //0x20,0.77xVcc

	s.write_Cmd(0x8d) //设置电荷泵开关
	s.write_Cmd(0x14) //开

	s.write_Cmd(0xaf) //开启显示
}

//坐标设置：也就是在哪里显示
func (s *ssd1306) OLED_Set_Pixel(x uint8, y uint8) {
	//以下3个寄存器只在页寻址的模式下有效
	s.write_Cmd(0xb0 + x)
	s.write_Cmd(((y & 0xf0) >> 4) | 0x10) //列高位地址设置
	s.write_Cmd(0x0f & y)                 //列低位地址设置
}

/*
func (s *ssd1306) ShowStr(x uint8, y uint8, ch []uint8, TextSize uint8) {
	//unsigned char c = 0,i = 0,j = 0,k = 0;
	if TextSize == 1 {
		for j := range ch {
			c := ch[j] - 32
			for i := 0; i < 6; i++ {
				s.OLED_Set_Pixel(x, y+uint8(j*6+i))
				s.write_Data(F6x8[c][i])
			}
		}
	} else if TextSize == 2 {
		for j := range ch {
			if ch[j] == '\r' {
				continue
			} else if ch[j] == '\n' {
				x++
				y = 0
			}
			c := ch[j] - 32
			//先构建上半部分
			for i := 0; i < 8; i++ {
				s.OLED_Set_Pixel(x, y+uint8(j*8+i))
				s.write_Data(F8X16[int(c)*16+i]) //上半截
				s.OLED_Set_Pixel(x+1, y+uint8(j*8+i))
				s.write_Data(F8X16[int(c)*16+i+8]) //下半截
			}
		}
	}
}*/
