package main

// CodeServer - Access VS Code through the browser
// Homepage: https://github.com/coder/code-server

import (
	"fmt"
	
	"os/exec"
)

func installCodeServer() {
	// Método 1: Descargar y extraer .tar.gz
	codeserver_tar_url := "https://registry.npmjs.org/code-server/-/code-server-4.92.2.tgz"
	codeserver_cmd_tar := exec.Command("curl", "-L", codeserver_tar_url, "-o", "package.tar.gz")
	err := codeserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codeserver_zip_url := "https://registry.npmjs.org/code-server/-/code-server-4.92.2.tgz"
	codeserver_cmd_zip := exec.Command("curl", "-L", codeserver_zip_url, "-o", "package.zip")
	err = codeserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codeserver_bin_url := "https://registry.npmjs.org/code-server/-/code-server-4.92.2.tgz"
	codeserver_cmd_bin := exec.Command("curl", "-L", codeserver_bin_url, "-o", "binary.bin")
	err = codeserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codeserver_src_url := "https://registry.npmjs.org/code-server/-/code-server-4.92.2.tgz"
	codeserver_cmd_src := exec.Command("curl", "-L", codeserver_src_url, "-o", "source.tar.gz")
	err = codeserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codeserver_cmd_direct := exec.Command("./binary")
	err = codeserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node@20")
exec.Command("latte", "install", "node@20")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: krb5")
exec.Command("latte", "install", "krb5")
	fmt.Println("Instalando dependencia: libsecret")
exec.Command("latte", "install", "libsecret")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxkbfile")
exec.Command("latte", "install", "libxkbfile")
}
