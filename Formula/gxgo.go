package main

// GxGo - Tool to use with the gx package manager for packages written in go
// Homepage: https://github.com/whyrusleeping/gx-go

import (
	"fmt"
	
	"os/exec"
)

func installGxGo() {
	// Método 1: Descargar y extraer .tar.gz
	gxgo_tar_url := "https://github.com/whyrusleeping/gx-go/archive/refs/tags/v1.9.0.tar.gz"
	gxgo_cmd_tar := exec.Command("curl", "-L", gxgo_tar_url, "-o", "package.tar.gz")
	err := gxgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gxgo_zip_url := "https://github.com/whyrusleeping/gx-go/archive/refs/tags/v1.9.0.zip"
	gxgo_cmd_zip := exec.Command("curl", "-L", gxgo_zip_url, "-o", "package.zip")
	err = gxgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gxgo_bin_url := "https://github.com/whyrusleeping/gx-go/archive/refs/tags/v1.9.0.bin"
	gxgo_cmd_bin := exec.Command("curl", "-L", gxgo_bin_url, "-o", "binary.bin")
	err = gxgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gxgo_src_url := "https://github.com/whyrusleeping/gx-go/archive/refs/tags/v1.9.0.src.tar.gz"
	gxgo_cmd_src := exec.Command("curl", "-L", gxgo_src_url, "-o", "source.tar.gz")
	err = gxgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gxgo_cmd_direct := exec.Command("./binary")
	err = gxgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
