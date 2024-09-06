package main

// Noir - Attack surface detector that identifies endpoints by static analysis
// Homepage: https://github.com/owasp-noir/noir

import (
	"fmt"
	
	"os/exec"
)

func installNoir() {
	// Método 1: Descargar y extraer .tar.gz
	noir_tar_url := "https://github.com/owasp-noir/noir/archive/refs/tags/v0.17.0.tar.gz"
	noir_cmd_tar := exec.Command("curl", "-L", noir_tar_url, "-o", "package.tar.gz")
	err := noir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	noir_zip_url := "https://github.com/owasp-noir/noir/archive/refs/tags/v0.17.0.zip"
	noir_cmd_zip := exec.Command("curl", "-L", noir_zip_url, "-o", "package.zip")
	err = noir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	noir_bin_url := "https://github.com/owasp-noir/noir/archive/refs/tags/v0.17.0.bin"
	noir_cmd_bin := exec.Command("curl", "-L", noir_bin_url, "-o", "binary.bin")
	err = noir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	noir_src_url := "https://github.com/owasp-noir/noir/archive/refs/tags/v0.17.0.src.tar.gz"
	noir_cmd_src := exec.Command("curl", "-L", noir_src_url, "-o", "source.tar.gz")
	err = noir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	noir_cmd_direct := exec.Command("./binary")
	err = noir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: crystal")
exec.Command("latte", "install", "crystal")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
