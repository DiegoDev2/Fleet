package main

// Killswitch - VPN kill switch for macOS
// Homepage: https://vpn-kill-switch.com

import (
	"fmt"
	
	"os/exec"
)

func installKillswitch() {
	// Método 1: Descargar y extraer .tar.gz
	killswitch_tar_url := "https://github.com/vpn-kill-switch/killswitch/archive/refs/tags/v0.7.3.tar.gz"
	killswitch_cmd_tar := exec.Command("curl", "-L", killswitch_tar_url, "-o", "package.tar.gz")
	err := killswitch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	killswitch_zip_url := "https://github.com/vpn-kill-switch/killswitch/archive/refs/tags/v0.7.3.zip"
	killswitch_cmd_zip := exec.Command("curl", "-L", killswitch_zip_url, "-o", "package.zip")
	err = killswitch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	killswitch_bin_url := "https://github.com/vpn-kill-switch/killswitch/archive/refs/tags/v0.7.3.bin"
	killswitch_cmd_bin := exec.Command("curl", "-L", killswitch_bin_url, "-o", "binary.bin")
	err = killswitch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	killswitch_src_url := "https://github.com/vpn-kill-switch/killswitch/archive/refs/tags/v0.7.3.src.tar.gz"
	killswitch_cmd_src := exec.Command("curl", "-L", killswitch_src_url, "-o", "source.tar.gz")
	err = killswitch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	killswitch_cmd_direct := exec.Command("./binary")
	err = killswitch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
