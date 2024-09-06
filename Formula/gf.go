package main

// Gf - App development framework of Golang
// Homepage: https://goframe.org

import (
	"fmt"
	
	"os/exec"
)

func installGf() {
	// Método 1: Descargar y extraer .tar.gz
	gf_tar_url := "https://github.com/gogf/gf/archive/refs/tags/v2.7.2.tar.gz"
	gf_cmd_tar := exec.Command("curl", "-L", gf_tar_url, "-o", "package.tar.gz")
	err := gf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gf_zip_url := "https://github.com/gogf/gf/archive/refs/tags/v2.7.2.zip"
	gf_cmd_zip := exec.Command("curl", "-L", gf_zip_url, "-o", "package.zip")
	err = gf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gf_bin_url := "https://github.com/gogf/gf/archive/refs/tags/v2.7.2.bin"
	gf_cmd_bin := exec.Command("curl", "-L", gf_bin_url, "-o", "binary.bin")
	err = gf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gf_src_url := "https://github.com/gogf/gf/archive/refs/tags/v2.7.2.src.tar.gz"
	gf_cmd_src := exec.Command("curl", "-L", gf_src_url, "-o", "source.tar.gz")
	err = gf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gf_cmd_direct := exec.Command("./binary")
	err = gf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
