package main

// Megacmd - Command-line client for mega.co.nz storage service
// Homepage: https://github.com/t3rm1n4l/megacmd

import (
	"fmt"
	
	"os/exec"
)

func installMegacmd() {
	// Método 1: Descargar y extraer .tar.gz
	megacmd_tar_url := "https://github.com/t3rm1n4l/megacmd/archive/refs/tags/0.016.tar.gz"
	megacmd_cmd_tar := exec.Command("curl", "-L", megacmd_tar_url, "-o", "package.tar.gz")
	err := megacmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	megacmd_zip_url := "https://github.com/t3rm1n4l/megacmd/archive/refs/tags/0.016.zip"
	megacmd_cmd_zip := exec.Command("curl", "-L", megacmd_zip_url, "-o", "package.zip")
	err = megacmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	megacmd_bin_url := "https://github.com/t3rm1n4l/megacmd/archive/refs/tags/0.016.bin"
	megacmd_cmd_bin := exec.Command("curl", "-L", megacmd_bin_url, "-o", "binary.bin")
	err = megacmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	megacmd_src_url := "https://github.com/t3rm1n4l/megacmd/archive/refs/tags/0.016.src.tar.gz"
	megacmd_cmd_src := exec.Command("curl", "-L", megacmd_src_url, "-o", "source.tar.gz")
	err = megacmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	megacmd_cmd_direct := exec.Command("./binary")
	err = megacmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
