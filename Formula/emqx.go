package main

// Emqx - MQTT broker for IoT
// Homepage: https://www.emqx.io/

import (
	"fmt"
	
	"os/exec"
)

func installEmqx() {
	// Método 1: Descargar y extraer .tar.gz
	emqx_tar_url := "https://github.com/emqx/emqx/archive/refs/tags/v5.7.2.tar.gz"
	emqx_cmd_tar := exec.Command("curl", "-L", emqx_tar_url, "-o", "package.tar.gz")
	err := emqx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emqx_zip_url := "https://github.com/emqx/emqx/archive/refs/tags/v5.7.2.zip"
	emqx_cmd_zip := exec.Command("curl", "-L", emqx_zip_url, "-o", "package.zip")
	err = emqx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emqx_bin_url := "https://github.com/emqx/emqx/archive/refs/tags/v5.7.2.bin"
	emqx_cmd_bin := exec.Command("curl", "-L", emqx_bin_url, "-o", "binary.bin")
	err = emqx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emqx_src_url := "https://github.com/emqx/emqx/archive/refs/tags/v5.7.2.src.tar.gz"
	emqx_cmd_src := exec.Command("curl", "-L", emqx_src_url, "-o", "source.tar.gz")
	err = emqx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emqx_cmd_direct := exec.Command("./binary")
	err = emqx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: ccache")
	exec.Command("latte", "install", "ccache").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: erlang")
	exec.Command("latte", "install", "erlang").Run()
	fmt.Println("Instalando dependencia: freetds")
	exec.Command("latte", "install", "freetds").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
}
