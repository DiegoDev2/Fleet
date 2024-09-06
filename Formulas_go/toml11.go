package main

// Toml11 - TOML for Modern C++
// Homepage: https://github.com/ToruNiina/toml11

import (
	"fmt"
	
	"os/exec"
)

func installToml11() {
	// Método 1: Descargar y extraer .tar.gz
	toml11_tar_url := "https://github.com/ToruNiina/toml11/archive/refs/tags/v4.2.0.tar.gz"
	toml11_cmd_tar := exec.Command("curl", "-L", toml11_tar_url, "-o", "package.tar.gz")
	err := toml11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toml11_zip_url := "https://github.com/ToruNiina/toml11/archive/refs/tags/v4.2.0.zip"
	toml11_cmd_zip := exec.Command("curl", "-L", toml11_zip_url, "-o", "package.zip")
	err = toml11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toml11_bin_url := "https://github.com/ToruNiina/toml11/archive/refs/tags/v4.2.0.bin"
	toml11_cmd_bin := exec.Command("curl", "-L", toml11_bin_url, "-o", "binary.bin")
	err = toml11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toml11_src_url := "https://github.com/ToruNiina/toml11/archive/refs/tags/v4.2.0.src.tar.gz"
	toml11_cmd_src := exec.Command("curl", "-L", toml11_src_url, "-o", "source.tar.gz")
	err = toml11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toml11_cmd_direct := exec.Command("./binary")
	err = toml11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
