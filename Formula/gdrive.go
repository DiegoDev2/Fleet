package main

// Gdrive - Google Drive CLI Client
// Homepage: https://github.com/glotlabs/gdrive

import (
	"fmt"
	
	"os/exec"
)

func installGdrive() {
	// Método 1: Descargar y extraer .tar.gz
	gdrive_tar_url := "https://github.com/glotlabs/gdrive/archive/refs/tags/3.9.1.tar.gz"
	gdrive_cmd_tar := exec.Command("curl", "-L", gdrive_tar_url, "-o", "package.tar.gz")
	err := gdrive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdrive_zip_url := "https://github.com/glotlabs/gdrive/archive/refs/tags/3.9.1.zip"
	gdrive_cmd_zip := exec.Command("curl", "-L", gdrive_zip_url, "-o", "package.zip")
	err = gdrive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdrive_bin_url := "https://github.com/glotlabs/gdrive/archive/refs/tags/3.9.1.bin"
	gdrive_cmd_bin := exec.Command("curl", "-L", gdrive_bin_url, "-o", "binary.bin")
	err = gdrive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdrive_src_url := "https://github.com/glotlabs/gdrive/archive/refs/tags/3.9.1.src.tar.gz"
	gdrive_cmd_src := exec.Command("curl", "-L", gdrive_src_url, "-o", "source.tar.gz")
	err = gdrive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdrive_cmd_direct := exec.Command("./binary")
	err = gdrive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
