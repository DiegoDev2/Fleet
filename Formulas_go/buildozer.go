package main

// Buildozer - Rewrite bazel BUILD files using standard commands
// Homepage: https://github.com/bazelbuild/buildtools

import (
	"fmt"
	
	"os/exec"
)

func installBuildozer() {
	// Método 1: Descargar y extraer .tar.gz
	buildozer_tar_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.tar.gz"
	buildozer_cmd_tar := exec.Command("curl", "-L", buildozer_tar_url, "-o", "package.tar.gz")
	err := buildozer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	buildozer_zip_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.zip"
	buildozer_cmd_zip := exec.Command("curl", "-L", buildozer_zip_url, "-o", "package.zip")
	err = buildozer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	buildozer_bin_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.bin"
	buildozer_cmd_bin := exec.Command("curl", "-L", buildozer_bin_url, "-o", "binary.bin")
	err = buildozer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	buildozer_src_url := "https://github.com/bazelbuild/buildtools/archive/refs/tags/v7.3.1.src.tar.gz"
	buildozer_cmd_src := exec.Command("curl", "-L", buildozer_src_url, "-o", "source.tar.gz")
	err = buildozer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	buildozer_cmd_direct := exec.Command("./binary")
	err = buildozer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
