package main

// Ratfor - Rational Fortran
// Homepage: http://www.dgate.org/ratfor/

import (
	"fmt"
	
	"os/exec"
)

func installRatfor() {
	// Método 1: Descargar y extraer .tar.gz
	ratfor_tar_url := "http://www.dgate.org/ratfor/tars/ratfor-1.05.tar.gz"
	ratfor_cmd_tar := exec.Command("curl", "-L", ratfor_tar_url, "-o", "package.tar.gz")
	err := ratfor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ratfor_zip_url := "http://www.dgate.org/ratfor/tars/ratfor-1.05.zip"
	ratfor_cmd_zip := exec.Command("curl", "-L", ratfor_zip_url, "-o", "package.zip")
	err = ratfor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ratfor_bin_url := "http://www.dgate.org/ratfor/tars/ratfor-1.05.bin"
	ratfor_cmd_bin := exec.Command("curl", "-L", ratfor_bin_url, "-o", "binary.bin")
	err = ratfor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ratfor_src_url := "http://www.dgate.org/ratfor/tars/ratfor-1.05.src.tar.gz"
	ratfor_cmd_src := exec.Command("curl", "-L", ratfor_src_url, "-o", "source.tar.gz")
	err = ratfor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ratfor_cmd_direct := exec.Command("./binary")
	err = ratfor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
