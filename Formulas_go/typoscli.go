package main

// TyposCli - Source code spell checker
// Homepage: https://github.com/crate-ci/typos

import (
	"fmt"
	
	"os/exec"
)

func installTyposCli() {
	// Método 1: Descargar y extraer .tar.gz
	typoscli_tar_url := "https://github.com/crate-ci/typos/archive/refs/tags/v1.24.5.tar.gz"
	typoscli_cmd_tar := exec.Command("curl", "-L", typoscli_tar_url, "-o", "package.tar.gz")
	err := typoscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typoscli_zip_url := "https://github.com/crate-ci/typos/archive/refs/tags/v1.24.5.zip"
	typoscli_cmd_zip := exec.Command("curl", "-L", typoscli_zip_url, "-o", "package.zip")
	err = typoscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typoscli_bin_url := "https://github.com/crate-ci/typos/archive/refs/tags/v1.24.5.bin"
	typoscli_cmd_bin := exec.Command("curl", "-L", typoscli_bin_url, "-o", "binary.bin")
	err = typoscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typoscli_src_url := "https://github.com/crate-ci/typos/archive/refs/tags/v1.24.5.src.tar.gz"
	typoscli_cmd_src := exec.Command("curl", "-L", typoscli_src_url, "-o", "source.tar.gz")
	err = typoscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typoscli_cmd_direct := exec.Command("./binary")
	err = typoscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
