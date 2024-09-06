package main

// Slowhttptest - Simulates application layer denial of service attacks
// Homepage: https://github.com/shekyan/slowhttptest

import (
	"fmt"
	
	"os/exec"
)

func installSlowhttptest() {
	// Método 1: Descargar y extraer .tar.gz
	slowhttptest_tar_url := "https://github.com/shekyan/slowhttptest/archive/refs/tags/v1.9.0.tar.gz"
	slowhttptest_cmd_tar := exec.Command("curl", "-L", slowhttptest_tar_url, "-o", "package.tar.gz")
	err := slowhttptest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slowhttptest_zip_url := "https://github.com/shekyan/slowhttptest/archive/refs/tags/v1.9.0.zip"
	slowhttptest_cmd_zip := exec.Command("curl", "-L", slowhttptest_zip_url, "-o", "package.zip")
	err = slowhttptest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slowhttptest_bin_url := "https://github.com/shekyan/slowhttptest/archive/refs/tags/v1.9.0.bin"
	slowhttptest_cmd_bin := exec.Command("curl", "-L", slowhttptest_bin_url, "-o", "binary.bin")
	err = slowhttptest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slowhttptest_src_url := "https://github.com/shekyan/slowhttptest/archive/refs/tags/v1.9.0.src.tar.gz"
	slowhttptest_cmd_src := exec.Command("curl", "-L", slowhttptest_src_url, "-o", "source.tar.gz")
	err = slowhttptest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slowhttptest_cmd_direct := exec.Command("./binary")
	err = slowhttptest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
