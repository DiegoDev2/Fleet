package main

// Testssl - Tool which checks for the support of TLS/SSL ciphers and flaws
// Homepage: https://testssl.sh/

import (
	"fmt"
	
	"os/exec"
)

func installTestssl() {
	// Método 1: Descargar y extraer .tar.gz
	testssl_tar_url := "https://github.com/drwetter/testssl.sh/archive/refs/tags/v3.0.9.tar.gz"
	testssl_cmd_tar := exec.Command("curl", "-L", testssl_tar_url, "-o", "package.tar.gz")
	err := testssl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	testssl_zip_url := "https://github.com/drwetter/testssl.sh/archive/refs/tags/v3.0.9.zip"
	testssl_cmd_zip := exec.Command("curl", "-L", testssl_zip_url, "-o", "package.zip")
	err = testssl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	testssl_bin_url := "https://github.com/drwetter/testssl.sh/archive/refs/tags/v3.0.9.bin"
	testssl_cmd_bin := exec.Command("curl", "-L", testssl_bin_url, "-o", "binary.bin")
	err = testssl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	testssl_src_url := "https://github.com/drwetter/testssl.sh/archive/refs/tags/v3.0.9.src.tar.gz"
	testssl_cmd_src := exec.Command("curl", "-L", testssl_src_url, "-o", "source.tar.gz")
	err = testssl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	testssl_cmd_direct := exec.Command("./binary")
	err = testssl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: bind")
	exec.Command("latte", "install", "bind").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
