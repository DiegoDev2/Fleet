package main

// Prestd - Simplify and accelerate development on any Postgres application, existing or new
// Homepage: https://github.com/prest/prest

import (
	"fmt"
	
	"os/exec"
)

func installPrestd() {
	// Método 1: Descargar y extraer .tar.gz
	prestd_tar_url := "https://github.com/prest/prest/archive/refs/tags/v1.5.5.tar.gz"
	prestd_cmd_tar := exec.Command("curl", "-L", prestd_tar_url, "-o", "package.tar.gz")
	err := prestd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prestd_zip_url := "https://github.com/prest/prest/archive/refs/tags/v1.5.5.zip"
	prestd_cmd_zip := exec.Command("curl", "-L", prestd_zip_url, "-o", "package.zip")
	err = prestd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prestd_bin_url := "https://github.com/prest/prest/archive/refs/tags/v1.5.5.bin"
	prestd_cmd_bin := exec.Command("curl", "-L", prestd_bin_url, "-o", "binary.bin")
	err = prestd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prestd_src_url := "https://github.com/prest/prest/archive/refs/tags/v1.5.5.src.tar.gz"
	prestd_cmd_src := exec.Command("curl", "-L", prestd_src_url, "-o", "source.tar.gz")
	err = prestd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prestd_cmd_direct := exec.Command("./binary")
	err = prestd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
