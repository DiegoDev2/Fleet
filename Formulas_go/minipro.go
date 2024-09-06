package main

// Minipro - Open controller for the MiniPRO TL866xx series of chip programmers
// Homepage: https://gitlab.com/DavidGriffith/minipro/

import (
	"fmt"
	
	"os/exec"
)

func installMinipro() {
	// Método 1: Descargar y extraer .tar.gz
	minipro_tar_url := "https://gitlab.com/DavidGriffith/minipro/-/archive/0.7.1/minipro-0.7.1.tar.gz"
	minipro_cmd_tar := exec.Command("curl", "-L", minipro_tar_url, "-o", "package.tar.gz")
	err := minipro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minipro_zip_url := "https://gitlab.com/DavidGriffith/minipro/-/archive/0.7.1/minipro-0.7.1.zip"
	minipro_cmd_zip := exec.Command("curl", "-L", minipro_zip_url, "-o", "package.zip")
	err = minipro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minipro_bin_url := "https://gitlab.com/DavidGriffith/minipro/-/archive/0.7.1/minipro-0.7.1.bin"
	minipro_cmd_bin := exec.Command("curl", "-L", minipro_bin_url, "-o", "binary.bin")
	err = minipro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minipro_src_url := "https://gitlab.com/DavidGriffith/minipro/-/archive/0.7.1/minipro-0.7.1.src.tar.gz"
	minipro_cmd_src := exec.Command("curl", "-L", minipro_src_url, "-o", "source.tar.gz")
	err = minipro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minipro_cmd_direct := exec.Command("./binary")
	err = minipro_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: srecord")
exec.Command("latte", "install", "srecord")
}
