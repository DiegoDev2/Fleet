package main

// Yajl - Yet Another JSON Library
// Homepage: https://lloyd.github.io/yajl/

import (
	"fmt"
	
	"os/exec"
)

func installYajl() {
	// Método 1: Descargar y extraer .tar.gz
	yajl_tar_url := "https://github.com/lloyd/yajl/archive/refs/tags/2.1.0.tar.gz"
	yajl_cmd_tar := exec.Command("curl", "-L", yajl_tar_url, "-o", "package.tar.gz")
	err := yajl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yajl_zip_url := "https://github.com/lloyd/yajl/archive/refs/tags/2.1.0.zip"
	yajl_cmd_zip := exec.Command("curl", "-L", yajl_zip_url, "-o", "package.zip")
	err = yajl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yajl_bin_url := "https://github.com/lloyd/yajl/archive/refs/tags/2.1.0.bin"
	yajl_cmd_bin := exec.Command("curl", "-L", yajl_bin_url, "-o", "binary.bin")
	err = yajl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yajl_src_url := "https://github.com/lloyd/yajl/archive/refs/tags/2.1.0.src.tar.gz"
	yajl_cmd_src := exec.Command("curl", "-L", yajl_src_url, "-o", "source.tar.gz")
	err = yajl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yajl_cmd_direct := exec.Command("./binary")
	err = yajl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
