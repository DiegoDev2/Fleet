package main

// Hq - Jq, but for HTML
// Homepage: https://github.com/orf/html-query

import (
	"fmt"
	
	"os/exec"
)

func installHq() {
	// Método 1: Descargar y extraer .tar.gz
	hq_tar_url := "https://github.com/orf/html-query/archive/refs/tags/html-query-v1.2.2.tar.gz"
	hq_cmd_tar := exec.Command("curl", "-L", hq_tar_url, "-o", "package.tar.gz")
	err := hq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hq_zip_url := "https://github.com/orf/html-query/archive/refs/tags/html-query-v1.2.2.zip"
	hq_cmd_zip := exec.Command("curl", "-L", hq_zip_url, "-o", "package.zip")
	err = hq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hq_bin_url := "https://github.com/orf/html-query/archive/refs/tags/html-query-v1.2.2.bin"
	hq_cmd_bin := exec.Command("curl", "-L", hq_bin_url, "-o", "binary.bin")
	err = hq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hq_src_url := "https://github.com/orf/html-query/archive/refs/tags/html-query-v1.2.2.src.tar.gz"
	hq_cmd_src := exec.Command("curl", "-L", hq_src_url, "-o", "source.tar.gz")
	err = hq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hq_cmd_direct := exec.Command("./binary")
	err = hq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
