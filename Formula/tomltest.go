package main

// TomlTest - Language agnostic test suite for TOML parsers
// Homepage: https://github.com/toml-lang/toml-test

import (
	"fmt"
	
	"os/exec"
)

func installTomlTest() {
	// Método 1: Descargar y extraer .tar.gz
	tomltest_tar_url := "https://github.com/toml-lang/toml-test/archive/refs/tags/v1.5.0.tar.gz"
	tomltest_cmd_tar := exec.Command("curl", "-L", tomltest_tar_url, "-o", "package.tar.gz")
	err := tomltest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomltest_zip_url := "https://github.com/toml-lang/toml-test/archive/refs/tags/v1.5.0.zip"
	tomltest_cmd_zip := exec.Command("curl", "-L", tomltest_zip_url, "-o", "package.zip")
	err = tomltest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomltest_bin_url := "https://github.com/toml-lang/toml-test/archive/refs/tags/v1.5.0.bin"
	tomltest_cmd_bin := exec.Command("curl", "-L", tomltest_bin_url, "-o", "binary.bin")
	err = tomltest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomltest_src_url := "https://github.com/toml-lang/toml-test/archive/refs/tags/v1.5.0.src.tar.gz"
	tomltest_cmd_src := exec.Command("curl", "-L", tomltest_src_url, "-o", "source.tar.gz")
	err = tomltest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomltest_cmd_direct := exec.Command("./binary")
	err = tomltest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
