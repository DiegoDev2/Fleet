package main

// Sad - CLI search and replace | Space Age seD
// Homepage: https://github.com/ms-jpq/sad

import (
	"fmt"
	
	"os/exec"
)

func installSad() {
	// Método 1: Descargar y extraer .tar.gz
	sad_tar_url := "https://github.com/ms-jpq/sad/archive/refs/tags/v0.4.31.tar.gz"
	sad_cmd_tar := exec.Command("curl", "-L", sad_tar_url, "-o", "package.tar.gz")
	err := sad_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sad_zip_url := "https://github.com/ms-jpq/sad/archive/refs/tags/v0.4.31.zip"
	sad_cmd_zip := exec.Command("curl", "-L", sad_zip_url, "-o", "package.zip")
	err = sad_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sad_bin_url := "https://github.com/ms-jpq/sad/archive/refs/tags/v0.4.31.bin"
	sad_cmd_bin := exec.Command("curl", "-L", sad_bin_url, "-o", "binary.bin")
	err = sad_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sad_src_url := "https://github.com/ms-jpq/sad/archive/refs/tags/v0.4.31.src.tar.gz"
	sad_cmd_src := exec.Command("curl", "-L", sad_src_url, "-o", "source.tar.gz")
	err = sad_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sad_cmd_direct := exec.Command("./binary")
	err = sad_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
