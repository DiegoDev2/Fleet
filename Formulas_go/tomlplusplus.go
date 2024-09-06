package main

// Tomlplusplus - Header-only TOML config file parser and serializer for C++17
// Homepage: https://marzer.github.io/tomlplusplus/

import (
	"fmt"
	
	"os/exec"
)

func installTomlplusplus() {
	// Método 1: Descargar y extraer .tar.gz
	tomlplusplus_tar_url := "https://github.com/marzer/tomlplusplus/archive/refs/tags/v3.4.0.tar.gz"
	tomlplusplus_cmd_tar := exec.Command("curl", "-L", tomlplusplus_tar_url, "-o", "package.tar.gz")
	err := tomlplusplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomlplusplus_zip_url := "https://github.com/marzer/tomlplusplus/archive/refs/tags/v3.4.0.zip"
	tomlplusplus_cmd_zip := exec.Command("curl", "-L", tomlplusplus_zip_url, "-o", "package.zip")
	err = tomlplusplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomlplusplus_bin_url := "https://github.com/marzer/tomlplusplus/archive/refs/tags/v3.4.0.bin"
	tomlplusplus_cmd_bin := exec.Command("curl", "-L", tomlplusplus_bin_url, "-o", "binary.bin")
	err = tomlplusplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomlplusplus_src_url := "https://github.com/marzer/tomlplusplus/archive/refs/tags/v3.4.0.src.tar.gz"
	tomlplusplus_cmd_src := exec.Command("curl", "-L", tomlplusplus_src_url, "-o", "source.tar.gz")
	err = tomlplusplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomlplusplus_cmd_direct := exec.Command("./binary")
	err = tomlplusplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
