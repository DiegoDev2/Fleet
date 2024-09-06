package main

// Ddh - Fast duplicate file finder
// Homepage: https://github.com/darakian/ddh

import (
	"fmt"
	
	"os/exec"
)

func installDdh() {
	// Método 1: Descargar y extraer .tar.gz
	ddh_tar_url := "https://github.com/darakian/ddh/archive/refs/tags/0.13.0.tar.gz"
	ddh_cmd_tar := exec.Command("curl", "-L", ddh_tar_url, "-o", "package.tar.gz")
	err := ddh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddh_zip_url := "https://github.com/darakian/ddh/archive/refs/tags/0.13.0.zip"
	ddh_cmd_zip := exec.Command("curl", "-L", ddh_zip_url, "-o", "package.zip")
	err = ddh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddh_bin_url := "https://github.com/darakian/ddh/archive/refs/tags/0.13.0.bin"
	ddh_cmd_bin := exec.Command("curl", "-L", ddh_bin_url, "-o", "binary.bin")
	err = ddh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddh_src_url := "https://github.com/darakian/ddh/archive/refs/tags/0.13.0.src.tar.gz"
	ddh_cmd_src := exec.Command("curl", "-L", ddh_src_url, "-o", "source.tar.gz")
	err = ddh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddh_cmd_direct := exec.Command("./binary")
	err = ddh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
