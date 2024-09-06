package main

// Liborigin - Library for reading OriginLab OPJ project files
// Homepage: https://sourceforge.net/projects/liborigin/

import (
	"fmt"
	
	"os/exec"
)

func installLiborigin() {
	// Método 1: Descargar y extraer .tar.gz
	liborigin_tar_url := "https://downloads.sourceforge.net/project/liborigin/liborigin/3.0/liborigin-3.0.3.tar.gz"
	liborigin_cmd_tar := exec.Command("curl", "-L", liborigin_tar_url, "-o", "package.tar.gz")
	err := liborigin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liborigin_zip_url := "https://downloads.sourceforge.net/project/liborigin/liborigin/3.0/liborigin-3.0.3.zip"
	liborigin_cmd_zip := exec.Command("curl", "-L", liborigin_zip_url, "-o", "package.zip")
	err = liborigin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liborigin_bin_url := "https://downloads.sourceforge.net/project/liborigin/liborigin/3.0/liborigin-3.0.3.bin"
	liborigin_cmd_bin := exec.Command("curl", "-L", liborigin_bin_url, "-o", "binary.bin")
	err = liborigin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liborigin_src_url := "https://downloads.sourceforge.net/project/liborigin/liborigin/3.0/liborigin-3.0.3.src.tar.gz"
	liborigin_cmd_src := exec.Command("curl", "-L", liborigin_src_url, "-o", "source.tar.gz")
	err = liborigin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liborigin_cmd_direct := exec.Command("./binary")
	err = liborigin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
