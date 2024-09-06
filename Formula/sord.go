package main

// Sord - C library for storing RDF data in memory
// Homepage: https://drobilla.net/software/sord.html

import (
	"fmt"
	
	"os/exec"
)

func installSord() {
	// Método 1: Descargar y extraer .tar.gz
	sord_tar_url := "https://download.drobilla.net/sord-0.16.16.tar.xz"
	sord_cmd_tar := exec.Command("curl", "-L", sord_tar_url, "-o", "package.tar.gz")
	err := sord_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sord_zip_url := "https://download.drobilla.net/sord-0.16.16.tar.xz"
	sord_cmd_zip := exec.Command("curl", "-L", sord_zip_url, "-o", "package.zip")
	err = sord_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sord_bin_url := "https://download.drobilla.net/sord-0.16.16.tar.xz"
	sord_cmd_bin := exec.Command("curl", "-L", sord_bin_url, "-o", "binary.bin")
	err = sord_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sord_src_url := "https://download.drobilla.net/sord-0.16.16.tar.xz"
	sord_cmd_src := exec.Command("curl", "-L", sord_src_url, "-o", "source.tar.gz")
	err = sord_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sord_cmd_direct := exec.Command("./binary")
	err = sord_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: serd")
	exec.Command("latte", "install", "serd").Run()
	fmt.Println("Instalando dependencia: zix")
	exec.Command("latte", "install", "zix").Run()
}
