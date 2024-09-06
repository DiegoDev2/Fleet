package main

// Libbpf - Berkeley Packet Filter library
// Homepage: https://github.com/libbpf/libbpf

import (
	"fmt"
	
	"os/exec"
)

func installLibbpf() {
	// Método 1: Descargar y extraer .tar.gz
	libbpf_tar_url := "https://github.com/libbpf/libbpf/archive/refs/tags/v1.4.6.tar.gz"
	libbpf_cmd_tar := exec.Command("curl", "-L", libbpf_tar_url, "-o", "package.tar.gz")
	err := libbpf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbpf_zip_url := "https://github.com/libbpf/libbpf/archive/refs/tags/v1.4.6.zip"
	libbpf_cmd_zip := exec.Command("curl", "-L", libbpf_zip_url, "-o", "package.zip")
	err = libbpf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbpf_bin_url := "https://github.com/libbpf/libbpf/archive/refs/tags/v1.4.6.bin"
	libbpf_cmd_bin := exec.Command("curl", "-L", libbpf_bin_url, "-o", "binary.bin")
	err = libbpf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbpf_src_url := "https://github.com/libbpf/libbpf/archive/refs/tags/v1.4.6.src.tar.gz"
	libbpf_cmd_src := exec.Command("curl", "-L", libbpf_src_url, "-o", "source.tar.gz")
	err = libbpf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbpf_cmd_direct := exec.Command("./binary")
	err = libbpf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
}
