package main

// Bind - Implementation of the DNS protocols
// Homepage: https://www.isc.org/bind/

import (
	"fmt"
	
	"os/exec"
)

func installBind() {
	// Método 1: Descargar y extraer .tar.gz
	bind_tar_url := "https://downloads.isc.org/isc/bind9/9.20.1/bind-9.20.1.tar.xz"
	bind_cmd_tar := exec.Command("curl", "-L", bind_tar_url, "-o", "package.tar.gz")
	err := bind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bind_zip_url := "https://downloads.isc.org/isc/bind9/9.20.1/bind-9.20.1.tar.xz"
	bind_cmd_zip := exec.Command("curl", "-L", bind_zip_url, "-o", "package.zip")
	err = bind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bind_bin_url := "https://downloads.isc.org/isc/bind9/9.20.1/bind-9.20.1.tar.xz"
	bind_cmd_bin := exec.Command("curl", "-L", bind_bin_url, "-o", "binary.bin")
	err = bind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bind_src_url := "https://downloads.isc.org/isc/bind9/9.20.1/bind-9.20.1.tar.xz"
	bind_cmd_src := exec.Command("curl", "-L", bind_src_url, "-o", "source.tar.gz")
	err = bind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bind_cmd_direct := exec.Command("./binary")
	err = bind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jemalloc")
	exec.Command("latte", "install", "jemalloc").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: libnghttp2")
	exec.Command("latte", "install", "libnghttp2").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: userspace-rcu")
	exec.Command("latte", "install", "userspace-rcu").Run()
	fmt.Println("Instalando dependencia: libcap")
	exec.Command("latte", "install", "libcap").Run()
}
