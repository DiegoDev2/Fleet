package main

// MicroInetd - Simple network service spawner
// Homepage: https://acme.com/software/micro_inetd/

import (
	"fmt"
	
	"os/exec"
)

func installMicroInetd() {
	// Método 1: Descargar y extraer .tar.gz
	microinetd_tar_url := "https://acme.com/software/micro_inetd/micro_inetd_14Aug2014.tar.gz"
	microinetd_cmd_tar := exec.Command("curl", "-L", microinetd_tar_url, "-o", "package.tar.gz")
	err := microinetd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	microinetd_zip_url := "https://acme.com/software/micro_inetd/micro_inetd_14Aug2014.zip"
	microinetd_cmd_zip := exec.Command("curl", "-L", microinetd_zip_url, "-o", "package.zip")
	err = microinetd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	microinetd_bin_url := "https://acme.com/software/micro_inetd/micro_inetd_14Aug2014.bin"
	microinetd_cmd_bin := exec.Command("curl", "-L", microinetd_bin_url, "-o", "binary.bin")
	err = microinetd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	microinetd_src_url := "https://acme.com/software/micro_inetd/micro_inetd_14Aug2014.src.tar.gz"
	microinetd_cmd_src := exec.Command("curl", "-L", microinetd_src_url, "-o", "source.tar.gz")
	err = microinetd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	microinetd_cmd_direct := exec.Command("./binary")
	err = microinetd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
