package main

// IrcdHybrid - High-performance secure IRC server
// Homepage: https://www.ircd-hybrid.org/

import (
	"fmt"
	
	"os/exec"
)

func installIrcdHybrid() {
	// Método 1: Descargar y extraer .tar.gz
	ircdhybrid_tar_url := "https://downloads.sourceforge.net/project/ircd-hybrid/ircd-hybrid/ircd-hybrid-8.2.45/ircd-hybrid-8.2.45.tgz"
	ircdhybrid_cmd_tar := exec.Command("curl", "-L", ircdhybrid_tar_url, "-o", "package.tar.gz")
	err := ircdhybrid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ircdhybrid_zip_url := "https://downloads.sourceforge.net/project/ircd-hybrid/ircd-hybrid/ircd-hybrid-8.2.45/ircd-hybrid-8.2.45.tgz"
	ircdhybrid_cmd_zip := exec.Command("curl", "-L", ircdhybrid_zip_url, "-o", "package.zip")
	err = ircdhybrid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ircdhybrid_bin_url := "https://downloads.sourceforge.net/project/ircd-hybrid/ircd-hybrid/ircd-hybrid-8.2.45/ircd-hybrid-8.2.45.tgz"
	ircdhybrid_cmd_bin := exec.Command("curl", "-L", ircdhybrid_bin_url, "-o", "binary.bin")
	err = ircdhybrid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ircdhybrid_src_url := "https://downloads.sourceforge.net/project/ircd-hybrid/ircd-hybrid/ircd-hybrid-8.2.45/ircd-hybrid-8.2.45.tgz"
	ircdhybrid_cmd_src := exec.Command("curl", "-L", ircdhybrid_src_url, "-o", "source.tar.gz")
	err = ircdhybrid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ircdhybrid_cmd_direct := exec.Command("./binary")
	err = ircdhybrid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
