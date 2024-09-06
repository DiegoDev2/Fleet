package main

// Nmap - Port scanning utility for large networks
// Homepage: https://nmap.org/

import (
	"fmt"
	
	"os/exec"
)

func installNmap() {
	// Método 1: Descargar y extraer .tar.gz
	nmap_tar_url := "https://nmap.org/dist/nmap-7.95.tar.bz2"
	nmap_cmd_tar := exec.Command("curl", "-L", nmap_tar_url, "-o", "package.tar.gz")
	err := nmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nmap_zip_url := "https://nmap.org/dist/nmap-7.95.tar.bz2"
	nmap_cmd_zip := exec.Command("curl", "-L", nmap_zip_url, "-o", "package.zip")
	err = nmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nmap_bin_url := "https://nmap.org/dist/nmap-7.95.tar.bz2"
	nmap_cmd_bin := exec.Command("curl", "-L", nmap_bin_url, "-o", "binary.bin")
	err = nmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nmap_src_url := "https://nmap.org/dist/nmap-7.95.tar.bz2"
	nmap_cmd_src := exec.Command("curl", "-L", nmap_src_url, "-o", "source.tar.gz")
	err = nmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nmap_cmd_direct := exec.Command("./binary")
	err = nmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: liblinear")
	exec.Command("latte", "install", "liblinear").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
