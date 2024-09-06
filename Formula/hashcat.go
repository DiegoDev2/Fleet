package main

// Hashcat - World's fastest and most advanced password recovery utility
// Homepage: https://hashcat.net/hashcat/

import (
	"fmt"
	
	"os/exec"
)

func installHashcat() {
	// Método 1: Descargar y extraer .tar.gz
	hashcat_tar_url := "https://hashcat.net/files/hashcat-6.2.6.tar.gz"
	hashcat_cmd_tar := exec.Command("curl", "-L", hashcat_tar_url, "-o", "package.tar.gz")
	err := hashcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hashcat_zip_url := "https://hashcat.net/files/hashcat-6.2.6.zip"
	hashcat_cmd_zip := exec.Command("curl", "-L", hashcat_zip_url, "-o", "package.zip")
	err = hashcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hashcat_bin_url := "https://hashcat.net/files/hashcat-6.2.6.bin"
	hashcat_cmd_bin := exec.Command("curl", "-L", hashcat_bin_url, "-o", "binary.bin")
	err = hashcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hashcat_src_url := "https://hashcat.net/files/hashcat-6.2.6.src.tar.gz"
	hashcat_cmd_src := exec.Command("curl", "-L", hashcat_src_url, "-o", "source.tar.gz")
	err = hashcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hashcat_cmd_direct := exec.Command("./binary")
	err = hashcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: xxhash")
	exec.Command("latte", "install", "xxhash").Run()
	fmt.Println("Instalando dependencia: opencl-headers")
	exec.Command("latte", "install", "opencl-headers").Run()
	fmt.Println("Instalando dependencia: opencl-icd-loader")
	exec.Command("latte", "install", "opencl-icd-loader").Run()
	fmt.Println("Instalando dependencia: pocl")
	exec.Command("latte", "install", "pocl").Run()
}
