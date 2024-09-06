package main

// Pkcs11Tools - Tools to manage objects on PKCS#11 crypotographic tokens
// Homepage: https://github.com/Mastercard/pkcs11-tools

import (
	"fmt"
	
	"os/exec"
)

func installPkcs11Tools() {
	// Método 1: Descargar y extraer .tar.gz
	pkcs11tools_tar_url := "https://github.com/Mastercard/pkcs11-tools/releases/download/v2.6.0/pkcs11-tools-2.6.0.tar.gz"
	pkcs11tools_cmd_tar := exec.Command("curl", "-L", pkcs11tools_tar_url, "-o", "package.tar.gz")
	err := pkcs11tools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pkcs11tools_zip_url := "https://github.com/Mastercard/pkcs11-tools/releases/download/v2.6.0/pkcs11-tools-2.6.0.zip"
	pkcs11tools_cmd_zip := exec.Command("curl", "-L", pkcs11tools_zip_url, "-o", "package.zip")
	err = pkcs11tools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pkcs11tools_bin_url := "https://github.com/Mastercard/pkcs11-tools/releases/download/v2.6.0/pkcs11-tools-2.6.0.bin"
	pkcs11tools_cmd_bin := exec.Command("curl", "-L", pkcs11tools_bin_url, "-o", "binary.bin")
	err = pkcs11tools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pkcs11tools_src_url := "https://github.com/Mastercard/pkcs11-tools/releases/download/v2.6.0/pkcs11-tools-2.6.0.src.tar.gz"
	pkcs11tools_cmd_src := exec.Command("curl", "-L", pkcs11tools_src_url, "-o", "source.tar.gz")
	err = pkcs11tools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pkcs11tools_cmd_direct := exec.Command("./binary")
	err = pkcs11tools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: softhsm")
exec.Command("latte", "install", "softhsm")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
