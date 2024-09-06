package main

// DotenvLinter - Lightning-fast linter for .env files written in Rust
// Homepage: https://dotenv-linter.github.io

import (
	"fmt"
	
	"os/exec"
)

func installDotenvLinter() {
	// Método 1: Descargar y extraer .tar.gz
	dotenvlinter_tar_url := "https://github.com/dotenv-linter/dotenv-linter/archive/refs/tags/v3.3.0.tar.gz"
	dotenvlinter_cmd_tar := exec.Command("curl", "-L", dotenvlinter_tar_url, "-o", "package.tar.gz")
	err := dotenvlinter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dotenvlinter_zip_url := "https://github.com/dotenv-linter/dotenv-linter/archive/refs/tags/v3.3.0.zip"
	dotenvlinter_cmd_zip := exec.Command("curl", "-L", dotenvlinter_zip_url, "-o", "package.zip")
	err = dotenvlinter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dotenvlinter_bin_url := "https://github.com/dotenv-linter/dotenv-linter/archive/refs/tags/v3.3.0.bin"
	dotenvlinter_cmd_bin := exec.Command("curl", "-L", dotenvlinter_bin_url, "-o", "binary.bin")
	err = dotenvlinter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dotenvlinter_src_url := "https://github.com/dotenv-linter/dotenv-linter/archive/refs/tags/v3.3.0.src.tar.gz"
	dotenvlinter_cmd_src := exec.Command("curl", "-L", dotenvlinter_src_url, "-o", "source.tar.gz")
	err = dotenvlinter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dotenvlinter_cmd_direct := exec.Command("./binary")
	err = dotenvlinter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
