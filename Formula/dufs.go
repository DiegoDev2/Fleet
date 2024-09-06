package main

// Dufs - Static file server
// Homepage: https://github.com/sigoden/dufs

import (
	"fmt"
	
	"os/exec"
)

func installDufs() {
	// Método 1: Descargar y extraer .tar.gz
	dufs_tar_url := "https://github.com/sigoden/dufs/archive/refs/tags/v0.42.0.tar.gz"
	dufs_cmd_tar := exec.Command("curl", "-L", dufs_tar_url, "-o", "package.tar.gz")
	err := dufs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dufs_zip_url := "https://github.com/sigoden/dufs/archive/refs/tags/v0.42.0.zip"
	dufs_cmd_zip := exec.Command("curl", "-L", dufs_zip_url, "-o", "package.zip")
	err = dufs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dufs_bin_url := "https://github.com/sigoden/dufs/archive/refs/tags/v0.42.0.bin"
	dufs_cmd_bin := exec.Command("curl", "-L", dufs_bin_url, "-o", "binary.bin")
	err = dufs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dufs_src_url := "https://github.com/sigoden/dufs/archive/refs/tags/v0.42.0.src.tar.gz"
	dufs_cmd_src := exec.Command("curl", "-L", dufs_src_url, "-o", "source.tar.gz")
	err = dufs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dufs_cmd_direct := exec.Command("./binary")
	err = dufs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
