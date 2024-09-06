package main

// Ocproxy - User-level SOCKS and port forwarding proxy
// Homepage: https://github.com/cernekee/ocproxy

import (
	"fmt"
	
	"os/exec"
)

func installOcproxy() {
	// Método 1: Descargar y extraer .tar.gz
	ocproxy_tar_url := "https://github.com/cernekee/ocproxy/archive/refs/tags/v1.60.tar.gz"
	ocproxy_cmd_tar := exec.Command("curl", "-L", ocproxy_tar_url, "-o", "package.tar.gz")
	err := ocproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocproxy_zip_url := "https://github.com/cernekee/ocproxy/archive/refs/tags/v1.60.zip"
	ocproxy_cmd_zip := exec.Command("curl", "-L", ocproxy_zip_url, "-o", "package.zip")
	err = ocproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocproxy_bin_url := "https://github.com/cernekee/ocproxy/archive/refs/tags/v1.60.bin"
	ocproxy_cmd_bin := exec.Command("curl", "-L", ocproxy_bin_url, "-o", "binary.bin")
	err = ocproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocproxy_src_url := "https://github.com/cernekee/ocproxy/archive/refs/tags/v1.60.src.tar.gz"
	ocproxy_cmd_src := exec.Command("curl", "-L", ocproxy_src_url, "-o", "source.tar.gz")
	err = ocproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocproxy_cmd_direct := exec.Command("./binary")
	err = ocproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
}
