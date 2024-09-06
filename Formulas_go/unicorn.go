package main

// Unicorn - Lightweight multi-architecture CPU emulation framework
// Homepage: https://www.unicorn-engine.org/

import (
	"fmt"
	
	"os/exec"
)

func installUnicorn() {
	// Método 1: Descargar y extraer .tar.gz
	unicorn_tar_url := "https://github.com/unicorn-engine/unicorn/archive/refs/tags/2.0.1.post1.tar.gz"
	unicorn_cmd_tar := exec.Command("curl", "-L", unicorn_tar_url, "-o", "package.tar.gz")
	err := unicorn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unicorn_zip_url := "https://github.com/unicorn-engine/unicorn/archive/refs/tags/2.0.1.post1.zip"
	unicorn_cmd_zip := exec.Command("curl", "-L", unicorn_zip_url, "-o", "package.zip")
	err = unicorn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unicorn_bin_url := "https://github.com/unicorn-engine/unicorn/archive/refs/tags/2.0.1.post1.bin"
	unicorn_cmd_bin := exec.Command("curl", "-L", unicorn_bin_url, "-o", "binary.bin")
	err = unicorn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unicorn_src_url := "https://github.com/unicorn-engine/unicorn/archive/refs/tags/2.0.1.post1.src.tar.gz"
	unicorn_cmd_src := exec.Command("curl", "-L", unicorn_src_url, "-o", "source.tar.gz")
	err = unicorn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unicorn_cmd_direct := exec.Command("./binary")
	err = unicorn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
