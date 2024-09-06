package main

// Serd - C library for RDF syntax
// Homepage: https://drobilla.net/software/serd.html

import (
	"fmt"
	
	"os/exec"
)

func installSerd() {
	// Método 1: Descargar y extraer .tar.gz
	serd_tar_url := "https://download.drobilla.net/serd-0.32.2.tar.xz"
	serd_cmd_tar := exec.Command("curl", "-L", serd_tar_url, "-o", "package.tar.gz")
	err := serd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	serd_zip_url := "https://download.drobilla.net/serd-0.32.2.tar.xz"
	serd_cmd_zip := exec.Command("curl", "-L", serd_zip_url, "-o", "package.zip")
	err = serd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	serd_bin_url := "https://download.drobilla.net/serd-0.32.2.tar.xz"
	serd_cmd_bin := exec.Command("curl", "-L", serd_bin_url, "-o", "binary.bin")
	err = serd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	serd_src_url := "https://download.drobilla.net/serd-0.32.2.tar.xz"
	serd_cmd_src := exec.Command("curl", "-L", serd_src_url, "-o", "source.tar.gz")
	err = serd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	serd_cmd_direct := exec.Command("./binary")
	err = serd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
}
