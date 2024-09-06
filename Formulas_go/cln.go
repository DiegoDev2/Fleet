package main

// Cln - Class Library for Numbers
// Homepage: https://www.ginac.de/CLN/

import (
	"fmt"
	
	"os/exec"
)

func installCln() {
	// Método 1: Descargar y extraer .tar.gz
	cln_tar_url := "https://www.ginac.de/CLN/cln-1.3.7.tar.bz2"
	cln_cmd_tar := exec.Command("curl", "-L", cln_tar_url, "-o", "package.tar.gz")
	err := cln_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cln_zip_url := "https://www.ginac.de/CLN/cln-1.3.7.tar.bz2"
	cln_cmd_zip := exec.Command("curl", "-L", cln_zip_url, "-o", "package.zip")
	err = cln_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cln_bin_url := "https://www.ginac.de/CLN/cln-1.3.7.tar.bz2"
	cln_cmd_bin := exec.Command("curl", "-L", cln_bin_url, "-o", "binary.bin")
	err = cln_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cln_src_url := "https://www.ginac.de/CLN/cln-1.3.7.tar.bz2"
	cln_cmd_src := exec.Command("curl", "-L", cln_src_url, "-o", "source.tar.gz")
	err = cln_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cln_cmd_direct := exec.Command("./binary")
	err = cln_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: wget")
exec.Command("latte", "install", "wget")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
