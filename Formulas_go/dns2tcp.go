package main

// Dns2tcp - TCP over DNS tunnel
// Homepage: https://packages.debian.org/sid/dns2tcp

import (
	"fmt"
	
	"os/exec"
)

func installDns2tcp() {
	// Método 1: Descargar y extraer .tar.gz
	dns2tcp_tar_url := "https://deb.debian.org/debian/pool/main/d/dns2tcp/dns2tcp_0.5.2.orig.tar.gz"
	dns2tcp_cmd_tar := exec.Command("curl", "-L", dns2tcp_tar_url, "-o", "package.tar.gz")
	err := dns2tcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dns2tcp_zip_url := "https://deb.debian.org/debian/pool/main/d/dns2tcp/dns2tcp_0.5.2.orig.zip"
	dns2tcp_cmd_zip := exec.Command("curl", "-L", dns2tcp_zip_url, "-o", "package.zip")
	err = dns2tcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dns2tcp_bin_url := "https://deb.debian.org/debian/pool/main/d/dns2tcp/dns2tcp_0.5.2.orig.bin"
	dns2tcp_cmd_bin := exec.Command("curl", "-L", dns2tcp_bin_url, "-o", "binary.bin")
	err = dns2tcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dns2tcp_src_url := "https://deb.debian.org/debian/pool/main/d/dns2tcp/dns2tcp_0.5.2.orig.src.tar.gz"
	dns2tcp_cmd_src := exec.Command("curl", "-L", dns2tcp_src_url, "-o", "source.tar.gz")
	err = dns2tcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dns2tcp_cmd_direct := exec.Command("./binary")
	err = dns2tcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
