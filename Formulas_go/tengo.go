package main

// Tengo - Fast script language for Go
// Homepage: https://tengolang.com

import (
	"fmt"
	
	"os/exec"
)

func installTengo() {
	// Método 1: Descargar y extraer .tar.gz
	tengo_tar_url := "https://github.com/d5/tengo/archive/refs/tags/v2.17.0.tar.gz"
	tengo_cmd_tar := exec.Command("curl", "-L", tengo_tar_url, "-o", "package.tar.gz")
	err := tengo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tengo_zip_url := "https://github.com/d5/tengo/archive/refs/tags/v2.17.0.zip"
	tengo_cmd_zip := exec.Command("curl", "-L", tengo_zip_url, "-o", "package.zip")
	err = tengo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tengo_bin_url := "https://github.com/d5/tengo/archive/refs/tags/v2.17.0.bin"
	tengo_cmd_bin := exec.Command("curl", "-L", tengo_bin_url, "-o", "binary.bin")
	err = tengo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tengo_src_url := "https://github.com/d5/tengo/archive/refs/tags/v2.17.0.src.tar.gz"
	tengo_cmd_src := exec.Command("curl", "-L", tengo_src_url, "-o", "source.tar.gz")
	err = tengo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tengo_cmd_direct := exec.Command("./binary")
	err = tengo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
