package main

// Veraxx - Programmable tool for C++ source code
// Homepage: https://bitbucket.org/verateam/vera

import (
	"fmt"
	
	"os/exec"
)

func installVeraxx() {
	// Método 1: Descargar y extraer .tar.gz
	veraxx_tar_url := "https://bitbucket.org/verateam/vera/downloads/vera++-1.3.0.tar.gz"
	veraxx_cmd_tar := exec.Command("curl", "-L", veraxx_tar_url, "-o", "package.tar.gz")
	err := veraxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	veraxx_zip_url := "https://bitbucket.org/verateam/vera/downloads/vera++-1.3.0.zip"
	veraxx_cmd_zip := exec.Command("curl", "-L", veraxx_zip_url, "-o", "package.zip")
	err = veraxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	veraxx_bin_url := "https://bitbucket.org/verateam/vera/downloads/vera++-1.3.0.bin"
	veraxx_cmd_bin := exec.Command("curl", "-L", veraxx_bin_url, "-o", "binary.bin")
	err = veraxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	veraxx_src_url := "https://bitbucket.org/verateam/vera/downloads/vera++-1.3.0.src.tar.gz"
	veraxx_cmd_src := exec.Command("curl", "-L", veraxx_src_url, "-o", "source.tar.gz")
	err = veraxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	veraxx_cmd_direct := exec.Command("./binary")
	err = veraxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
