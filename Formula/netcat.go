package main

// Netcat - Utility for managing network connections
// Homepage: https://netcat.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installNetcat() {
	// Método 1: Descargar y extraer .tar.gz
	netcat_tar_url := "https://downloads.sourceforge.net/project/netcat/netcat/0.7.1/netcat-0.7.1.tar.bz2"
	netcat_cmd_tar := exec.Command("curl", "-L", netcat_tar_url, "-o", "package.tar.gz")
	err := netcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netcat_zip_url := "https://downloads.sourceforge.net/project/netcat/netcat/0.7.1/netcat-0.7.1.tar.bz2"
	netcat_cmd_zip := exec.Command("curl", "-L", netcat_zip_url, "-o", "package.zip")
	err = netcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netcat_bin_url := "https://downloads.sourceforge.net/project/netcat/netcat/0.7.1/netcat-0.7.1.tar.bz2"
	netcat_cmd_bin := exec.Command("curl", "-L", netcat_bin_url, "-o", "binary.bin")
	err = netcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netcat_src_url := "https://downloads.sourceforge.net/project/netcat/netcat/0.7.1/netcat-0.7.1.tar.bz2"
	netcat_cmd_src := exec.Command("curl", "-L", netcat_src_url, "-o", "source.tar.gz")
	err = netcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netcat_cmd_direct := exec.Command("./binary")
	err = netcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
