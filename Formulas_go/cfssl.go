package main

// Cfssl - CloudFlare's PKI toolkit
// Homepage: https://cfssl.org/

import (
	"fmt"
	
	"os/exec"
)

func installCfssl() {
	// Método 1: Descargar y extraer .tar.gz
	cfssl_tar_url := "https://github.com/cloudflare/cfssl/archive/refs/tags/v1.6.5.tar.gz"
	cfssl_cmd_tar := exec.Command("curl", "-L", cfssl_tar_url, "-o", "package.tar.gz")
	err := cfssl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfssl_zip_url := "https://github.com/cloudflare/cfssl/archive/refs/tags/v1.6.5.zip"
	cfssl_cmd_zip := exec.Command("curl", "-L", cfssl_zip_url, "-o", "package.zip")
	err = cfssl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfssl_bin_url := "https://github.com/cloudflare/cfssl/archive/refs/tags/v1.6.5.bin"
	cfssl_cmd_bin := exec.Command("curl", "-L", cfssl_bin_url, "-o", "binary.bin")
	err = cfssl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfssl_src_url := "https://github.com/cloudflare/cfssl/archive/refs/tags/v1.6.5.src.tar.gz"
	cfssl_cmd_src := exec.Command("curl", "-L", cfssl_src_url, "-o", "source.tar.gz")
	err = cfssl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfssl_cmd_direct := exec.Command("./binary")
	err = cfssl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
