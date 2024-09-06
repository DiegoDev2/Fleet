package main

// Tealdeer - Very fast implementation of tldr in Rust
// Homepage: https://github.com/dbrgn/tealdeer

import (
	"fmt"
	
	"os/exec"
)

func installTealdeer() {
	// Método 1: Descargar y extraer .tar.gz
	tealdeer_tar_url := "https://github.com/dbrgn/tealdeer/archive/refs/tags/v1.6.1.tar.gz"
	tealdeer_cmd_tar := exec.Command("curl", "-L", tealdeer_tar_url, "-o", "package.tar.gz")
	err := tealdeer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tealdeer_zip_url := "https://github.com/dbrgn/tealdeer/archive/refs/tags/v1.6.1.zip"
	tealdeer_cmd_zip := exec.Command("curl", "-L", tealdeer_zip_url, "-o", "package.zip")
	err = tealdeer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tealdeer_bin_url := "https://github.com/dbrgn/tealdeer/archive/refs/tags/v1.6.1.bin"
	tealdeer_cmd_bin := exec.Command("curl", "-L", tealdeer_bin_url, "-o", "binary.bin")
	err = tealdeer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tealdeer_src_url := "https://github.com/dbrgn/tealdeer/archive/refs/tags/v1.6.1.src.tar.gz"
	tealdeer_cmd_src := exec.Command("curl", "-L", tealdeer_src_url, "-o", "source.tar.gz")
	err = tealdeer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tealdeer_cmd_direct := exec.Command("./binary")
	err = tealdeer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
