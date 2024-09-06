package main

// Ftnchek - Fortran 77 program checker
// Homepage: https://www.dsm.fordham.edu/~ftnchek/

import (
	"fmt"
	
	"os/exec"
)

func installFtnchek() {
	// Método 1: Descargar y extraer .tar.gz
	ftnchek_tar_url := "https://www.dsm.fordham.edu/~ftnchek/download/ftnchek-3.3.1.tar.gz"
	ftnchek_cmd_tar := exec.Command("curl", "-L", ftnchek_tar_url, "-o", "package.tar.gz")
	err := ftnchek_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ftnchek_zip_url := "https://www.dsm.fordham.edu/~ftnchek/download/ftnchek-3.3.1.zip"
	ftnchek_cmd_zip := exec.Command("curl", "-L", ftnchek_zip_url, "-o", "package.zip")
	err = ftnchek_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ftnchek_bin_url := "https://www.dsm.fordham.edu/~ftnchek/download/ftnchek-3.3.1.bin"
	ftnchek_cmd_bin := exec.Command("curl", "-L", ftnchek_bin_url, "-o", "binary.bin")
	err = ftnchek_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ftnchek_src_url := "https://www.dsm.fordham.edu/~ftnchek/download/ftnchek-3.3.1.src.tar.gz"
	ftnchek_cmd_src := exec.Command("curl", "-L", ftnchek_src_url, "-o", "source.tar.gz")
	err = ftnchek_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ftnchek_cmd_direct := exec.Command("./binary")
	err = ftnchek_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: groff")
exec.Command("latte", "install", "groff")
}
