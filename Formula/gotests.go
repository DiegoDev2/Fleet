package main

// Gotests - Automatically generate Go test boilerplate from your source code
// Homepage: https://github.com/cweill/gotests

import (
	"fmt"
	
	"os/exec"
)

func installGotests() {
	// Método 1: Descargar y extraer .tar.gz
	gotests_tar_url := "https://github.com/cweill/gotests/archive/refs/tags/v1.6.0.tar.gz"
	gotests_cmd_tar := exec.Command("curl", "-L", gotests_tar_url, "-o", "package.tar.gz")
	err := gotests_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotests_zip_url := "https://github.com/cweill/gotests/archive/refs/tags/v1.6.0.zip"
	gotests_cmd_zip := exec.Command("curl", "-L", gotests_zip_url, "-o", "package.zip")
	err = gotests_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotests_bin_url := "https://github.com/cweill/gotests/archive/refs/tags/v1.6.0.bin"
	gotests_cmd_bin := exec.Command("curl", "-L", gotests_bin_url, "-o", "binary.bin")
	err = gotests_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotests_src_url := "https://github.com/cweill/gotests/archive/refs/tags/v1.6.0.src.tar.gz"
	gotests_cmd_src := exec.Command("curl", "-L", gotests_src_url, "-o", "source.tar.gz")
	err = gotests_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotests_cmd_direct := exec.Command("./binary")
	err = gotests_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
