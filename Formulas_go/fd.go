package main

// Fd - Simple, fast and user-friendly alternative to find
// Homepage: https://github.com/sharkdp/fd

import (
	"fmt"
	
	"os/exec"
)

func installFd() {
	// Método 1: Descargar y extraer .tar.gz
	fd_tar_url := "https://github.com/sharkdp/fd/archive/refs/tags/v10.2.0.tar.gz"
	fd_cmd_tar := exec.Command("curl", "-L", fd_tar_url, "-o", "package.tar.gz")
	err := fd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fd_zip_url := "https://github.com/sharkdp/fd/archive/refs/tags/v10.2.0.zip"
	fd_cmd_zip := exec.Command("curl", "-L", fd_zip_url, "-o", "package.zip")
	err = fd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fd_bin_url := "https://github.com/sharkdp/fd/archive/refs/tags/v10.2.0.bin"
	fd_cmd_bin := exec.Command("curl", "-L", fd_bin_url, "-o", "binary.bin")
	err = fd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fd_src_url := "https://github.com/sharkdp/fd/archive/refs/tags/v10.2.0.src.tar.gz"
	fd_cmd_src := exec.Command("curl", "-L", fd_src_url, "-o", "source.tar.gz")
	err = fd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fd_cmd_direct := exec.Command("./binary")
	err = fd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
