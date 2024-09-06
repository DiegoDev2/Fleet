package main

// Pkcs11Helper - Library to simplify the interaction with PKCS#11
// Homepage: https://github.com/OpenSC/OpenSC/wiki/pkcs11-helper

import (
	"fmt"
	
	"os/exec"
)

func installPkcs11Helper() {
	// Método 1: Descargar y extraer .tar.gz
	pkcs11helper_tar_url := "https://github.com/OpenSC/pkcs11-helper/releases/download/pkcs11-helper-1.30.0/pkcs11-helper-1.30.0.tar.bz2"
	pkcs11helper_cmd_tar := exec.Command("curl", "-L", pkcs11helper_tar_url, "-o", "package.tar.gz")
	err := pkcs11helper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pkcs11helper_zip_url := "https://github.com/OpenSC/pkcs11-helper/releases/download/pkcs11-helper-1.30.0/pkcs11-helper-1.30.0.tar.bz2"
	pkcs11helper_cmd_zip := exec.Command("curl", "-L", pkcs11helper_zip_url, "-o", "package.zip")
	err = pkcs11helper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pkcs11helper_bin_url := "https://github.com/OpenSC/pkcs11-helper/releases/download/pkcs11-helper-1.30.0/pkcs11-helper-1.30.0.tar.bz2"
	pkcs11helper_cmd_bin := exec.Command("curl", "-L", pkcs11helper_bin_url, "-o", "binary.bin")
	err = pkcs11helper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pkcs11helper_src_url := "https://github.com/OpenSC/pkcs11-helper/releases/download/pkcs11-helper-1.30.0/pkcs11-helper-1.30.0.tar.bz2"
	pkcs11helper_cmd_src := exec.Command("curl", "-L", pkcs11helper_src_url, "-o", "source.tar.gz")
	err = pkcs11helper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pkcs11helper_cmd_direct := exec.Command("./binary")
	err = pkcs11helper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
