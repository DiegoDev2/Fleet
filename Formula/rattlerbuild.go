package main

// RattlerBuild - Universal conda package builder
// Homepage: https://github.com/prefix-dev/rattler-build

import (
	"fmt"
	
	"os/exec"
)

func installRattlerBuild() {
	// Método 1: Descargar y extraer .tar.gz
	rattlerbuild_tar_url := "https://github.com/prefix-dev/rattler-build/archive/refs/tags/v0.20.0.tar.gz"
	rattlerbuild_cmd_tar := exec.Command("curl", "-L", rattlerbuild_tar_url, "-o", "package.tar.gz")
	err := rattlerbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rattlerbuild_zip_url := "https://github.com/prefix-dev/rattler-build/archive/refs/tags/v0.20.0.zip"
	rattlerbuild_cmd_zip := exec.Command("curl", "-L", rattlerbuild_zip_url, "-o", "package.zip")
	err = rattlerbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rattlerbuild_bin_url := "https://github.com/prefix-dev/rattler-build/archive/refs/tags/v0.20.0.bin"
	rattlerbuild_cmd_bin := exec.Command("curl", "-L", rattlerbuild_bin_url, "-o", "binary.bin")
	err = rattlerbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rattlerbuild_src_url := "https://github.com/prefix-dev/rattler-build/archive/refs/tags/v0.20.0.src.tar.gz"
	rattlerbuild_cmd_src := exec.Command("curl", "-L", rattlerbuild_src_url, "-o", "source.tar.gz")
	err = rattlerbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rattlerbuild_cmd_direct := exec.Command("./binary")
	err = rattlerbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
