package main

// Pixiewps - Offline Wi-Fi Protected Setup brute-force utility
// Homepage: https://github.com/wiire-a/pixiewps

import (
	"fmt"
	
	"os/exec"
)

func installPixiewps() {
	// Método 1: Descargar y extraer .tar.gz
	pixiewps_tar_url := "https://github.com/wiire-a/pixiewps/releases/download/v1.4.2/pixiewps-1.4.2.tar.xz"
	pixiewps_cmd_tar := exec.Command("curl", "-L", pixiewps_tar_url, "-o", "package.tar.gz")
	err := pixiewps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pixiewps_zip_url := "https://github.com/wiire-a/pixiewps/releases/download/v1.4.2/pixiewps-1.4.2.tar.xz"
	pixiewps_cmd_zip := exec.Command("curl", "-L", pixiewps_zip_url, "-o", "package.zip")
	err = pixiewps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pixiewps_bin_url := "https://github.com/wiire-a/pixiewps/releases/download/v1.4.2/pixiewps-1.4.2.tar.xz"
	pixiewps_cmd_bin := exec.Command("curl", "-L", pixiewps_bin_url, "-o", "binary.bin")
	err = pixiewps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pixiewps_src_url := "https://github.com/wiire-a/pixiewps/releases/download/v1.4.2/pixiewps-1.4.2.tar.xz"
	pixiewps_cmd_src := exec.Command("curl", "-L", pixiewps_src_url, "-o", "source.tar.gz")
	err = pixiewps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pixiewps_cmd_direct := exec.Command("./binary")
	err = pixiewps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
