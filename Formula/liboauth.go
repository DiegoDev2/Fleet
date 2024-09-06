package main

// Liboauth - C library for the OAuth Core RFC 5849 standard
// Homepage: https://sourceforge.net/projects/liboauth/

import (
	"fmt"
	
	"os/exec"
)

func installLiboauth() {
	// Método 1: Descargar y extraer .tar.gz
	liboauth_tar_url := "https://downloads.sourceforge.net/project/liboauth/liboauth-1.0.3.tar.gz"
	liboauth_cmd_tar := exec.Command("curl", "-L", liboauth_tar_url, "-o", "package.tar.gz")
	err := liboauth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liboauth_zip_url := "https://downloads.sourceforge.net/project/liboauth/liboauth-1.0.3.zip"
	liboauth_cmd_zip := exec.Command("curl", "-L", liboauth_zip_url, "-o", "package.zip")
	err = liboauth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liboauth_bin_url := "https://downloads.sourceforge.net/project/liboauth/liboauth-1.0.3.bin"
	liboauth_cmd_bin := exec.Command("curl", "-L", liboauth_bin_url, "-o", "binary.bin")
	err = liboauth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liboauth_src_url := "https://downloads.sourceforge.net/project/liboauth/liboauth-1.0.3.src.tar.gz"
	liboauth_cmd_src := exec.Command("curl", "-L", liboauth_src_url, "-o", "source.tar.gz")
	err = liboauth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liboauth_cmd_direct := exec.Command("./binary")
	err = liboauth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
