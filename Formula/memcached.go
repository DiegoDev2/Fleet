package main

// Memcached - High performance, distributed memory object caching system
// Homepage: https://memcached.org/

import (
	"fmt"
	
	"os/exec"
)

func installMemcached() {
	// Método 1: Descargar y extraer .tar.gz
	memcached_tar_url := "https://www.memcached.org/files/memcached-1.6.30.tar.gz"
	memcached_cmd_tar := exec.Command("curl", "-L", memcached_tar_url, "-o", "package.tar.gz")
	err := memcached_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	memcached_zip_url := "https://www.memcached.org/files/memcached-1.6.30.zip"
	memcached_cmd_zip := exec.Command("curl", "-L", memcached_zip_url, "-o", "package.zip")
	err = memcached_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	memcached_bin_url := "https://www.memcached.org/files/memcached-1.6.30.bin"
	memcached_cmd_bin := exec.Command("curl", "-L", memcached_bin_url, "-o", "binary.bin")
	err = memcached_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	memcached_src_url := "https://www.memcached.org/files/memcached-1.6.30.src.tar.gz"
	memcached_cmd_src := exec.Command("curl", "-L", memcached_src_url, "-o", "source.tar.gz")
	err = memcached_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	memcached_cmd_direct := exec.Command("./binary")
	err = memcached_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
