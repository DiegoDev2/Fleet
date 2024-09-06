package main

// Ntl - C++ number theory library
// Homepage: https://libntl.org

import (
	"fmt"
	
	"os/exec"
)

func installNtl() {
	// Método 1: Descargar y extraer .tar.gz
	ntl_tar_url := "https://libntl.org/ntl-11.5.1.tar.gz"
	ntl_cmd_tar := exec.Command("curl", "-L", ntl_tar_url, "-o", "package.tar.gz")
	err := ntl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ntl_zip_url := "https://libntl.org/ntl-11.5.1.zip"
	ntl_cmd_zip := exec.Command("curl", "-L", ntl_zip_url, "-o", "package.zip")
	err = ntl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ntl_bin_url := "https://libntl.org/ntl-11.5.1.bin"
	ntl_cmd_bin := exec.Command("curl", "-L", ntl_bin_url, "-o", "binary.bin")
	err = ntl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ntl_src_url := "https://libntl.org/ntl-11.5.1.src.tar.gz"
	ntl_cmd_src := exec.Command("curl", "-L", ntl_src_url, "-o", "source.tar.gz")
	err = ntl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ntl_cmd_direct := exec.Command("./binary")
	err = ntl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
