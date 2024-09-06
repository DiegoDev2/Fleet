package main

// Libbpg - Image format meant to improve on JPEG quality and file size
// Homepage: https://bellard.org/bpg/

import (
	"fmt"
	
	"os/exec"
)

func installLibbpg() {
	// Método 1: Descargar y extraer .tar.gz
	libbpg_tar_url := "https://bellard.org/bpg/libbpg-0.9.8.tar.gz"
	libbpg_cmd_tar := exec.Command("curl", "-L", libbpg_tar_url, "-o", "package.tar.gz")
	err := libbpg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbpg_zip_url := "https://bellard.org/bpg/libbpg-0.9.8.zip"
	libbpg_cmd_zip := exec.Command("curl", "-L", libbpg_zip_url, "-o", "package.zip")
	err = libbpg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbpg_bin_url := "https://bellard.org/bpg/libbpg-0.9.8.bin"
	libbpg_cmd_bin := exec.Command("curl", "-L", libbpg_bin_url, "-o", "binary.bin")
	err = libbpg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbpg_src_url := "https://bellard.org/bpg/libbpg-0.9.8.src.tar.gz"
	libbpg_cmd_src := exec.Command("curl", "-L", libbpg_src_url, "-o", "source.tar.gz")
	err = libbpg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbpg_cmd_direct := exec.Command("./binary")
	err = libbpg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: yasm")
	exec.Command("latte", "install", "yasm").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
