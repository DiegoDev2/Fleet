package main

// Apib - HTTP performance-testing tool
// Homepage: https://github.com/apigee/apib

import (
	"fmt"
	
	"os/exec"
)

func installApib() {
	// Método 1: Descargar y extraer .tar.gz
	apib_tar_url := "https://github.com/apigee/apib/archive/refs/tags/APIB_1_2_1.tar.gz"
	apib_cmd_tar := exec.Command("curl", "-L", apib_tar_url, "-o", "package.tar.gz")
	err := apib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apib_zip_url := "https://github.com/apigee/apib/archive/refs/tags/APIB_1_2_1.zip"
	apib_cmd_zip := exec.Command("curl", "-L", apib_zip_url, "-o", "package.zip")
	err = apib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apib_bin_url := "https://github.com/apigee/apib/archive/refs/tags/APIB_1_2_1.bin"
	apib_cmd_bin := exec.Command("curl", "-L", apib_bin_url, "-o", "binary.bin")
	err = apib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apib_src_url := "https://github.com/apigee/apib/archive/refs/tags/APIB_1_2_1.src.tar.gz"
	apib_cmd_src := exec.Command("curl", "-L", apib_src_url, "-o", "source.tar.gz")
	err = apib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apib_cmd_direct := exec.Command("./binary")
	err = apib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
