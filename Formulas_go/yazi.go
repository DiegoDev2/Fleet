package main

// Yazi - Blazing fast terminal file manager written in Rust, based on async I/O
// Homepage: https://github.com/sxyazi/yazi

import (
	"fmt"
	
	"os/exec"
)

func installYazi() {
	// Método 1: Descargar y extraer .tar.gz
	yazi_tar_url := "https://github.com/sxyazi/yazi/archive/refs/tags/v0.3.3.tar.gz"
	yazi_cmd_tar := exec.Command("curl", "-L", yazi_tar_url, "-o", "package.tar.gz")
	err := yazi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yazi_zip_url := "https://github.com/sxyazi/yazi/archive/refs/tags/v0.3.3.zip"
	yazi_cmd_zip := exec.Command("curl", "-L", yazi_zip_url, "-o", "package.zip")
	err = yazi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yazi_bin_url := "https://github.com/sxyazi/yazi/archive/refs/tags/v0.3.3.bin"
	yazi_cmd_bin := exec.Command("curl", "-L", yazi_bin_url, "-o", "binary.bin")
	err = yazi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yazi_src_url := "https://github.com/sxyazi/yazi/archive/refs/tags/v0.3.3.src.tar.gz"
	yazi_cmd_src := exec.Command("curl", "-L", yazi_src_url, "-o", "source.tar.gz")
	err = yazi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yazi_cmd_direct := exec.Command("./binary")
	err = yazi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
