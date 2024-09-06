package main

// Trash - CLI tool that moves files or folder to the trash
// Homepage: https://hasseg.org/trash/

import (
	"fmt"
	
	"os/exec"
)

func installTrash() {
	// Método 1: Descargar y extraer .tar.gz
	trash_tar_url := "https://github.com/ali-rantakari/trash/archive/refs/tags/v0.9.2.tar.gz"
	trash_cmd_tar := exec.Command("curl", "-L", trash_tar_url, "-o", "package.tar.gz")
	err := trash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trash_zip_url := "https://github.com/ali-rantakari/trash/archive/refs/tags/v0.9.2.zip"
	trash_cmd_zip := exec.Command("curl", "-L", trash_zip_url, "-o", "package.zip")
	err = trash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trash_bin_url := "https://github.com/ali-rantakari/trash/archive/refs/tags/v0.9.2.bin"
	trash_cmd_bin := exec.Command("curl", "-L", trash_bin_url, "-o", "binary.bin")
	err = trash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trash_src_url := "https://github.com/ali-rantakari/trash/archive/refs/tags/v0.9.2.src.tar.gz"
	trash_cmd_src := exec.Command("curl", "-L", trash_src_url, "-o", "source.tar.gz")
	err = trash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trash_cmd_direct := exec.Command("./binary")
	err = trash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
