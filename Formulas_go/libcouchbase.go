package main

// Libcouchbase - C library for Couchbase
// Homepage: https://docs.couchbase.com/c-sdk/current/hello-world/start-using-sdk.html

import (
	"fmt"
	
	"os/exec"
)

func installLibcouchbase() {
	// Método 1: Descargar y extraer .tar.gz
	libcouchbase_tar_url := "https://packages.couchbase.com/clients/c/libcouchbase-3.3.12.tar.gz"
	libcouchbase_cmd_tar := exec.Command("curl", "-L", libcouchbase_tar_url, "-o", "package.tar.gz")
	err := libcouchbase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcouchbase_zip_url := "https://packages.couchbase.com/clients/c/libcouchbase-3.3.12.zip"
	libcouchbase_cmd_zip := exec.Command("curl", "-L", libcouchbase_zip_url, "-o", "package.zip")
	err = libcouchbase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcouchbase_bin_url := "https://packages.couchbase.com/clients/c/libcouchbase-3.3.12.bin"
	libcouchbase_cmd_bin := exec.Command("curl", "-L", libcouchbase_bin_url, "-o", "binary.bin")
	err = libcouchbase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcouchbase_src_url := "https://packages.couchbase.com/clients/c/libcouchbase-3.3.12.src.tar.gz"
	libcouchbase_cmd_src := exec.Command("curl", "-L", libcouchbase_src_url, "-o", "source.tar.gz")
	err = libcouchbase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcouchbase_cmd_direct := exec.Command("./binary")
	err = libcouchbase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
