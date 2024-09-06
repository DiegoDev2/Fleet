package main

// Ulfius - HTTP Framework for REST Applications in C
// Homepage: https://github.com/babelouest/ulfius/

import (
	"fmt"
	
	"os/exec"
)

func installUlfius() {
	// Método 1: Descargar y extraer .tar.gz
	ulfius_tar_url := "https://github.com/babelouest/ulfius/archive/refs/tags/v2.7.15.tar.gz"
	ulfius_cmd_tar := exec.Command("curl", "-L", ulfius_tar_url, "-o", "package.tar.gz")
	err := ulfius_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ulfius_zip_url := "https://github.com/babelouest/ulfius/archive/refs/tags/v2.7.15.zip"
	ulfius_cmd_zip := exec.Command("curl", "-L", ulfius_zip_url, "-o", "package.zip")
	err = ulfius_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ulfius_bin_url := "https://github.com/babelouest/ulfius/archive/refs/tags/v2.7.15.bin"
	ulfius_cmd_bin := exec.Command("curl", "-L", ulfius_bin_url, "-o", "binary.bin")
	err = ulfius_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ulfius_src_url := "https://github.com/babelouest/ulfius/archive/refs/tags/v2.7.15.src.tar.gz"
	ulfius_cmd_src := exec.Command("curl", "-L", ulfius_src_url, "-o", "source.tar.gz")
	err = ulfius_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ulfius_cmd_direct := exec.Command("./binary")
	err = ulfius_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: libmicrohttpd")
exec.Command("latte", "install", "libmicrohttpd")
	fmt.Println("Instalando dependencia: orcania")
exec.Command("latte", "install", "orcania")
	fmt.Println("Instalando dependencia: yder")
exec.Command("latte", "install", "yder")
}
