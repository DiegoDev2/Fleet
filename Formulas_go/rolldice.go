package main

// Rolldice - Rolls an amount of virtual dice
// Homepage: https://github.com/sstrickl/rolldice

import (
	"fmt"
	
	"os/exec"
)

func installRolldice() {
	// Método 1: Descargar y extraer .tar.gz
	rolldice_tar_url := "https://github.com/sstrickl/rolldice/archive/refs/tags/v1.16.tar.gz"
	rolldice_cmd_tar := exec.Command("curl", "-L", rolldice_tar_url, "-o", "package.tar.gz")
	err := rolldice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rolldice_zip_url := "https://github.com/sstrickl/rolldice/archive/refs/tags/v1.16.zip"
	rolldice_cmd_zip := exec.Command("curl", "-L", rolldice_zip_url, "-o", "package.zip")
	err = rolldice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rolldice_bin_url := "https://github.com/sstrickl/rolldice/archive/refs/tags/v1.16.bin"
	rolldice_cmd_bin := exec.Command("curl", "-L", rolldice_bin_url, "-o", "binary.bin")
	err = rolldice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rolldice_src_url := "https://github.com/sstrickl/rolldice/archive/refs/tags/v1.16.src.tar.gz"
	rolldice_cmd_src := exec.Command("curl", "-L", rolldice_src_url, "-o", "source.tar.gz")
	err = rolldice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rolldice_cmd_direct := exec.Command("./binary")
	err = rolldice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
