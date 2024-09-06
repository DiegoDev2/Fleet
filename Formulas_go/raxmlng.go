package main

// RaxmlNg - RAxML Next Generation: faster, easier-to-use and more flexible
// Homepage: https://cme.h-its.org/exelixis/web/software/raxml/

import (
	"fmt"
	
	"os/exec"
)

func installRaxmlNg() {
	// Método 1: Descargar y extraer .tar.gz
	raxmlng_tar_url := "https://github.com/amkozlov/raxml-ng.git"
	raxmlng_cmd_tar := exec.Command("curl", "-L", raxmlng_tar_url, "-o", "package.tar.gz")
	err := raxmlng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	raxmlng_zip_url := "https://github.com/amkozlov/raxml-ng.git"
	raxmlng_cmd_zip := exec.Command("curl", "-L", raxmlng_zip_url, "-o", "package.zip")
	err = raxmlng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	raxmlng_bin_url := "https://github.com/amkozlov/raxml-ng.git"
	raxmlng_cmd_bin := exec.Command("curl", "-L", raxmlng_bin_url, "-o", "binary.bin")
	err = raxmlng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	raxmlng_src_url := "https://github.com/amkozlov/raxml-ng.git"
	raxmlng_cmd_src := exec.Command("curl", "-L", raxmlng_src_url, "-o", "source.tar.gz")
	err = raxmlng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	raxmlng_cmd_direct := exec.Command("./binary")
	err = raxmlng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
}
