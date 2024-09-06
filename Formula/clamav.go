package main

// Clamav - Anti-virus software
// Homepage: https://www.clamav.net/

import (
	"fmt"
	
	"os/exec"
)

func installClamav() {
	// Método 1: Descargar y extraer .tar.gz
	clamav_tar_url := "https://github.com/Cisco-Talos/clamav/releases/download/clamav-1.4.1/clamav-1.4.1.tar.gz"
	clamav_cmd_tar := exec.Command("curl", "-L", clamav_tar_url, "-o", "package.tar.gz")
	err := clamav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clamav_zip_url := "https://github.com/Cisco-Talos/clamav/releases/download/clamav-1.4.1/clamav-1.4.1.zip"
	clamav_cmd_zip := exec.Command("curl", "-L", clamav_zip_url, "-o", "package.zip")
	err = clamav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clamav_bin_url := "https://github.com/Cisco-Talos/clamav/releases/download/clamav-1.4.1/clamav-1.4.1.bin"
	clamav_cmd_bin := exec.Command("curl", "-L", clamav_bin_url, "-o", "binary.bin")
	err = clamav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clamav_src_url := "https://github.com/Cisco-Talos/clamav/releases/download/clamav-1.4.1/clamav-1.4.1.src.tar.gz"
	clamav_cmd_src := exec.Command("curl", "-L", clamav_src_url, "-o", "source.tar.gz")
	err = clamav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clamav_cmd_direct := exec.Command("./binary")
	err = clamav_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: yara")
	exec.Command("latte", "install", "yara").Run()
}
