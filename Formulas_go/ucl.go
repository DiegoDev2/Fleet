package main

// Ucl - Data compression library with small memory footprint
// Homepage: https://www.oberhumer.com/opensource/ucl/

import (
	"fmt"
	
	"os/exec"
)

func installUcl() {
	// Método 1: Descargar y extraer .tar.gz
	ucl_tar_url := "https://www.oberhumer.com/opensource/ucl/download/ucl-1.03.tar.gz"
	ucl_cmd_tar := exec.Command("curl", "-L", ucl_tar_url, "-o", "package.tar.gz")
	err := ucl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ucl_zip_url := "https://www.oberhumer.com/opensource/ucl/download/ucl-1.03.zip"
	ucl_cmd_zip := exec.Command("curl", "-L", ucl_zip_url, "-o", "package.zip")
	err = ucl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ucl_bin_url := "https://www.oberhumer.com/opensource/ucl/download/ucl-1.03.bin"
	ucl_cmd_bin := exec.Command("curl", "-L", ucl_bin_url, "-o", "binary.bin")
	err = ucl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ucl_src_url := "https://www.oberhumer.com/opensource/ucl/download/ucl-1.03.src.tar.gz"
	ucl_cmd_src := exec.Command("curl", "-L", ucl_src_url, "-o", "source.tar.gz")
	err = ucl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ucl_cmd_direct := exec.Command("./binary")
	err = ucl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
