package main

// Burp - Network backup and restore
// Homepage: https://burp.grke.org/

import (
	"fmt"
	
	"os/exec"
)

func installBurp() {
	// Método 1: Descargar y extraer .tar.gz
	burp_tar_url := "https://github.com/grke/burp/releases/download/2.4.0/burp-2.4.0.tar.bz2"
	burp_cmd_tar := exec.Command("curl", "-L", burp_tar_url, "-o", "package.tar.gz")
	err := burp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	burp_zip_url := "https://github.com/grke/burp/releases/download/2.4.0/burp-2.4.0.tar.bz2"
	burp_cmd_zip := exec.Command("curl", "-L", burp_zip_url, "-o", "package.zip")
	err = burp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	burp_bin_url := "https://github.com/grke/burp/releases/download/2.4.0/burp-2.4.0.tar.bz2"
	burp_cmd_bin := exec.Command("curl", "-L", burp_bin_url, "-o", "binary.bin")
	err = burp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	burp_src_url := "https://github.com/grke/burp/releases/download/2.4.0/burp-2.4.0.tar.bz2"
	burp_cmd_src := exec.Command("curl", "-L", burp_src_url, "-o", "source.tar.gz")
	err = burp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	burp_cmd_direct := exec.Command("./binary")
	err = burp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librsync")
	exec.Command("latte", "install", "librsync").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
