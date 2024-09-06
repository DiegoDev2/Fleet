package main

// Dust - More intuitive version of du in rust
// Homepage: https://github.com/bootandy/dust

import (
	"fmt"
	
	"os/exec"
)

func installDust() {
	// Método 1: Descargar y extraer .tar.gz
	dust_tar_url := "https://github.com/bootandy/dust/archive/refs/tags/v1.1.1.tar.gz"
	dust_cmd_tar := exec.Command("curl", "-L", dust_tar_url, "-o", "package.tar.gz")
	err := dust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dust_zip_url := "https://github.com/bootandy/dust/archive/refs/tags/v1.1.1.zip"
	dust_cmd_zip := exec.Command("curl", "-L", dust_zip_url, "-o", "package.zip")
	err = dust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dust_bin_url := "https://github.com/bootandy/dust/archive/refs/tags/v1.1.1.bin"
	dust_cmd_bin := exec.Command("curl", "-L", dust_bin_url, "-o", "binary.bin")
	err = dust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dust_src_url := "https://github.com/bootandy/dust/archive/refs/tags/v1.1.1.src.tar.gz"
	dust_cmd_src := exec.Command("curl", "-L", dust_src_url, "-o", "source.tar.gz")
	err = dust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dust_cmd_direct := exec.Command("./binary")
	err = dust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
