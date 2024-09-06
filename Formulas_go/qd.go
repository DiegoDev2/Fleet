package main

// Qd - C++/Fortran-90 double-double and quad-double package
// Homepage: https://www.davidhbailey.com/dhbsoftware/

import (
	"fmt"
	
	"os/exec"
)

func installQd() {
	// Método 1: Descargar y extraer .tar.gz
	qd_tar_url := "https://www.davidhbailey.com/dhbsoftware/qd-2.3.24.tar.gz"
	qd_cmd_tar := exec.Command("curl", "-L", qd_tar_url, "-o", "package.tar.gz")
	err := qd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qd_zip_url := "https://www.davidhbailey.com/dhbsoftware/qd-2.3.24.zip"
	qd_cmd_zip := exec.Command("curl", "-L", qd_zip_url, "-o", "package.zip")
	err = qd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qd_bin_url := "https://www.davidhbailey.com/dhbsoftware/qd-2.3.24.bin"
	qd_cmd_bin := exec.Command("curl", "-L", qd_bin_url, "-o", "binary.bin")
	err = qd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qd_src_url := "https://www.davidhbailey.com/dhbsoftware/qd-2.3.24.src.tar.gz"
	qd_cmd_src := exec.Command("curl", "-L", qd_src_url, "-o", "source.tar.gz")
	err = qd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qd_cmd_direct := exec.Command("./binary")
	err = qd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
