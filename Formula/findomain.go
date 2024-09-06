package main

// Findomain - Cross-platform subdomain enumerator
// Homepage: https://github.com/Findomain/Findomain

import (
	"fmt"
	
	"os/exec"
)

func installFindomain() {
	// Método 1: Descargar y extraer .tar.gz
	findomain_tar_url := "https://github.com/Findomain/Findomain/archive/refs/tags/9.0.4.tar.gz"
	findomain_cmd_tar := exec.Command("curl", "-L", findomain_tar_url, "-o", "package.tar.gz")
	err := findomain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	findomain_zip_url := "https://github.com/Findomain/Findomain/archive/refs/tags/9.0.4.zip"
	findomain_cmd_zip := exec.Command("curl", "-L", findomain_zip_url, "-o", "package.zip")
	err = findomain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	findomain_bin_url := "https://github.com/Findomain/Findomain/archive/refs/tags/9.0.4.bin"
	findomain_cmd_bin := exec.Command("curl", "-L", findomain_bin_url, "-o", "binary.bin")
	err = findomain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	findomain_src_url := "https://github.com/Findomain/Findomain/archive/refs/tags/9.0.4.src.tar.gz"
	findomain_cmd_src := exec.Command("curl", "-L", findomain_src_url, "-o", "source.tar.gz")
	err = findomain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	findomain_cmd_direct := exec.Command("./binary")
	err = findomain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
