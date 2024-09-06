package main

// Dub - Build tool for D projects
// Homepage: https://code.dlang.org/getting_started

import (
	"fmt"
	
	"os/exec"
)

func installDub() {
	// Método 1: Descargar y extraer .tar.gz
	dub_tar_url := "https://github.com/dlang/dub/archive/refs/tags/v1.38.1.tar.gz"
	dub_cmd_tar := exec.Command("curl", "-L", dub_tar_url, "-o", "package.tar.gz")
	err := dub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dub_zip_url := "https://github.com/dlang/dub/archive/refs/tags/v1.38.1.zip"
	dub_cmd_zip := exec.Command("curl", "-L", dub_zip_url, "-o", "package.zip")
	err = dub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dub_bin_url := "https://github.com/dlang/dub/archive/refs/tags/v1.38.1.bin"
	dub_cmd_bin := exec.Command("curl", "-L", dub_bin_url, "-o", "binary.bin")
	err = dub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dub_src_url := "https://github.com/dlang/dub/archive/refs/tags/v1.38.1.src.tar.gz"
	dub_cmd_src := exec.Command("curl", "-L", dub_src_url, "-o", "source.tar.gz")
	err = dub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dub_cmd_direct := exec.Command("./binary")
	err = dub_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ldc")
exec.Command("latte", "install", "ldc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
