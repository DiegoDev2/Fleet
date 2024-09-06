package main

// Tor - Anonymizing overlay network for TCP
// Homepage: https://www.torproject.org/

import (
	"fmt"
	
	"os/exec"
)

func installTor() {
	// Método 1: Descargar y extraer .tar.gz
	tor_tar_url := "https://www.torproject.org/dist/tor-0.4.8.12.tar.gz"
	tor_cmd_tar := exec.Command("curl", "-L", tor_tar_url, "-o", "package.tar.gz")
	err := tor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tor_zip_url := "https://www.torproject.org/dist/tor-0.4.8.12.zip"
	tor_cmd_zip := exec.Command("curl", "-L", tor_zip_url, "-o", "package.zip")
	err = tor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tor_bin_url := "https://www.torproject.org/dist/tor-0.4.8.12.bin"
	tor_cmd_bin := exec.Command("curl", "-L", tor_bin_url, "-o", "binary.bin")
	err = tor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tor_src_url := "https://www.torproject.org/dist/tor-0.4.8.12.src.tar.gz"
	tor_cmd_src := exec.Command("curl", "-L", tor_src_url, "-o", "source.tar.gz")
	err = tor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tor_cmd_direct := exec.Command("./binary")
	err = tor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libscrypt")
	exec.Command("latte", "install", "libscrypt").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
