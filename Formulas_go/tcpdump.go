package main

// Tcpdump - Command-line packet analyzer
// Homepage: https://www.tcpdump.org/

import (
	"fmt"
	
	"os/exec"
)

func installTcpdump() {
	// Método 1: Descargar y extraer .tar.gz
	tcpdump_tar_url := "https://www.tcpdump.org/release/tcpdump-4.99.5.tar.gz"
	tcpdump_cmd_tar := exec.Command("curl", "-L", tcpdump_tar_url, "-o", "package.tar.gz")
	err := tcpdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcpdump_zip_url := "https://www.tcpdump.org/release/tcpdump-4.99.5.zip"
	tcpdump_cmd_zip := exec.Command("curl", "-L", tcpdump_zip_url, "-o", "package.zip")
	err = tcpdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcpdump_bin_url := "https://www.tcpdump.org/release/tcpdump-4.99.5.bin"
	tcpdump_cmd_bin := exec.Command("curl", "-L", tcpdump_bin_url, "-o", "binary.bin")
	err = tcpdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcpdump_src_url := "https://www.tcpdump.org/release/tcpdump-4.99.5.src.tar.gz"
	tcpdump_cmd_src := exec.Command("curl", "-L", tcpdump_src_url, "-o", "source.tar.gz")
	err = tcpdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcpdump_cmd_direct := exec.Command("./binary")
	err = tcpdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpcap")
exec.Command("latte", "install", "libpcap")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
