package main

// Foma - Finite-state compiler and C library
// Homepage: https://code.google.com/p/foma/

import (
	"fmt"
	
	"os/exec"
)

func installFoma() {
	// Método 1: Descargar y extraer .tar.gz
	foma_tar_url := "https://bitbucket.org/mhulden/foma/downloads/foma-0.9.18.tar.gz"
	foma_cmd_tar := exec.Command("curl", "-L", foma_tar_url, "-o", "package.tar.gz")
	err := foma_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	foma_zip_url := "https://bitbucket.org/mhulden/foma/downloads/foma-0.9.18.zip"
	foma_cmd_zip := exec.Command("curl", "-L", foma_zip_url, "-o", "package.zip")
	err = foma_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	foma_bin_url := "https://bitbucket.org/mhulden/foma/downloads/foma-0.9.18.bin"
	foma_cmd_bin := exec.Command("curl", "-L", foma_bin_url, "-o", "binary.bin")
	err = foma_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	foma_src_url := "https://bitbucket.org/mhulden/foma/downloads/foma-0.9.18.src.tar.gz"
	foma_cmd_src := exec.Command("curl", "-L", foma_src_url, "-o", "source.tar.gz")
	err = foma_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	foma_cmd_direct := exec.Command("./binary")
	err = foma_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
