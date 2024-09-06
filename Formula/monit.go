package main

// Monit - Manage and monitor processes, files, directories, and devices
// Homepage: https://mmonit.com/monit/

import (
	"fmt"
	
	"os/exec"
)

func installMonit() {
	// Método 1: Descargar y extraer .tar.gz
	monit_tar_url := "https://mmonit.com/monit/dist/monit-5.34.0.tar.gz"
	monit_cmd_tar := exec.Command("curl", "-L", monit_tar_url, "-o", "package.tar.gz")
	err := monit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	monit_zip_url := "https://mmonit.com/monit/dist/monit-5.34.0.zip"
	monit_cmd_zip := exec.Command("curl", "-L", monit_zip_url, "-o", "package.zip")
	err = monit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	monit_bin_url := "https://mmonit.com/monit/dist/monit-5.34.0.bin"
	monit_cmd_bin := exec.Command("curl", "-L", monit_bin_url, "-o", "binary.bin")
	err = monit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	monit_src_url := "https://mmonit.com/monit/dist/monit-5.34.0.src.tar.gz"
	monit_cmd_src := exec.Command("curl", "-L", monit_src_url, "-o", "source.tar.gz")
	err = monit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	monit_cmd_direct := exec.Command("./binary")
	err = monit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
}
