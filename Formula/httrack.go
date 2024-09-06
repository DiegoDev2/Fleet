package main

// Httrack - Website copier/offline browser
// Homepage: https://www.httrack.com/

import (
	"fmt"
	
	"os/exec"
)

func installHttrack() {
	// Método 1: Descargar y extraer .tar.gz
	httrack_tar_url := "https://mirror.httrack.com/historical/httrack-3.49.2.tar.gz"
	httrack_cmd_tar := exec.Command("curl", "-L", httrack_tar_url, "-o", "package.tar.gz")
	err := httrack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httrack_zip_url := "https://mirror.httrack.com/historical/httrack-3.49.2.zip"
	httrack_cmd_zip := exec.Command("curl", "-L", httrack_zip_url, "-o", "package.zip")
	err = httrack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httrack_bin_url := "https://mirror.httrack.com/historical/httrack-3.49.2.bin"
	httrack_cmd_bin := exec.Command("curl", "-L", httrack_bin_url, "-o", "binary.bin")
	err = httrack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httrack_src_url := "https://mirror.httrack.com/historical/httrack-3.49.2.src.tar.gz"
	httrack_cmd_src := exec.Command("curl", "-L", httrack_src_url, "-o", "source.tar.gz")
	err = httrack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httrack_cmd_direct := exec.Command("./binary")
	err = httrack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
