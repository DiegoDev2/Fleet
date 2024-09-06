package main

// Cointop - Interactive terminal based UI application for tracking cryptocurrencies
// Homepage: https://cointop.sh

import (
	"fmt"
	
	"os/exec"
)

func installCointop() {
	// Método 1: Descargar y extraer .tar.gz
	cointop_tar_url := "https://github.com/cointop-sh/cointop/archive/refs/tags/v1.6.10.tar.gz"
	cointop_cmd_tar := exec.Command("curl", "-L", cointop_tar_url, "-o", "package.tar.gz")
	err := cointop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cointop_zip_url := "https://github.com/cointop-sh/cointop/archive/refs/tags/v1.6.10.zip"
	cointop_cmd_zip := exec.Command("curl", "-L", cointop_zip_url, "-o", "package.zip")
	err = cointop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cointop_bin_url := "https://github.com/cointop-sh/cointop/archive/refs/tags/v1.6.10.bin"
	cointop_cmd_bin := exec.Command("curl", "-L", cointop_bin_url, "-o", "binary.bin")
	err = cointop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cointop_src_url := "https://github.com/cointop-sh/cointop/archive/refs/tags/v1.6.10.src.tar.gz"
	cointop_cmd_src := exec.Command("curl", "-L", cointop_src_url, "-o", "source.tar.gz")
	err = cointop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cointop_cmd_direct := exec.Command("./binary")
	err = cointop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
