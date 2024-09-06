package main

// Frizbee - Throw a tag at and it comes back with a checksum
// Homepage: https://github.com/stacklok/frizbee

import (
	"fmt"
	
	"os/exec"
)

func installFrizbee() {
	// Método 1: Descargar y extraer .tar.gz
	frizbee_tar_url := "https://github.com/stacklok/frizbee/archive/refs/tags/v0.1.2.tar.gz"
	frizbee_cmd_tar := exec.Command("curl", "-L", frizbee_tar_url, "-o", "package.tar.gz")
	err := frizbee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frizbee_zip_url := "https://github.com/stacklok/frizbee/archive/refs/tags/v0.1.2.zip"
	frizbee_cmd_zip := exec.Command("curl", "-L", frizbee_zip_url, "-o", "package.zip")
	err = frizbee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frizbee_bin_url := "https://github.com/stacklok/frizbee/archive/refs/tags/v0.1.2.bin"
	frizbee_cmd_bin := exec.Command("curl", "-L", frizbee_bin_url, "-o", "binary.bin")
	err = frizbee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frizbee_src_url := "https://github.com/stacklok/frizbee/archive/refs/tags/v0.1.2.src.tar.gz"
	frizbee_cmd_src := exec.Command("curl", "-L", frizbee_src_url, "-o", "source.tar.gz")
	err = frizbee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frizbee_cmd_direct := exec.Command("./binary")
	err = frizbee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
