package main

// Luau - Fast, safe, gradually typed embeddable scripting language derived from Lua
// Homepage: https://luau-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installLuau() {
	// Método 1: Descargar y extraer .tar.gz
	luau_tar_url := "https://github.com/luau-lang/luau/archive/refs/tags/0.641.tar.gz"
	luau_cmd_tar := exec.Command("curl", "-L", luau_tar_url, "-o", "package.tar.gz")
	err := luau_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luau_zip_url := "https://github.com/luau-lang/luau/archive/refs/tags/0.641.zip"
	luau_cmd_zip := exec.Command("curl", "-L", luau_zip_url, "-o", "package.zip")
	err = luau_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luau_bin_url := "https://github.com/luau-lang/luau/archive/refs/tags/0.641.bin"
	luau_cmd_bin := exec.Command("curl", "-L", luau_bin_url, "-o", "binary.bin")
	err = luau_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luau_src_url := "https://github.com/luau-lang/luau/archive/refs/tags/0.641.src.tar.gz"
	luau_cmd_src := exec.Command("curl", "-L", luau_src_url, "-o", "source.tar.gz")
	err = luau_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luau_cmd_direct := exec.Command("./binary")
	err = luau_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
