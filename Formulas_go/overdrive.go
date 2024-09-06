package main

// Overdrive - Bash script to download mp3s from the OverDrive audiobook service
// Homepage: https://github.com/chbrown/overdrive

import (
	"fmt"
	
	"os/exec"
)

func installOverdrive() {
	// Método 1: Descargar y extraer .tar.gz
	overdrive_tar_url := "https://github.com/chbrown/overdrive/archive/refs/tags/2.4.0.tar.gz"
	overdrive_cmd_tar := exec.Command("curl", "-L", overdrive_tar_url, "-o", "package.tar.gz")
	err := overdrive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	overdrive_zip_url := "https://github.com/chbrown/overdrive/archive/refs/tags/2.4.0.zip"
	overdrive_cmd_zip := exec.Command("curl", "-L", overdrive_zip_url, "-o", "package.zip")
	err = overdrive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	overdrive_bin_url := "https://github.com/chbrown/overdrive/archive/refs/tags/2.4.0.bin"
	overdrive_cmd_bin := exec.Command("curl", "-L", overdrive_bin_url, "-o", "binary.bin")
	err = overdrive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	overdrive_src_url := "https://github.com/chbrown/overdrive/archive/refs/tags/2.4.0.src.tar.gz"
	overdrive_cmd_src := exec.Command("curl", "-L", overdrive_src_url, "-o", "source.tar.gz")
	err = overdrive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	overdrive_cmd_direct := exec.Command("./binary")
	err = overdrive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
