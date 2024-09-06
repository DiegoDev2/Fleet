package main

// Ssldump - SSLv3/TLS network protocol analyzer
// Homepage: https://adulau.github.io/ssldump/

import (
	"fmt"
	
	"os/exec"
)

func installSsldump() {
	// Método 1: Descargar y extraer .tar.gz
	ssldump_tar_url := "https://github.com/adulau/ssldump/archive/refs/tags/v1.8.tar.gz"
	ssldump_cmd_tar := exec.Command("curl", "-L", ssldump_tar_url, "-o", "package.tar.gz")
	err := ssldump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssldump_zip_url := "https://github.com/adulau/ssldump/archive/refs/tags/v1.8.zip"
	ssldump_cmd_zip := exec.Command("curl", "-L", ssldump_zip_url, "-o", "package.zip")
	err = ssldump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssldump_bin_url := "https://github.com/adulau/ssldump/archive/refs/tags/v1.8.bin"
	ssldump_cmd_bin := exec.Command("curl", "-L", ssldump_bin_url, "-o", "binary.bin")
	err = ssldump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssldump_src_url := "https://github.com/adulau/ssldump/archive/refs/tags/v1.8.src.tar.gz"
	ssldump_cmd_src := exec.Command("curl", "-L", ssldump_src_url, "-o", "source.tar.gz")
	err = ssldump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssldump_cmd_direct := exec.Command("./binary")
	err = ssldump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
	fmt.Println("Instalando dependencia: libpcap")
exec.Command("latte", "install", "libpcap")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
