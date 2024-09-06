package main

// Asdf - Extendable version manager with support for Ruby, Node.js, Erlang & more
// Homepage: https://asdf-vm.com/

import (
	"fmt"
	
	"os/exec"
)

func installAsdf() {
	// Método 1: Descargar y extraer .tar.gz
	asdf_tar_url := "https://github.com/asdf-vm/asdf/archive/refs/tags/v0.14.1.tar.gz"
	asdf_cmd_tar := exec.Command("curl", "-L", asdf_tar_url, "-o", "package.tar.gz")
	err := asdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asdf_zip_url := "https://github.com/asdf-vm/asdf/archive/refs/tags/v0.14.1.zip"
	asdf_cmd_zip := exec.Command("curl", "-L", asdf_zip_url, "-o", "package.zip")
	err = asdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asdf_bin_url := "https://github.com/asdf-vm/asdf/archive/refs/tags/v0.14.1.bin"
	asdf_cmd_bin := exec.Command("curl", "-L", asdf_bin_url, "-o", "binary.bin")
	err = asdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asdf_src_url := "https://github.com/asdf-vm/asdf/archive/refs/tags/v0.14.1.src.tar.gz"
	asdf_cmd_src := exec.Command("curl", "-L", asdf_src_url, "-o", "source.tar.gz")
	err = asdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asdf_cmd_direct := exec.Command("./binary")
	err = asdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
}
