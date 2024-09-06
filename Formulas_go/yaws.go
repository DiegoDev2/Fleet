package main

// Yaws - Webserver for dynamic content (written in Erlang)
// Homepage: https://erlyaws.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installYaws() {
	// Método 1: Descargar y extraer .tar.gz
	yaws_tar_url := "https://github.com/erlyaws/yaws/archive/refs/tags/yaws-2.2.0.tar.gz"
	yaws_cmd_tar := exec.Command("curl", "-L", yaws_tar_url, "-o", "package.tar.gz")
	err := yaws_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yaws_zip_url := "https://github.com/erlyaws/yaws/archive/refs/tags/yaws-2.2.0.zip"
	yaws_cmd_zip := exec.Command("curl", "-L", yaws_zip_url, "-o", "package.zip")
	err = yaws_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yaws_bin_url := "https://github.com/erlyaws/yaws/archive/refs/tags/yaws-2.2.0.bin"
	yaws_cmd_bin := exec.Command("curl", "-L", yaws_bin_url, "-o", "binary.bin")
	err = yaws_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yaws_src_url := "https://github.com/erlyaws/yaws/archive/refs/tags/yaws-2.2.0.src.tar.gz"
	yaws_cmd_src := exec.Command("curl", "-L", yaws_src_url, "-o", "source.tar.gz")
	err = yaws_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yaws_cmd_direct := exec.Command("./binary")
	err = yaws_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: erlang")
exec.Command("latte", "install", "erlang")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
