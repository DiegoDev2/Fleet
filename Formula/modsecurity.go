package main

// Modsecurity - Libmodsecurity is one component of the ModSecurity v3 project
// Homepage: https://github.com/owasp-modsecurity/ModSecurity

import (
	"fmt"
	
	"os/exec"
)

func installModsecurity() {
	// Método 1: Descargar y extraer .tar.gz
	modsecurity_tar_url := "https://github.com/owasp-modsecurity/ModSecurity/releases/download/v3.0.13/modsecurity-v3.0.13.tar.gz"
	modsecurity_cmd_tar := exec.Command("curl", "-L", modsecurity_tar_url, "-o", "package.tar.gz")
	err := modsecurity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	modsecurity_zip_url := "https://github.com/owasp-modsecurity/ModSecurity/releases/download/v3.0.13/modsecurity-v3.0.13.zip"
	modsecurity_cmd_zip := exec.Command("curl", "-L", modsecurity_zip_url, "-o", "package.zip")
	err = modsecurity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	modsecurity_bin_url := "https://github.com/owasp-modsecurity/ModSecurity/releases/download/v3.0.13/modsecurity-v3.0.13.bin"
	modsecurity_cmd_bin := exec.Command("curl", "-L", modsecurity_bin_url, "-o", "binary.bin")
	err = modsecurity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	modsecurity_src_url := "https://github.com/owasp-modsecurity/ModSecurity/releases/download/v3.0.13/modsecurity-v3.0.13.src.tar.gz"
	modsecurity_cmd_src := exec.Command("curl", "-L", modsecurity_src_url, "-o", "source.tar.gz")
	err = modsecurity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	modsecurity_cmd_direct := exec.Command("./binary")
	err = modsecurity_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libmaxminddb")
	exec.Command("latte", "install", "libmaxminddb").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: yajl")
	exec.Command("latte", "install", "yajl").Run()
}
