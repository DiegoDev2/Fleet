package main

// OpenOcd - On-chip debugging, in-system programming and boundary-scan testing
// Homepage: https://openocd.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenOcd() {
	// Método 1: Descargar y extraer .tar.gz
	openocd_tar_url := "https://downloads.sourceforge.net/project/openocd/openocd/0.12.0/openocd-0.12.0.tar.bz2"
	openocd_cmd_tar := exec.Command("curl", "-L", openocd_tar_url, "-o", "package.tar.gz")
	err := openocd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openocd_zip_url := "https://downloads.sourceforge.net/project/openocd/openocd/0.12.0/openocd-0.12.0.tar.bz2"
	openocd_cmd_zip := exec.Command("curl", "-L", openocd_zip_url, "-o", "package.zip")
	err = openocd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openocd_bin_url := "https://downloads.sourceforge.net/project/openocd/openocd/0.12.0/openocd-0.12.0.tar.bz2"
	openocd_cmd_bin := exec.Command("curl", "-L", openocd_bin_url, "-o", "binary.bin")
	err = openocd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openocd_src_url := "https://downloads.sourceforge.net/project/openocd/openocd/0.12.0/openocd-0.12.0.tar.bz2"
	openocd_cmd_src := exec.Command("curl", "-L", openocd_src_url, "-o", "source.tar.gz")
	err = openocd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openocd_cmd_direct := exec.Command("./binary")
	err = openocd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: hidapi")
	exec.Command("latte", "install", "hidapi").Run()
	fmt.Println("Instalando dependencia: libftdi")
	exec.Command("latte", "install", "libftdi").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
