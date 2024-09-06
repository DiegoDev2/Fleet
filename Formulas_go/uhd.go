package main

// Uhd - Hardware driver for all USRP devices
// Homepage: https://files.ettus.com/manual/

import (
	"fmt"
	
	"os/exec"
)

func installUhd() {
	// Método 1: Descargar y extraer .tar.gz
	uhd_tar_url := "https://github.com/EttusResearch/uhd/archive/refs/tags/v4.7.0.0.tar.gz"
	uhd_cmd_tar := exec.Command("curl", "-L", uhd_tar_url, "-o", "package.tar.gz")
	err := uhd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uhd_zip_url := "https://github.com/EttusResearch/uhd/archive/refs/tags/v4.7.0.0.zip"
	uhd_cmd_zip := exec.Command("curl", "-L", uhd_zip_url, "-o", "package.zip")
	err = uhd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uhd_bin_url := "https://github.com/EttusResearch/uhd/archive/refs/tags/v4.7.0.0.bin"
	uhd_cmd_bin := exec.Command("curl", "-L", uhd_bin_url, "-o", "binary.bin")
	err = uhd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uhd_src_url := "https://github.com/EttusResearch/uhd/archive/refs/tags/v4.7.0.0.src.tar.gz"
	uhd_cmd_src := exec.Command("curl", "-L", uhd_src_url, "-o", "source.tar.gz")
	err = uhd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uhd_cmd_direct := exec.Command("./binary")
	err = uhd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
