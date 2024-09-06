package main

// JsonFortran - Fortran 2008 JSON API
// Homepage: https://github.com/jacobwilliams/json-fortran

import (
	"fmt"
	
	"os/exec"
)

func installJsonFortran() {
	// Método 1: Descargar y extraer .tar.gz
	jsonfortran_tar_url := "https://github.com/jacobwilliams/json-fortran/archive/refs/tags/9.0.2.tar.gz"
	jsonfortran_cmd_tar := exec.Command("curl", "-L", jsonfortran_tar_url, "-o", "package.tar.gz")
	err := jsonfortran_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonfortran_zip_url := "https://github.com/jacobwilliams/json-fortran/archive/refs/tags/9.0.2.zip"
	jsonfortran_cmd_zip := exec.Command("curl", "-L", jsonfortran_zip_url, "-o", "package.zip")
	err = jsonfortran_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonfortran_bin_url := "https://github.com/jacobwilliams/json-fortran/archive/refs/tags/9.0.2.bin"
	jsonfortran_cmd_bin := exec.Command("curl", "-L", jsonfortran_bin_url, "-o", "binary.bin")
	err = jsonfortran_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonfortran_src_url := "https://github.com/jacobwilliams/json-fortran/archive/refs/tags/9.0.2.src.tar.gz"
	jsonfortran_cmd_src := exec.Command("curl", "-L", jsonfortran_src_url, "-o", "source.tar.gz")
	err = jsonfortran_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonfortran_cmd_direct := exec.Command("./binary")
	err = jsonfortran_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ford")
exec.Command("latte", "install", "ford")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
