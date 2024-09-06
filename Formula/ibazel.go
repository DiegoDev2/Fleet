package main

// Ibazel - Tools for building Bazel targets when source files change
// Homepage: https://github.com/bazelbuild/bazel-watcher

import (
	"fmt"
	
	"os/exec"
)

func installIbazel() {
	// Método 1: Descargar y extraer .tar.gz
	ibazel_tar_url := "https://github.com/bazelbuild/bazel-watcher/archive/refs/tags/v0.25.3.tar.gz"
	ibazel_cmd_tar := exec.Command("curl", "-L", ibazel_tar_url, "-o", "package.tar.gz")
	err := ibazel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ibazel_zip_url := "https://github.com/bazelbuild/bazel-watcher/archive/refs/tags/v0.25.3.zip"
	ibazel_cmd_zip := exec.Command("curl", "-L", ibazel_zip_url, "-o", "package.zip")
	err = ibazel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ibazel_bin_url := "https://github.com/bazelbuild/bazel-watcher/archive/refs/tags/v0.25.3.bin"
	ibazel_cmd_bin := exec.Command("curl", "-L", ibazel_bin_url, "-o", "binary.bin")
	err = ibazel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ibazel_src_url := "https://github.com/bazelbuild/bazel-watcher/archive/refs/tags/v0.25.3.src.tar.gz"
	ibazel_cmd_src := exec.Command("curl", "-L", ibazel_src_url, "-o", "source.tar.gz")
	err = ibazel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ibazel_cmd_direct := exec.Command("./binary")
	err = ibazel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bazelisk")
	exec.Command("latte", "install", "bazelisk").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
