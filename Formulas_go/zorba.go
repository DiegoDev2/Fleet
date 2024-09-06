package main

// Zorba - NoSQL query processor
// Homepage: http://www.zorba.io/

import (
	"fmt"
	
	"os/exec"
)

func installZorba() {
	// Método 1: Descargar y extraer .tar.gz
	zorba_tar_url := "https://github.com/28msec/zorba/archive/refs/tags/3.1.tar.gz"
	zorba_cmd_tar := exec.Command("curl", "-L", zorba_tar_url, "-o", "package.tar.gz")
	err := zorba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zorba_zip_url := "https://github.com/28msec/zorba/archive/refs/tags/3.1.zip"
	zorba_cmd_zip := exec.Command("curl", "-L", zorba_zip_url, "-o", "package.zip")
	err = zorba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zorba_bin_url := "https://github.com/28msec/zorba/archive/refs/tags/3.1.bin"
	zorba_cmd_bin := exec.Command("curl", "-L", zorba_bin_url, "-o", "binary.bin")
	err = zorba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zorba_src_url := "https://github.com/28msec/zorba/archive/refs/tags/3.1.src.tar.gz"
	zorba_cmd_src := exec.Command("curl", "-L", zorba_src_url, "-o", "source.tar.gz")
	err = zorba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zorba_cmd_direct := exec.Command("./binary")
	err = zorba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: flex")
exec.Command("latte", "install", "flex")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
}
