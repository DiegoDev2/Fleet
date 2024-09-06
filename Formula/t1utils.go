package main

// T1utils - Command-line tools for dealing with Type 1 fonts
// Homepage: https://www.lcdf.org/type/

import (
	"fmt"
	
	"os/exec"
)

func installT1utils() {
	// Método 1: Descargar y extraer .tar.gz
	t1utils_tar_url := "https://www.lcdf.org/type/t1utils-1.42.tar.gz"
	t1utils_cmd_tar := exec.Command("curl", "-L", t1utils_tar_url, "-o", "package.tar.gz")
	err := t1utils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	t1utils_zip_url := "https://www.lcdf.org/type/t1utils-1.42.zip"
	t1utils_cmd_zip := exec.Command("curl", "-L", t1utils_zip_url, "-o", "package.zip")
	err = t1utils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	t1utils_bin_url := "https://www.lcdf.org/type/t1utils-1.42.bin"
	t1utils_cmd_bin := exec.Command("curl", "-L", t1utils_bin_url, "-o", "binary.bin")
	err = t1utils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	t1utils_src_url := "https://www.lcdf.org/type/t1utils-1.42.src.tar.gz"
	t1utils_cmd_src := exec.Command("curl", "-L", t1utils_src_url, "-o", "source.tar.gz")
	err = t1utils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	t1utils_cmd_direct := exec.Command("./binary")
	err = t1utils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
