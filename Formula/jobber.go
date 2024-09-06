package main

// Jobber - Alternative to cron, with better status-reporting and error-handling
// Homepage: https://dshearer.github.io/jobber/

import (
	"fmt"
	
	"os/exec"
)

func installJobber() {
	// Método 1: Descargar y extraer .tar.gz
	jobber_tar_url := "https://github.com/dshearer/jobber/archive/refs/tags/v1.4.4.tar.gz"
	jobber_cmd_tar := exec.Command("curl", "-L", jobber_tar_url, "-o", "package.tar.gz")
	err := jobber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jobber_zip_url := "https://github.com/dshearer/jobber/archive/refs/tags/v1.4.4.zip"
	jobber_cmd_zip := exec.Command("curl", "-L", jobber_zip_url, "-o", "package.zip")
	err = jobber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jobber_bin_url := "https://github.com/dshearer/jobber/archive/refs/tags/v1.4.4.bin"
	jobber_cmd_bin := exec.Command("curl", "-L", jobber_bin_url, "-o", "binary.bin")
	err = jobber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jobber_src_url := "https://github.com/dshearer/jobber/archive/refs/tags/v1.4.4.src.tar.gz"
	jobber_cmd_src := exec.Command("curl", "-L", jobber_src_url, "-o", "source.tar.gz")
	err = jobber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jobber_cmd_direct := exec.Command("./binary")
	err = jobber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
