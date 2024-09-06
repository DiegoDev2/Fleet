package main

// Memcacheq - Queue service for memcache
// Homepage: https://code.google.com/archive/p/memcacheq/

import (
	"fmt"
	
	"os/exec"
)

func installMemcacheq() {
	// Método 1: Descargar y extraer .tar.gz
	memcacheq_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcacheq/memcacheq-0.2.0.tar.gz"
	memcacheq_cmd_tar := exec.Command("curl", "-L", memcacheq_tar_url, "-o", "package.tar.gz")
	err := memcacheq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	memcacheq_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcacheq/memcacheq-0.2.0.zip"
	memcacheq_cmd_zip := exec.Command("curl", "-L", memcacheq_zip_url, "-o", "package.zip")
	err = memcacheq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	memcacheq_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcacheq/memcacheq-0.2.0.bin"
	memcacheq_cmd_bin := exec.Command("curl", "-L", memcacheq_bin_url, "-o", "binary.bin")
	err = memcacheq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	memcacheq_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcacheq/memcacheq-0.2.0.src.tar.gz"
	memcacheq_cmd_src := exec.Command("curl", "-L", memcacheq_src_url, "-o", "source.tar.gz")
	err = memcacheq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	memcacheq_cmd_direct := exec.Command("./binary")
	err = memcacheq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
}
