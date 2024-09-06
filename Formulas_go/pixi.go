package main

// Pixi - Package management made easy
// Homepage: https://pixi.sh

import (
	"fmt"
	
	"os/exec"
)

func installPixi() {
	// Método 1: Descargar y extraer .tar.gz
	pixi_tar_url := "https://github.com/prefix-dev/pixi/archive/refs/tags/v0.29.0.tar.gz"
	pixi_cmd_tar := exec.Command("curl", "-L", pixi_tar_url, "-o", "package.tar.gz")
	err := pixi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pixi_zip_url := "https://github.com/prefix-dev/pixi/archive/refs/tags/v0.29.0.zip"
	pixi_cmd_zip := exec.Command("curl", "-L", pixi_zip_url, "-o", "package.zip")
	err = pixi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pixi_bin_url := "https://github.com/prefix-dev/pixi/archive/refs/tags/v0.29.0.bin"
	pixi_cmd_bin := exec.Command("curl", "-L", pixi_bin_url, "-o", "binary.bin")
	err = pixi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pixi_src_url := "https://github.com/prefix-dev/pixi/archive/refs/tags/v0.29.0.src.tar.gz"
	pixi_cmd_src := exec.Command("curl", "-L", pixi_src_url, "-o", "source.tar.gz")
	err = pixi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pixi_cmd_direct := exec.Command("./binary")
	err = pixi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
