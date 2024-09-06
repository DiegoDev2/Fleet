package main

// LibcouchbaseAT2 - C library for Couchbase
// Homepage: https://docs-archive.couchbase.com/c-sdk/2.10/start-using-sdk.html

import (
	"fmt"
	
	"os/exec"
)

func installLibcouchbaseAT2() {
	// Método 1: Descargar y extraer .tar.gz
	libcouchbaseat2_tar_url := "https://packages.couchbase.com/clients/c/libcouchbase-2.10.9.tar.gz"
	libcouchbaseat2_cmd_tar := exec.Command("curl", "-L", libcouchbaseat2_tar_url, "-o", "package.tar.gz")
	err := libcouchbaseat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcouchbaseat2_zip_url := "https://packages.couchbase.com/clients/c/libcouchbase-2.10.9.zip"
	libcouchbaseat2_cmd_zip := exec.Command("curl", "-L", libcouchbaseat2_zip_url, "-o", "package.zip")
	err = libcouchbaseat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcouchbaseat2_bin_url := "https://packages.couchbase.com/clients/c/libcouchbase-2.10.9.bin"
	libcouchbaseat2_cmd_bin := exec.Command("curl", "-L", libcouchbaseat2_bin_url, "-o", "binary.bin")
	err = libcouchbaseat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcouchbaseat2_src_url := "https://packages.couchbase.com/clients/c/libcouchbase-2.10.9.src.tar.gz"
	libcouchbaseat2_cmd_src := exec.Command("curl", "-L", libcouchbaseat2_src_url, "-o", "source.tar.gz")
	err = libcouchbaseat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcouchbaseat2_cmd_direct := exec.Command("./binary")
	err = libcouchbaseat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libev")
	exec.Command("latte", "install", "libev").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
