package main

// Wimlib - Library to create, extract, and modify Windows Imaging files
// Homepage: https://wimlib.net/

import (
	"fmt"
	
	"os/exec"
)

func installWimlib() {
	// Método 1: Descargar y extraer .tar.gz
	wimlib_tar_url := "https://wimlib.net/downloads/wimlib-1.14.4.tar.gz"
	wimlib_cmd_tar := exec.Command("curl", "-L", wimlib_tar_url, "-o", "package.tar.gz")
	err := wimlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wimlib_zip_url := "https://wimlib.net/downloads/wimlib-1.14.4.zip"
	wimlib_cmd_zip := exec.Command("curl", "-L", wimlib_zip_url, "-o", "package.zip")
	err = wimlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wimlib_bin_url := "https://wimlib.net/downloads/wimlib-1.14.4.bin"
	wimlib_cmd_bin := exec.Command("curl", "-L", wimlib_bin_url, "-o", "binary.bin")
	err = wimlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wimlib_src_url := "https://wimlib.net/downloads/wimlib-1.14.4.src.tar.gz"
	wimlib_cmd_src := exec.Command("curl", "-L", wimlib_src_url, "-o", "source.tar.gz")
	err = wimlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wimlib_cmd_direct := exec.Command("./binary")
	err = wimlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
