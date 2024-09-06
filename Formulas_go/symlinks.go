package main

// Symlinks - Symbolic link maintenance utility
// Homepage: https://github.com/brandt/symlinks

import (
	"fmt"
	
	"os/exec"
)

func installSymlinks() {
	// Método 1: Descargar y extraer .tar.gz
	symlinks_tar_url := "https://github.com/brandt/symlinks/archive/refs/tags/v1.4.3.tar.gz"
	symlinks_cmd_tar := exec.Command("curl", "-L", symlinks_tar_url, "-o", "package.tar.gz")
	err := symlinks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	symlinks_zip_url := "https://github.com/brandt/symlinks/archive/refs/tags/v1.4.3.zip"
	symlinks_cmd_zip := exec.Command("curl", "-L", symlinks_zip_url, "-o", "package.zip")
	err = symlinks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	symlinks_bin_url := "https://github.com/brandt/symlinks/archive/refs/tags/v1.4.3.bin"
	symlinks_cmd_bin := exec.Command("curl", "-L", symlinks_bin_url, "-o", "binary.bin")
	err = symlinks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	symlinks_src_url := "https://github.com/brandt/symlinks/archive/refs/tags/v1.4.3.src.tar.gz"
	symlinks_cmd_src := exec.Command("curl", "-L", symlinks_src_url, "-o", "source.tar.gz")
	err = symlinks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	symlinks_cmd_direct := exec.Command("./binary")
	err = symlinks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
