package main

// Libmemcached - C and C++ client library to the memcached server
// Homepage: https://libmemcached.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibmemcached() {
	// Método 1: Descargar y extraer .tar.gz
	libmemcached_tar_url := "https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.tar.gz"
	libmemcached_cmd_tar := exec.Command("curl", "-L", libmemcached_tar_url, "-o", "package.tar.gz")
	err := libmemcached_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmemcached_zip_url := "https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.zip"
	libmemcached_cmd_zip := exec.Command("curl", "-L", libmemcached_zip_url, "-o", "package.zip")
	err = libmemcached_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmemcached_bin_url := "https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.bin"
	libmemcached_cmd_bin := exec.Command("curl", "-L", libmemcached_bin_url, "-o", "binary.bin")
	err = libmemcached_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmemcached_src_url := "https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.src.tar.gz"
	libmemcached_cmd_src := exec.Command("curl", "-L", libmemcached_src_url, "-o", "source.tar.gz")
	err = libmemcached_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmemcached_cmd_direct := exec.Command("./binary")
	err = libmemcached_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: memcached")
exec.Command("latte", "install", "memcached")
}
