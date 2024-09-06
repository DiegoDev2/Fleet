package main

// Redex - Bytecode optimizer for Android apps
// Homepage: https://github.com/facebook/redex

import (
	"fmt"
	
	"os/exec"
)

func installRedex() {
	// Método 1: Descargar y extraer .tar.gz
	redex_tar_url := "https://github.com/facebook/redex/archive/refs/tags/v2017.10.31.tar.gz"
	redex_cmd_tar := exec.Command("curl", "-L", redex_tar_url, "-o", "package.tar.gz")
	err := redex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redex_zip_url := "https://github.com/facebook/redex/archive/refs/tags/v2017.10.31.zip"
	redex_cmd_zip := exec.Command("curl", "-L", redex_zip_url, "-o", "package.zip")
	err = redex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redex_bin_url := "https://github.com/facebook/redex/archive/refs/tags/v2017.10.31.bin"
	redex_cmd_bin := exec.Command("curl", "-L", redex_bin_url, "-o", "binary.bin")
	err = redex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redex_src_url := "https://github.com/facebook/redex/archive/refs/tags/v2017.10.31.src.tar.gz"
	redex_cmd_src := exec.Command("curl", "-L", redex_src_url, "-o", "source.tar.gz")
	err = redex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redex_cmd_direct := exec.Command("./binary")
	err = redex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: jsoncpp")
	exec.Command("latte", "install", "jsoncpp").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
