package main

// Privoxy - Advanced filtering web proxy
// Homepage: https://www.privoxy.org/

import (
	"fmt"
	
	"os/exec"
)

func installPrivoxy() {
	// Método 1: Descargar y extraer .tar.gz
	privoxy_tar_url := "https://downloads.sourceforge.net/project/ijbswa/Sources/3.0.34%20%28stable%29/privoxy-3.0.34-stable-src.tar.gz"
	privoxy_cmd_tar := exec.Command("curl", "-L", privoxy_tar_url, "-o", "package.tar.gz")
	err := privoxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	privoxy_zip_url := "https://downloads.sourceforge.net/project/ijbswa/Sources/3.0.34%20%28stable%29/privoxy-3.0.34-stable-src.zip"
	privoxy_cmd_zip := exec.Command("curl", "-L", privoxy_zip_url, "-o", "package.zip")
	err = privoxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	privoxy_bin_url := "https://downloads.sourceforge.net/project/ijbswa/Sources/3.0.34%20%28stable%29/privoxy-3.0.34-stable-src.bin"
	privoxy_cmd_bin := exec.Command("curl", "-L", privoxy_bin_url, "-o", "binary.bin")
	err = privoxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	privoxy_src_url := "https://downloads.sourceforge.net/project/ijbswa/Sources/3.0.34%20%28stable%29/privoxy-3.0.34-stable-src.src.tar.gz"
	privoxy_cmd_src := exec.Command("curl", "-L", privoxy_src_url, "-o", "source.tar.gz")
	err = privoxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	privoxy_cmd_direct := exec.Command("./binary")
	err = privoxy_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
