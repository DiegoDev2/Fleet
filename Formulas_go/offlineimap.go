package main

// Offlineimap - Synchronizes emails between two repositories
// Homepage: https://github.com/OfflineIMAP/offlineimap3

import (
	"fmt"
	
	"os/exec"
)

func installOfflineimap() {
	// Método 1: Descargar y extraer .tar.gz
	offlineimap_tar_url := "https://github.com/OfflineIMAP/offlineimap3/archive/refs/tags/v8.0.0.tar.gz"
	offlineimap_cmd_tar := exec.Command("curl", "-L", offlineimap_tar_url, "-o", "package.tar.gz")
	err := offlineimap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	offlineimap_zip_url := "https://github.com/OfflineIMAP/offlineimap3/archive/refs/tags/v8.0.0.zip"
	offlineimap_cmd_zip := exec.Command("curl", "-L", offlineimap_zip_url, "-o", "package.zip")
	err = offlineimap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	offlineimap_bin_url := "https://github.com/OfflineIMAP/offlineimap3/archive/refs/tags/v8.0.0.bin"
	offlineimap_cmd_bin := exec.Command("curl", "-L", offlineimap_bin_url, "-o", "binary.bin")
	err = offlineimap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	offlineimap_src_url := "https://github.com/OfflineIMAP/offlineimap3/archive/refs/tags/v8.0.0.src.tar.gz"
	offlineimap_cmd_src := exec.Command("curl", "-L", offlineimap_src_url, "-o", "source.tar.gz")
	err = offlineimap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	offlineimap_cmd_direct := exec.Command("./binary")
	err = offlineimap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
