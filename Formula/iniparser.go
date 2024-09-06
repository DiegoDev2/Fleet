package main

// Iniparser - Library for parsing ini files
// Homepage: https://gitlab.com/iniparser/iniparser

import (
	"fmt"
	
	"os/exec"
)

func installIniparser() {
	// Método 1: Descargar y extraer .tar.gz
	iniparser_tar_url := "https://gitlab.com/iniparser/iniparser/-/archive/v4.2.4/iniparser-v4.2.4.tar.bz2"
	iniparser_cmd_tar := exec.Command("curl", "-L", iniparser_tar_url, "-o", "package.tar.gz")
	err := iniparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iniparser_zip_url := "https://gitlab.com/iniparser/iniparser/-/archive/v4.2.4/iniparser-v4.2.4.tar.bz2"
	iniparser_cmd_zip := exec.Command("curl", "-L", iniparser_zip_url, "-o", "package.zip")
	err = iniparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iniparser_bin_url := "https://gitlab.com/iniparser/iniparser/-/archive/v4.2.4/iniparser-v4.2.4.tar.bz2"
	iniparser_cmd_bin := exec.Command("curl", "-L", iniparser_bin_url, "-o", "binary.bin")
	err = iniparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iniparser_src_url := "https://gitlab.com/iniparser/iniparser/-/archive/v4.2.4/iniparser-v4.2.4.tar.bz2"
	iniparser_cmd_src := exec.Command("curl", "-L", iniparser_src_url, "-o", "source.tar.gz")
	err = iniparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iniparser_cmd_direct := exec.Command("./binary")
	err = iniparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
}
