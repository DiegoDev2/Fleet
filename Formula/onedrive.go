package main

// Onedrive - Folder synchronization with OneDrive
// Homepage: https://github.com/abraunegg/onedrive

import (
	"fmt"
	
	"os/exec"
)

func installOnedrive() {
	// Método 1: Descargar y extraer .tar.gz
	onedrive_tar_url := "https://github.com/abraunegg/onedrive/archive/refs/tags/v2.4.25.tar.gz"
	onedrive_cmd_tar := exec.Command("curl", "-L", onedrive_tar_url, "-o", "package.tar.gz")
	err := onedrive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onedrive_zip_url := "https://github.com/abraunegg/onedrive/archive/refs/tags/v2.4.25.zip"
	onedrive_cmd_zip := exec.Command("curl", "-L", onedrive_zip_url, "-o", "package.zip")
	err = onedrive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onedrive_bin_url := "https://github.com/abraunegg/onedrive/archive/refs/tags/v2.4.25.bin"
	onedrive_cmd_bin := exec.Command("curl", "-L", onedrive_bin_url, "-o", "binary.bin")
	err = onedrive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onedrive_src_url := "https://github.com/abraunegg/onedrive/archive/refs/tags/v2.4.25.src.tar.gz"
	onedrive_cmd_src := exec.Command("curl", "-L", onedrive_src_url, "-o", "source.tar.gz")
	err = onedrive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onedrive_cmd_direct := exec.Command("./binary")
	err = onedrive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ldc")
	exec.Command("latte", "install", "ldc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
