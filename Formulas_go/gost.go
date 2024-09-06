package main

// Gost - GO Simple Tunnel - a simple tunnel written in golang
// Homepage: https://github.com/ginuerzh/gost

import (
	"fmt"
	
	"os/exec"
)

func installGost() {
	// Método 1: Descargar y extraer .tar.gz
	gost_tar_url := "https://github.com/ginuerzh/gost/archive/refs/tags/v2.11.5.tar.gz"
	gost_cmd_tar := exec.Command("curl", "-L", gost_tar_url, "-o", "package.tar.gz")
	err := gost_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gost_zip_url := "https://github.com/ginuerzh/gost/archive/refs/tags/v2.11.5.zip"
	gost_cmd_zip := exec.Command("curl", "-L", gost_zip_url, "-o", "package.zip")
	err = gost_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gost_bin_url := "https://github.com/ginuerzh/gost/archive/refs/tags/v2.11.5.bin"
	gost_cmd_bin := exec.Command("curl", "-L", gost_bin_url, "-o", "binary.bin")
	err = gost_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gost_src_url := "https://github.com/ginuerzh/gost/archive/refs/tags/v2.11.5.src.tar.gz"
	gost_cmd_src := exec.Command("curl", "-L", gost_src_url, "-o", "source.tar.gz")
	err = gost_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gost_cmd_direct := exec.Command("./binary")
	err = gost_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.20")
exec.Command("latte", "install", "go@1.20")
}
