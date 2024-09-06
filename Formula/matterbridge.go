package main

// Matterbridge - Protocol bridge for multiple chat platforms
// Homepage: https://github.com/42wim/matterbridge

import (
	"fmt"
	
	"os/exec"
)

func installMatterbridge() {
	// Método 1: Descargar y extraer .tar.gz
	matterbridge_tar_url := "https://github.com/42wim/matterbridge/archive/refs/tags/v1.26.0.tar.gz"
	matterbridge_cmd_tar := exec.Command("curl", "-L", matterbridge_tar_url, "-o", "package.tar.gz")
	err := matterbridge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	matterbridge_zip_url := "https://github.com/42wim/matterbridge/archive/refs/tags/v1.26.0.zip"
	matterbridge_cmd_zip := exec.Command("curl", "-L", matterbridge_zip_url, "-o", "package.zip")
	err = matterbridge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	matterbridge_bin_url := "https://github.com/42wim/matterbridge/archive/refs/tags/v1.26.0.bin"
	matterbridge_cmd_bin := exec.Command("curl", "-L", matterbridge_bin_url, "-o", "binary.bin")
	err = matterbridge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	matterbridge_src_url := "https://github.com/42wim/matterbridge/archive/refs/tags/v1.26.0.src.tar.gz"
	matterbridge_cmd_src := exec.Command("curl", "-L", matterbridge_src_url, "-o", "source.tar.gz")
	err = matterbridge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	matterbridge_cmd_direct := exec.Command("./binary")
	err = matterbridge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
