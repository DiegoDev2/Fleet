package main

// Libdbi - Database-independent abstraction layer in C, similar to DBI/DBD in Perl
// Homepage: https://libdbi.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibdbi() {
	// Método 1: Descargar y extraer .tar.gz
	libdbi_tar_url := "https://downloads.sourceforge.net/project/libdbi/libdbi/libdbi-0.9.0/libdbi-0.9.0.tar.gz"
	libdbi_cmd_tar := exec.Command("curl", "-L", libdbi_tar_url, "-o", "package.tar.gz")
	err := libdbi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdbi_zip_url := "https://downloads.sourceforge.net/project/libdbi/libdbi/libdbi-0.9.0/libdbi-0.9.0.zip"
	libdbi_cmd_zip := exec.Command("curl", "-L", libdbi_zip_url, "-o", "package.zip")
	err = libdbi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdbi_bin_url := "https://downloads.sourceforge.net/project/libdbi/libdbi/libdbi-0.9.0/libdbi-0.9.0.bin"
	libdbi_cmd_bin := exec.Command("curl", "-L", libdbi_bin_url, "-o", "binary.bin")
	err = libdbi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdbi_src_url := "https://downloads.sourceforge.net/project/libdbi/libdbi/libdbi-0.9.0/libdbi-0.9.0.src.tar.gz"
	libdbi_cmd_src := exec.Command("curl", "-L", libdbi_src_url, "-o", "source.tar.gz")
	err = libdbi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdbi_cmd_direct := exec.Command("./binary")
	err = libdbi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
