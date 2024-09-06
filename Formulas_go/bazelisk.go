package main

// Bazelisk - User-friendly launcher for Bazel
// Homepage: https://github.com/bazelbuild/bazelisk/

import (
	"fmt"
	
	"os/exec"
)

func installBazelisk() {
	// Método 1: Descargar y extraer .tar.gz
	bazelisk_tar_url := "https://github.com/bazelbuild/bazelisk/archive/refs/tags/v1.20.0.tar.gz"
	bazelisk_cmd_tar := exec.Command("curl", "-L", bazelisk_tar_url, "-o", "package.tar.gz")
	err := bazelisk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bazelisk_zip_url := "https://github.com/bazelbuild/bazelisk/archive/refs/tags/v1.20.0.zip"
	bazelisk_cmd_zip := exec.Command("curl", "-L", bazelisk_zip_url, "-o", "package.zip")
	err = bazelisk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bazelisk_bin_url := "https://github.com/bazelbuild/bazelisk/archive/refs/tags/v1.20.0.bin"
	bazelisk_cmd_bin := exec.Command("curl", "-L", bazelisk_bin_url, "-o", "binary.bin")
	err = bazelisk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bazelisk_src_url := "https://github.com/bazelbuild/bazelisk/archive/refs/tags/v1.20.0.src.tar.gz"
	bazelisk_cmd_src := exec.Command("curl", "-L", bazelisk_src_url, "-o", "source.tar.gz")
	err = bazelisk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bazelisk_cmd_direct := exec.Command("./binary")
	err = bazelisk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
