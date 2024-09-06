package main

// Premake - Write once, build anywhere Lua-based build system
// Homepage: https://premake.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installPremake() {
	// Método 1: Descargar y extraer .tar.gz
	premake_tar_url := "https://github.com/premake/premake-core/releases/download/v5.0.0-beta2/premake-5.0.0-beta2-src.zip"
	premake_cmd_tar := exec.Command("curl", "-L", premake_tar_url, "-o", "package.tar.gz")
	err := premake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	premake_zip_url := "https://github.com/premake/premake-core/releases/download/v5.0.0-beta2/premake-5.0.0-beta2-src.zip"
	premake_cmd_zip := exec.Command("curl", "-L", premake_zip_url, "-o", "package.zip")
	err = premake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	premake_bin_url := "https://github.com/premake/premake-core/releases/download/v5.0.0-beta2/premake-5.0.0-beta2-src.zip"
	premake_cmd_bin := exec.Command("curl", "-L", premake_bin_url, "-o", "binary.bin")
	err = premake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	premake_src_url := "https://github.com/premake/premake-core/releases/download/v5.0.0-beta2/premake-5.0.0-beta2-src.zip"
	premake_cmd_src := exec.Command("curl", "-L", premake_src_url, "-o", "source.tar.gz")
	err = premake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	premake_cmd_direct := exec.Command("./binary")
	err = premake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
