package main

// VpnSlice - Vpnc-script replacement for easy and secure split-tunnel VPN setup
// Homepage: https://github.com/dlenski/vpn-slice

import (
	"fmt"
	
	"os/exec"
)

func installVpnSlice() {
	// Método 1: Descargar y extraer .tar.gz
	vpnslice_tar_url := "https://files.pythonhosted.org/packages/74/fd/6c9472e8ed83695abace098d83ba0df4ea48e29e7b2f6c77ced73b9f7dce/vpn-slice-0.16.1.tar.gz"
	vpnslice_cmd_tar := exec.Command("curl", "-L", vpnslice_tar_url, "-o", "package.tar.gz")
	err := vpnslice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vpnslice_zip_url := "https://files.pythonhosted.org/packages/74/fd/6c9472e8ed83695abace098d83ba0df4ea48e29e7b2f6c77ced73b9f7dce/vpn-slice-0.16.1.zip"
	vpnslice_cmd_zip := exec.Command("curl", "-L", vpnslice_zip_url, "-o", "package.zip")
	err = vpnslice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vpnslice_bin_url := "https://files.pythonhosted.org/packages/74/fd/6c9472e8ed83695abace098d83ba0df4ea48e29e7b2f6c77ced73b9f7dce/vpn-slice-0.16.1.bin"
	vpnslice_cmd_bin := exec.Command("curl", "-L", vpnslice_bin_url, "-o", "binary.bin")
	err = vpnslice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vpnslice_src_url := "https://files.pythonhosted.org/packages/74/fd/6c9472e8ed83695abace098d83ba0df4ea48e29e7b2f6c77ced73b9f7dce/vpn-slice-0.16.1.src.tar.gz"
	vpnslice_cmd_src := exec.Command("curl", "-L", vpnslice_src_url, "-o", "source.tar.gz")
	err = vpnslice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vpnslice_cmd_direct := exec.Command("./binary")
	err = vpnslice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
