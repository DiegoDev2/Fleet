package main

// Alp - Access Log Profiler
// Homepage: https://github.com/tkuchiki/alp

import (
	"fmt"
	
	"os/exec"
)

func installAlp() {
	// Método 1: Descargar y extraer .tar.gz
	alp_tar_url := "https://github.com/tkuchiki/alp/archive/refs/tags/v1.0.21.tar.gz"
	alp_cmd_tar := exec.Command("curl", "-L", alp_tar_url, "-o", "package.tar.gz")
	err := alp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alp_zip_url := "https://github.com/tkuchiki/alp/archive/refs/tags/v1.0.21.zip"
	alp_cmd_zip := exec.Command("curl", "-L", alp_zip_url, "-o", "package.zip")
	err = alp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alp_bin_url := "https://github.com/tkuchiki/alp/archive/refs/tags/v1.0.21.bin"
	alp_cmd_bin := exec.Command("curl", "-L", alp_bin_url, "-o", "binary.bin")
	err = alp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alp_src_url := "https://github.com/tkuchiki/alp/archive/refs/tags/v1.0.21.src.tar.gz"
	alp_cmd_src := exec.Command("curl", "-L", alp_src_url, "-o", "source.tar.gz")
	err = alp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alp_cmd_direct := exec.Command("./binary")
	err = alp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
