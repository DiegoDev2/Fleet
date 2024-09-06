package main

// S2n - Implementation of the TLS/SSL protocols
// Homepage: https://github.com/aws/s2n-tls

import (
	"fmt"
	
	"os/exec"
)

func installS2n() {
	// Método 1: Descargar y extraer .tar.gz
	s2n_tar_url := "https://github.com/aws/s2n-tls/archive/refs/tags/v1.5.1.tar.gz"
	s2n_cmd_tar := exec.Command("curl", "-L", s2n_tar_url, "-o", "package.tar.gz")
	err := s2n_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s2n_zip_url := "https://github.com/aws/s2n-tls/archive/refs/tags/v1.5.1.zip"
	s2n_cmd_zip := exec.Command("curl", "-L", s2n_zip_url, "-o", "package.zip")
	err = s2n_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s2n_bin_url := "https://github.com/aws/s2n-tls/archive/refs/tags/v1.5.1.bin"
	s2n_cmd_bin := exec.Command("curl", "-L", s2n_bin_url, "-o", "binary.bin")
	err = s2n_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s2n_src_url := "https://github.com/aws/s2n-tls/archive/refs/tags/v1.5.1.src.tar.gz"
	s2n_cmd_src := exec.Command("curl", "-L", s2n_src_url, "-o", "source.tar.gz")
	err = s2n_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s2n_cmd_direct := exec.Command("./binary")
	err = s2n_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
