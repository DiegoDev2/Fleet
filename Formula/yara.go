package main

// Yara - Malware identification and classification tool
// Homepage: https://github.com/VirusTotal/yara/

import (
	"fmt"
	
	"os/exec"
)

func installYara() {
	// Método 1: Descargar y extraer .tar.gz
	yara_tar_url := "https://github.com/VirusTotal/yara/archive/refs/tags/v4.5.1.tar.gz"
	yara_cmd_tar := exec.Command("curl", "-L", yara_tar_url, "-o", "package.tar.gz")
	err := yara_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yara_zip_url := "https://github.com/VirusTotal/yara/archive/refs/tags/v4.5.1.zip"
	yara_cmd_zip := exec.Command("curl", "-L", yara_zip_url, "-o", "package.zip")
	err = yara_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yara_bin_url := "https://github.com/VirusTotal/yara/archive/refs/tags/v4.5.1.bin"
	yara_cmd_bin := exec.Command("curl", "-L", yara_bin_url, "-o", "binary.bin")
	err = yara_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yara_src_url := "https://github.com/VirusTotal/yara/archive/refs/tags/v4.5.1.src.tar.gz"
	yara_cmd_src := exec.Command("curl", "-L", yara_src_url, "-o", "source.tar.gz")
	err = yara_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yara_cmd_direct := exec.Command("./binary")
	err = yara_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf-c")
	exec.Command("latte", "install", "protobuf-c").Run()
}
