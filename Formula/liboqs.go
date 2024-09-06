package main

// Liboqs - Library for quantum-safe cryptography
// Homepage: https://openquantumsafe.org/

import (
	"fmt"
	
	"os/exec"
)

func installLiboqs() {
	// Método 1: Descargar y extraer .tar.gz
	liboqs_tar_url := "https://github.com/open-quantum-safe/liboqs/archive/refs/tags/0.10.1.tar.gz"
	liboqs_cmd_tar := exec.Command("curl", "-L", liboqs_tar_url, "-o", "package.tar.gz")
	err := liboqs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liboqs_zip_url := "https://github.com/open-quantum-safe/liboqs/archive/refs/tags/0.10.1.zip"
	liboqs_cmd_zip := exec.Command("curl", "-L", liboqs_zip_url, "-o", "package.zip")
	err = liboqs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liboqs_bin_url := "https://github.com/open-quantum-safe/liboqs/archive/refs/tags/0.10.1.bin"
	liboqs_cmd_bin := exec.Command("curl", "-L", liboqs_bin_url, "-o", "binary.bin")
	err = liboqs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liboqs_src_url := "https://github.com/open-quantum-safe/liboqs/archive/refs/tags/0.10.1.src.tar.gz"
	liboqs_cmd_src := exec.Command("curl", "-L", liboqs_src_url, "-o", "source.tar.gz")
	err = liboqs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liboqs_cmd_direct := exec.Command("./binary")
	err = liboqs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
