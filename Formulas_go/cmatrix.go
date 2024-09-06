package main

// Cmatrix - Console Matrix
// Homepage: https://github.com/abishekvashok/cmatrix/

import (
	"fmt"
	
	"os/exec"
)

func installCmatrix() {
	// Método 1: Descargar y extraer .tar.gz
	cmatrix_tar_url := "https://github.com/abishekvashok/cmatrix/archive/refs/tags/v2.0.tar.gz"
	cmatrix_cmd_tar := exec.Command("curl", "-L", cmatrix_tar_url, "-o", "package.tar.gz")
	err := cmatrix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmatrix_zip_url := "https://github.com/abishekvashok/cmatrix/archive/refs/tags/v2.0.zip"
	cmatrix_cmd_zip := exec.Command("curl", "-L", cmatrix_zip_url, "-o", "package.zip")
	err = cmatrix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmatrix_bin_url := "https://github.com/abishekvashok/cmatrix/archive/refs/tags/v2.0.bin"
	cmatrix_cmd_bin := exec.Command("curl", "-L", cmatrix_bin_url, "-o", "binary.bin")
	err = cmatrix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmatrix_src_url := "https://github.com/abishekvashok/cmatrix/archive/refs/tags/v2.0.src.tar.gz"
	cmatrix_cmd_src := exec.Command("curl", "-L", cmatrix_src_url, "-o", "source.tar.gz")
	err = cmatrix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmatrix_cmd_direct := exec.Command("./binary")
	err = cmatrix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
