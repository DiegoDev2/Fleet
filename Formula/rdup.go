package main

// Rdup - Utility to create a file list suitable for making backups
// Homepage: https://github.com/miekg/rdup

import (
	"fmt"
	
	"os/exec"
)

func installRdup() {
	// Método 1: Descargar y extraer .tar.gz
	rdup_tar_url := "https://github.com/miekg/rdup/archive/refs/tags/1.1.15.tar.gz"
	rdup_cmd_tar := exec.Command("curl", "-L", rdup_tar_url, "-o", "package.tar.gz")
	err := rdup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rdup_zip_url := "https://github.com/miekg/rdup/archive/refs/tags/1.1.15.zip"
	rdup_cmd_zip := exec.Command("curl", "-L", rdup_zip_url, "-o", "package.zip")
	err = rdup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rdup_bin_url := "https://github.com/miekg/rdup/archive/refs/tags/1.1.15.bin"
	rdup_cmd_bin := exec.Command("curl", "-L", rdup_bin_url, "-o", "binary.bin")
	err = rdup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rdup_src_url := "https://github.com/miekg/rdup/archive/refs/tags/1.1.15.src.tar.gz"
	rdup_cmd_src := exec.Command("curl", "-L", rdup_src_url, "-o", "source.tar.gz")
	err = rdup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rdup_cmd_direct := exec.Command("./binary")
	err = rdup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: mcrypt")
	exec.Command("latte", "install", "mcrypt").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
