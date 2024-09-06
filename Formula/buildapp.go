package main

// Buildapp - Creates executables with SBCL
// Homepage: https://www.xach.com/lisp/buildapp/

import (
	"fmt"
	
	"os/exec"
)

func installBuildapp() {
	// Método 1: Descargar y extraer .tar.gz
	buildapp_tar_url := "https://github.com/xach/buildapp/archive/refs/tags/release-1.5.6.tar.gz"
	buildapp_cmd_tar := exec.Command("curl", "-L", buildapp_tar_url, "-o", "package.tar.gz")
	err := buildapp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	buildapp_zip_url := "https://github.com/xach/buildapp/archive/refs/tags/release-1.5.6.zip"
	buildapp_cmd_zip := exec.Command("curl", "-L", buildapp_zip_url, "-o", "package.zip")
	err = buildapp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	buildapp_bin_url := "https://github.com/xach/buildapp/archive/refs/tags/release-1.5.6.bin"
	buildapp_cmd_bin := exec.Command("curl", "-L", buildapp_bin_url, "-o", "binary.bin")
	err = buildapp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	buildapp_src_url := "https://github.com/xach/buildapp/archive/refs/tags/release-1.5.6.src.tar.gz"
	buildapp_cmd_src := exec.Command("curl", "-L", buildapp_src_url, "-o", "source.tar.gz")
	err = buildapp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	buildapp_cmd_direct := exec.Command("./binary")
	err = buildapp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbcl")
	exec.Command("latte", "install", "sbcl").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
