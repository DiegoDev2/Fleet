package main

// Jemalloc - Implementation of malloc emphasizing fragmentation avoidance
// Homepage: https://jemalloc.net/

import (
	"fmt"
	
	"os/exec"
)

func installJemalloc() {
	// Método 1: Descargar y extraer .tar.gz
	jemalloc_tar_url := "https://github.com/jemalloc/jemalloc/releases/download/5.3.0/jemalloc-5.3.0.tar.bz2"
	jemalloc_cmd_tar := exec.Command("curl", "-L", jemalloc_tar_url, "-o", "package.tar.gz")
	err := jemalloc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jemalloc_zip_url := "https://github.com/jemalloc/jemalloc/releases/download/5.3.0/jemalloc-5.3.0.tar.bz2"
	jemalloc_cmd_zip := exec.Command("curl", "-L", jemalloc_zip_url, "-o", "package.zip")
	err = jemalloc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jemalloc_bin_url := "https://github.com/jemalloc/jemalloc/releases/download/5.3.0/jemalloc-5.3.0.tar.bz2"
	jemalloc_cmd_bin := exec.Command("curl", "-L", jemalloc_bin_url, "-o", "binary.bin")
	err = jemalloc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jemalloc_src_url := "https://github.com/jemalloc/jemalloc/releases/download/5.3.0/jemalloc-5.3.0.tar.bz2"
	jemalloc_cmd_src := exec.Command("curl", "-L", jemalloc_src_url, "-o", "source.tar.gz")
	err = jemalloc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jemalloc_cmd_direct := exec.Command("./binary")
	err = jemalloc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
}
