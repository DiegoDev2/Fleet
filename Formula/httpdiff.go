package main

// Httpdiff - Compare two HTTP(S) responses
// Homepage: https://github.com/jgrahamc/httpdiff

import (
	"fmt"
	
	"os/exec"
)

func installHttpdiff() {
	// Método 1: Descargar y extraer .tar.gz
	httpdiff_tar_url := "https://github.com/jgrahamc/httpdiff/archive/refs/tags/v1.0.0.tar.gz"
	httpdiff_cmd_tar := exec.Command("curl", "-L", httpdiff_tar_url, "-o", "package.tar.gz")
	err := httpdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpdiff_zip_url := "https://github.com/jgrahamc/httpdiff/archive/refs/tags/v1.0.0.zip"
	httpdiff_cmd_zip := exec.Command("curl", "-L", httpdiff_zip_url, "-o", "package.zip")
	err = httpdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpdiff_bin_url := "https://github.com/jgrahamc/httpdiff/archive/refs/tags/v1.0.0.bin"
	httpdiff_cmd_bin := exec.Command("curl", "-L", httpdiff_bin_url, "-o", "binary.bin")
	err = httpdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpdiff_src_url := "https://github.com/jgrahamc/httpdiff/archive/refs/tags/v1.0.0.src.tar.gz"
	httpdiff_cmd_src := exec.Command("curl", "-L", httpdiff_src_url, "-o", "source.tar.gz")
	err = httpdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpdiff_cmd_direct := exec.Command("./binary")
	err = httpdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
