package main

// Curlie - Power of curl, ease of use of httpie
// Homepage: https://curlie.io

import (
	"fmt"
	
	"os/exec"
)

func installCurlie() {
	// Método 1: Descargar y extraer .tar.gz
	curlie_tar_url := "https://github.com/rs/curlie/archive/refs/tags/v1.7.2.tar.gz"
	curlie_cmd_tar := exec.Command("curl", "-L", curlie_tar_url, "-o", "package.tar.gz")
	err := curlie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curlie_zip_url := "https://github.com/rs/curlie/archive/refs/tags/v1.7.2.zip"
	curlie_cmd_zip := exec.Command("curl", "-L", curlie_zip_url, "-o", "package.zip")
	err = curlie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curlie_bin_url := "https://github.com/rs/curlie/archive/refs/tags/v1.7.2.bin"
	curlie_cmd_bin := exec.Command("curl", "-L", curlie_bin_url, "-o", "binary.bin")
	err = curlie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curlie_src_url := "https://github.com/rs/curlie/archive/refs/tags/v1.7.2.src.tar.gz"
	curlie_cmd_src := exec.Command("curl", "-L", curlie_src_url, "-o", "source.tar.gz")
	err = curlie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curlie_cmd_direct := exec.Command("./binary")
	err = curlie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
