package main

// Squiid - Do advanced algebraic and RPN calculations
// Homepage: https://imaginaryinfinity.net/projects/squiid/

import (
	"fmt"
	
	"os/exec"
)

func installSquiid() {
	// Método 1: Descargar y extraer .tar.gz
	squiid_tar_url := "https://gitlab.com/ImaginaryInfinity/squiid-calculator/squiid/-/archive/1.1.2/squiid-1.1.2.tar.bz2"
	squiid_cmd_tar := exec.Command("curl", "-L", squiid_tar_url, "-o", "package.tar.gz")
	err := squiid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	squiid_zip_url := "https://gitlab.com/ImaginaryInfinity/squiid-calculator/squiid/-/archive/1.1.2/squiid-1.1.2.tar.bz2"
	squiid_cmd_zip := exec.Command("curl", "-L", squiid_zip_url, "-o", "package.zip")
	err = squiid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	squiid_bin_url := "https://gitlab.com/ImaginaryInfinity/squiid-calculator/squiid/-/archive/1.1.2/squiid-1.1.2.tar.bz2"
	squiid_cmd_bin := exec.Command("curl", "-L", squiid_bin_url, "-o", "binary.bin")
	err = squiid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	squiid_src_url := "https://gitlab.com/ImaginaryInfinity/squiid-calculator/squiid/-/archive/1.1.2/squiid-1.1.2.tar.bz2"
	squiid_cmd_src := exec.Command("curl", "-L", squiid_src_url, "-o", "source.tar.gz")
	err = squiid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	squiid_cmd_direct := exec.Command("./binary")
	err = squiid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: nng")
exec.Command("latte", "install", "nng")
}
