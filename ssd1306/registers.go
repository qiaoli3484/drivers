package ssd1306

const (
	Address        = 0x3D
	Address_128_32 = 0x3C

	SETCONTRAST         = 0x81
	DISPLAYALLON_RESUME = 0xA4
	DISPLAYALLON        = 0xA5
	NORMALDISPLAY       = 0xA6 //正常显示
	INVERTDISPLAY       = 0xA7 //反色显示
	DISPLAYOFF          = 0xAE //关闭显示
	DISPLAYON           = 0xAF //开启显示
	SETDISPLAYOFFSET    = 0xD3
	SETCOMPINS          = 0xDA //设置列引脚硬件配置
	SETVCOMDETECT       = 0xDB //设置 vcomh反压值
	SETDISPLAYCLOCKDIV  = 0xD5 //设置显示刷新率
	SETPRECHARGE        = 0xD9 //设置预充电周期
	SETMULTIPLEX        = 0xA8 //设置复用率(0-63)
	//0x80  1-f 越大速度越快
	SETLOWCOLUMN                         = 0x00 //设置低位栏
	SETHIGHCOLUMN                        = 0x10 //设置高位栏
	SETSTARTLINE                         = 0x40
	MEMORYMODE                           = 0x20
	COLUMNADDR                           = 0x21
	PAGEADDR                             = 0x22
	COMSCANINC                           = 0xC0 //上下镜像
	COMSCANDEC                           = 0xC8 //正常显示
	SEGREMAP                             = 0xA0 //左右镜像
	NOSEGREMAP                           = 0xA1 //正常显示
	CHARGEPUMP                           = 0x8D //设置电荷泵开启
	ACTIVATE_SCROLL                      = 0x2F //激活滚屏
	DEACTIVATE_SCROLL                    = 0x2E
	SET_VERTICAL_SCROLL_AREA             = 0xA3
	RIGHT_HORIZONTAL_SCROLL              = 0x26 //右侧_水平_滚动
	LEFT_HORIZONTAL_SCROLL               = 0x27 //左_水平_滚动
	VERTICAL_AND_RIGHT_HORIZONTAL_SCROLL = 0x29 //垂直和右侧水平滚动
	VERTICAL_AND_LEFT_HORIZONTAL_SCROLL  = 0x2A //垂直和左水平滚动

	EXTERNALVCC  VccMode = 0x1
	SWITCHCAPVCC VccMode = 0x2
)
