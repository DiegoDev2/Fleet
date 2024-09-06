package main

// MemcacheTop - Grab real-time stats from memcache
// Homepage: https://code.google.com/archive/p/memcache-top/

import (
	"fmt"
	
	"os/exec"
)

func installMemcacheTop() {
	// Método 1: Descargar y extraer .tar.gz
	memcachetop_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcache-top/memcache-top-v0.6"
	memcachetop_cmd_tar := exec.Command("curl", "-L", memcachetop_tar_url, "-o", "package.tar.gz")
	err := memcachetop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	memcachetop_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcache-top/memcache-top-v0.6"
	memcachetop_cmd_zip := exec.Command("curl", "-L", memcachetop_zip_url, "-o", "package.zip")
	err = memcachetop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	memcachetop_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcache-top/memcache-top-v0.6"
	memcachetop_cmd_bin := exec.Command("curl", "-L", memcachetop_bin_url, "-o", "binary.bin")
	err = memcachetop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	memcachetop_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/memcache-top/memcache-top-v0.6"
	memcachetop_cmd_src := exec.Command("curl", "-L", memcachetop_src_url, "-o", "source.tar.gz")
	err = memcachetop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	memcachetop_cmd_direct := exec.Command("./binary")
	err = memcachetop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
