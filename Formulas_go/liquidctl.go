package main

// Liquidctl - Cross-platform tool and drivers for liquid coolers and other devices
// Homepage: https://github.com/liquidctl/liquidctl

import (
	"fmt"
	
	"os/exec"
)

func installLiquidctl() {
	// Método 1: Descargar y extraer .tar.gz
	liquidctl_tar_url := "https://files.pythonhosted.org/packages/99/d9/15bfe9dc11f2910b7483693b0bab16a382e5ad16cee657ff8133b7cae56d/liquidctl-1.13.0.tar.gz"
	liquidctl_cmd_tar := exec.Command("curl", "-L", liquidctl_tar_url, "-o", "package.tar.gz")
	err := liquidctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liquidctl_zip_url := "https://files.pythonhosted.org/packages/99/d9/15bfe9dc11f2910b7483693b0bab16a382e5ad16cee657ff8133b7cae56d/liquidctl-1.13.0.zip"
	liquidctl_cmd_zip := exec.Command("curl", "-L", liquidctl_zip_url, "-o", "package.zip")
	err = liquidctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liquidctl_bin_url := "https://files.pythonhosted.org/packages/99/d9/15bfe9dc11f2910b7483693b0bab16a382e5ad16cee657ff8133b7cae56d/liquidctl-1.13.0.bin"
	liquidctl_cmd_bin := exec.Command("curl", "-L", liquidctl_bin_url, "-o", "binary.bin")
	err = liquidctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liquidctl_src_url := "https://files.pythonhosted.org/packages/99/d9/15bfe9dc11f2910b7483693b0bab16a382e5ad16cee657ff8133b7cae56d/liquidctl-1.13.0.src.tar.gz"
	liquidctl_cmd_src := exec.Command("curl", "-L", liquidctl_src_url, "-o", "source.tar.gz")
	err = liquidctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liquidctl_cmd_direct := exec.Command("./binary")
	err = liquidctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hidapi")
exec.Command("latte", "install", "hidapi")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: i2c-tools")
exec.Command("latte", "install", "i2c-tools")
}
