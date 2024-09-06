package main

// Popt - Library like getopt(3) with a number of enhancements
// Homepage: https://github.com/rpm-software-management/popt

import (
	"fmt"
	
	"os/exec"
)

func installPopt() {
	// Método 1: Descargar y extraer .tar.gz
	popt_tar_url := "http://ftp.rpm.org/popt/releases/popt-1.x/popt-1.19.tar.gz"
	popt_cmd_tar := exec.Command("curl", "-L", popt_tar_url, "-o", "package.tar.gz")
	err := popt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	popt_zip_url := "http://ftp.rpm.org/popt/releases/popt-1.x/popt-1.19.zip"
	popt_cmd_zip := exec.Command("curl", "-L", popt_zip_url, "-o", "package.zip")
	err = popt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	popt_bin_url := "http://ftp.rpm.org/popt/releases/popt-1.x/popt-1.19.bin"
	popt_cmd_bin := exec.Command("curl", "-L", popt_bin_url, "-o", "binary.bin")
	err = popt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	popt_src_url := "http://ftp.rpm.org/popt/releases/popt-1.x/popt-1.19.src.tar.gz"
	popt_cmd_src := exec.Command("curl", "-L", popt_src_url, "-o", "source.tar.gz")
	err = popt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	popt_cmd_direct := exec.Command("./binary")
	err = popt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
