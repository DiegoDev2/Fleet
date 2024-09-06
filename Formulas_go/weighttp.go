package main

// Weighttp - Webserver benchmarking tool that supports multithreading
// Homepage: https://redmine.lighttpd.net/projects/weighttp/wiki

import (
	"fmt"
	
	"os/exec"
)

func installWeighttp() {
	// Método 1: Descargar y extraer .tar.gz
	weighttp_tar_url := "https://github.com/lighttpd/weighttp/archive/refs/tags/weighttp-0.4.tar.gz"
	weighttp_cmd_tar := exec.Command("curl", "-L", weighttp_tar_url, "-o", "package.tar.gz")
	err := weighttp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	weighttp_zip_url := "https://github.com/lighttpd/weighttp/archive/refs/tags/weighttp-0.4.zip"
	weighttp_cmd_zip := exec.Command("curl", "-L", weighttp_zip_url, "-o", "package.zip")
	err = weighttp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	weighttp_bin_url := "https://github.com/lighttpd/weighttp/archive/refs/tags/weighttp-0.4.bin"
	weighttp_cmd_bin := exec.Command("curl", "-L", weighttp_bin_url, "-o", "binary.bin")
	err = weighttp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	weighttp_src_url := "https://github.com/lighttpd/weighttp/archive/refs/tags/weighttp-0.4.src.tar.gz"
	weighttp_cmd_src := exec.Command("curl", "-L", weighttp_src_url, "-o", "source.tar.gz")
	err = weighttp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	weighttp_cmd_direct := exec.Command("./binary")
	err = weighttp_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
}
