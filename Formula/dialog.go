package main

// Dialog - Display user-friendly message boxes from shell scripts
// Homepage: https://invisible-island.net/dialog/

import (
	"fmt"
	
	"os/exec"
)

func installDialog() {
	// Método 1: Descargar y extraer .tar.gz
	dialog_tar_url := "https://invisible-mirror.net/archives/dialog/dialog-1.3-20240619.tgz"
	dialog_cmd_tar := exec.Command("curl", "-L", dialog_tar_url, "-o", "package.tar.gz")
	err := dialog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dialog_zip_url := "https://invisible-mirror.net/archives/dialog/dialog-1.3-20240619.tgz"
	dialog_cmd_zip := exec.Command("curl", "-L", dialog_zip_url, "-o", "package.zip")
	err = dialog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dialog_bin_url := "https://invisible-mirror.net/archives/dialog/dialog-1.3-20240619.tgz"
	dialog_cmd_bin := exec.Command("curl", "-L", dialog_bin_url, "-o", "binary.bin")
	err = dialog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dialog_src_url := "https://invisible-mirror.net/archives/dialog/dialog-1.3-20240619.tgz"
	dialog_cmd_src := exec.Command("curl", "-L", dialog_src_url, "-o", "source.tar.gz")
	err = dialog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dialog_cmd_direct := exec.Command("./binary")
	err = dialog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
