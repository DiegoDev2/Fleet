package main

// Sslsplit - Man-in-the-middle attacks against SSL encrypted network connections
// Homepage: https://www.roe.ch/SSLsplit

import (
	"fmt"
	
	"os/exec"
)

func installSslsplit() {
	// Método 1: Descargar y extraer .tar.gz
	sslsplit_tar_url := "https://github.com/droe/sslsplit/archive/refs/tags/0.5.5.tar.gz"
	sslsplit_cmd_tar := exec.Command("curl", "-L", sslsplit_tar_url, "-o", "package.tar.gz")
	err := sslsplit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sslsplit_zip_url := "https://github.com/droe/sslsplit/archive/refs/tags/0.5.5.zip"
	sslsplit_cmd_zip := exec.Command("curl", "-L", sslsplit_zip_url, "-o", "package.zip")
	err = sslsplit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sslsplit_bin_url := "https://github.com/droe/sslsplit/archive/refs/tags/0.5.5.bin"
	sslsplit_cmd_bin := exec.Command("curl", "-L", sslsplit_bin_url, "-o", "binary.bin")
	err = sslsplit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sslsplit_src_url := "https://github.com/droe/sslsplit/archive/refs/tags/0.5.5.src.tar.gz"
	sslsplit_cmd_src := exec.Command("curl", "-L", sslsplit_src_url, "-o", "source.tar.gz")
	err = sslsplit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sslsplit_cmd_direct := exec.Command("./binary")
	err = sslsplit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: check")
	exec.Command("latte", "install", "check").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libnet")
	exec.Command("latte", "install", "libnet").Run()
	fmt.Println("Instalando dependencia: libpcap")
	exec.Command("latte", "install", "libpcap").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
