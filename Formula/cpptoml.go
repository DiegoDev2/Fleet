package main

// Cpptoml - Header-only library for parsing TOML
// Homepage: https://github.com/skystrife/cpptoml

import (
	"fmt"
	
	"os/exec"
)

func installCpptoml() {
	// Método 1: Descargar y extraer .tar.gz
	cpptoml_tar_url := "https://github.com/skystrife/cpptoml/archive/refs/tags/v0.1.1.tar.gz"
	cpptoml_cmd_tar := exec.Command("curl", "-L", cpptoml_tar_url, "-o", "package.tar.gz")
	err := cpptoml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpptoml_zip_url := "https://github.com/skystrife/cpptoml/archive/refs/tags/v0.1.1.zip"
	cpptoml_cmd_zip := exec.Command("curl", "-L", cpptoml_zip_url, "-o", "package.zip")
	err = cpptoml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpptoml_bin_url := "https://github.com/skystrife/cpptoml/archive/refs/tags/v0.1.1.bin"
	cpptoml_cmd_bin := exec.Command("curl", "-L", cpptoml_bin_url, "-o", "binary.bin")
	err = cpptoml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpptoml_src_url := "https://github.com/skystrife/cpptoml/archive/refs/tags/v0.1.1.src.tar.gz"
	cpptoml_cmd_src := exec.Command("curl", "-L", cpptoml_src_url, "-o", "source.tar.gz")
	err = cpptoml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpptoml_cmd_direct := exec.Command("./binary")
	err = cpptoml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
