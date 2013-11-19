package overlays

import (
	"os/exec"
)

func Compile() {
	// SPI Overlays
	exec.Command("dtc", "-O", "dtb", "-o", "overlays/ADAFRUIT-SPI0-00A0.dtbo", "-b", "o", "-@", "overlays/ADAFRUIT-SPI0-00A0.dts").Run()
	exec.Command("dtc", "-O", "dtb", "-o", "overlays/ADAFRUIT-SPI1-00A0.dtbo", "-b", "o", "-@", "overlays/ADAFRUIT-SPI1-00A0.dts").Run()
	// // UART Overlayss
	exec.Command("dtc", "-O", "dtb", "-o", "overlays/ADAFRUIT-UART1-00A0.dtbo", "-b", "o", "-@", "overlays/ADAFRUIT-UART1-00A0.dts").Run()
	exec.Command("dtc", "-O", "dtb", "-o", "overlays/ADAFRUIT-UART2-00A0.dtbo", "-b", "o", "-@", "overlays/ADAFRUIT-UART2-00A0.dts").Run()
	exec.Command("dtc", "-O", "dtb", "-o", "overlays/ADAFRUIT-UART4-00A0.dtbo", "-b", "o", "-@", "overlays/ADAFRUIT-UART4-00A0.dts").Run()
	exec.Command("dtc", "-O", "dtb", "-o", "overlays/ADAFRUIT-UART5-00A0.dtbo", "-b", "o", "-@", "overlays/ADAFRUIT-UART5-00A0.dts").Run()
}

func Copy() {
	// for fl in glob.glob("/lib/firmware/ADAFRUIT-SPI*-00A0.dtbo"):
	//     os.remove(fl)
	// for fl in glob.glob("/lib/firmware/ADAFRUIT-UART*-00A0.dtbo"):
	//     os.remove(fl)
	// // SPI Overlays
	// shutil.move("overlays/ADAFRUIT-SPI0-00A0.dtbo", "/lib/firmware/ADAFRUIT-SPI0-00A0.dtbo")
	// shutil.move("overlays/ADAFRUIT-SPI1-00A0.dtbo", "/lib/firmware/ADAFRUIT-SPI1-00A0.dtbo")
	// // UART Overlays
	// shutil.move("overlays/ADAFRUIT-UART1-00A0.dtbo", "/lib/firmware/ADAFRUIT-UART1-00A0.dtbo")
	// shutil.move("overlays/ADAFRUIT-UART2-00A0.dtbo", "/lib/firmware/ADAFRUIT-UART2-00A0.dtbo")
	// shutil.move("overlays/ADAFRUIT-UART4-00A0.dtbo", "/lib/firmware/ADAFRUIT-UART4-00A0.dtbo")
	// shutil.move("overlays/ADAFRUIT-UART5-00A0.dtbo", "/lib/firmware/ADAFRUIT-UART5-00A0.dtbo")
}
