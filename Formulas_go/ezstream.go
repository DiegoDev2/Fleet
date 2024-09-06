package main

// Ezstream - Client for Icecast streaming servers
// Homepage: https://icecast.org/ezstream/

import (
	"fmt"
	
	"os/exec"
)

func installEzstream() {
	// Método 1: Descargar y extraer .tar.gz
	ezstream_tar_url := "https://downloads.xiph.org/releases/ezstream/ezstream-1.0.2.tar.gz"
	ezstream_cmd_tar := exec.Command("curl", "-L", ezstream_tar_url, "-o", "package.tar.gz")
	err := ezstream_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ezstream_zip_url := "https://downloads.xiph.org/releases/ezstream/ezstream-1.0.2.zip"
	ezstream_cmd_zip := exec.Command("curl", "-L", ezstream_zip_url, "-o", "package.zip")
	err = ezstream_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ezstream_bin_url := "https://downloads.xiph.org/releases/ezstream/ezstream-1.0.2.bin"
	ezstream_cmd_bin := exec.Command("curl", "-L", ezstream_bin_url, "-o", "binary.bin")
	err = ezstream_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ezstream_src_url := "https://downloads.xiph.org/releases/ezstream/ezstream-1.0.2.src.tar.gz"
	ezstream_cmd_src := exec.Command("curl", "-L", ezstream_src_url, "-o", "source.tar.gz")
	err = ezstream_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ezstream_cmd_direct := exec.Command("./binary")
	err = ezstream_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: check")
exec.Command("latte", "install", "check")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libshout")
exec.Command("latte", "install", "libshout")
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
}
