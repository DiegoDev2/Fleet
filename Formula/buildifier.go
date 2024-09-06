package main

// Buildifier - Format bazel BUILD files with a standard convention
// Homepage: https://github.com/bazelbuild/buildtools

import (
	"fmt"
	
	"os/exec"
)

func installBuildifier() {
	// Método 1: Descargar y extraer .tar.gz
	buildifier_tar_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.tar.gz"
	buildifier_cmd_tar := exec.Command("curl", "-L", buildifier_tar_url, "-o", "package.tar.gz")
	err := buildifier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	buildifier_zip_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.zip"
	buildifier_cmd_zip := exec.Command("curl", "-L", buildifier_zip_url, "-o", "package.zip")
	err = buildifier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	buildifier_bin_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.bin"
	buildifier_cmd_bin := exec.Command("curl", "-L", buildifier_bin_url, "-o", "binary.bin")
	err = buildifier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	buildifier_src_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.src.tar.gz"
	buildifier_cmd_src := exec.Command("curl", "-L", buildifier_src_url, "-o", "source.tar.gz")
	err = buildifier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	buildifier_cmd_direct := exec.Command("./binary")
	err = buildifier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
