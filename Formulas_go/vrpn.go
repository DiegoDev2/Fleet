package main

// Vrpn - Virtual reality peripheral network
// Homepage: https://github.com/vrpn/vrpn/wiki

import (
	"fmt"
	
	"os/exec"
)

func installVrpn() {
	// Método 1: Descargar y extraer .tar.gz
	vrpn_tar_url := "https://github.com/vrpn/vrpn/releases/download/version_07.35/vrpn_07.35.zip"
	vrpn_cmd_tar := exec.Command("curl", "-L", vrpn_tar_url, "-o", "package.tar.gz")
	err := vrpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vrpn_zip_url := "https://github.com/vrpn/vrpn/releases/download/version_07.35/vrpn_07.35.zip"
	vrpn_cmd_zip := exec.Command("curl", "-L", vrpn_zip_url, "-o", "package.zip")
	err = vrpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vrpn_bin_url := "https://github.com/vrpn/vrpn/releases/download/version_07.35/vrpn_07.35.zip"
	vrpn_cmd_bin := exec.Command("curl", "-L", vrpn_bin_url, "-o", "binary.bin")
	err = vrpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vrpn_src_url := "https://github.com/vrpn/vrpn/releases/download/version_07.35/vrpn_07.35.zip"
	vrpn_cmd_src := exec.Command("curl", "-L", vrpn_src_url, "-o", "source.tar.gz")
	err = vrpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vrpn_cmd_direct := exec.Command("./binary")
	err = vrpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
