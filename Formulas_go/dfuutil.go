package main

// DfuUtil - USB programmer
// Homepage: https://dfu-util.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDfuUtil() {
	// Método 1: Descargar y extraer .tar.gz
	dfuutil_tar_url := "https://downloads.sourceforge.net/project/dfu-util/dfu-util-0.11.tar.gz"
	dfuutil_cmd_tar := exec.Command("curl", "-L", dfuutil_tar_url, "-o", "package.tar.gz")
	err := dfuutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dfuutil_zip_url := "https://downloads.sourceforge.net/project/dfu-util/dfu-util-0.11.zip"
	dfuutil_cmd_zip := exec.Command("curl", "-L", dfuutil_zip_url, "-o", "package.zip")
	err = dfuutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dfuutil_bin_url := "https://downloads.sourceforge.net/project/dfu-util/dfu-util-0.11.bin"
	dfuutil_cmd_bin := exec.Command("curl", "-L", dfuutil_bin_url, "-o", "binary.bin")
	err = dfuutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dfuutil_src_url := "https://downloads.sourceforge.net/project/dfu-util/dfu-util-0.11.src.tar.gz"
	dfuutil_cmd_src := exec.Command("curl", "-L", dfuutil_src_url, "-o", "source.tar.gz")
	err = dfuutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dfuutil_cmd_direct := exec.Command("./binary")
	err = dfuutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
