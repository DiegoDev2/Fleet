package main

// Cmdshelf - Better scripting life with cmdshelf
// Homepage: https://github.com/toshi0383/cmdshelf

import (
	"fmt"
	
	"os/exec"
)

func installCmdshelf() {
	// Método 1: Descargar y extraer .tar.gz
	cmdshelf_tar_url := "https://github.com/toshi0383/cmdshelf/archive/refs/tags/2.0.2.tar.gz"
	cmdshelf_cmd_tar := exec.Command("curl", "-L", cmdshelf_tar_url, "-o", "package.tar.gz")
	err := cmdshelf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmdshelf_zip_url := "https://github.com/toshi0383/cmdshelf/archive/refs/tags/2.0.2.zip"
	cmdshelf_cmd_zip := exec.Command("curl", "-L", cmdshelf_zip_url, "-o", "package.zip")
	err = cmdshelf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmdshelf_bin_url := "https://github.com/toshi0383/cmdshelf/archive/refs/tags/2.0.2.bin"
	cmdshelf_cmd_bin := exec.Command("curl", "-L", cmdshelf_bin_url, "-o", "binary.bin")
	err = cmdshelf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmdshelf_src_url := "https://github.com/toshi0383/cmdshelf/archive/refs/tags/2.0.2.src.tar.gz"
	cmdshelf_cmd_src := exec.Command("curl", "-L", cmdshelf_src_url, "-o", "source.tar.gz")
	err = cmdshelf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmdshelf_cmd_direct := exec.Command("./binary")
	err = cmdshelf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
