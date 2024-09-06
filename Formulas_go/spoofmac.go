package main

// SpoofMac - Spoof your MAC address in macOS
// Homepage: https://github.com/feross/SpoofMAC

import (
	"fmt"
	
	"os/exec"
)

func installSpoofMac() {
	// Método 1: Descargar y extraer .tar.gz
	spoofmac_tar_url := "https://files.pythonhosted.org/packages/9c/59/cc52a4c5d97b01fac7ff048353f8dc96f217eadc79022f78455e85144028/SpoofMAC-2.1.1.tar.gz"
	spoofmac_cmd_tar := exec.Command("curl", "-L", spoofmac_tar_url, "-o", "package.tar.gz")
	err := spoofmac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spoofmac_zip_url := "https://files.pythonhosted.org/packages/9c/59/cc52a4c5d97b01fac7ff048353f8dc96f217eadc79022f78455e85144028/SpoofMAC-2.1.1.zip"
	spoofmac_cmd_zip := exec.Command("curl", "-L", spoofmac_zip_url, "-o", "package.zip")
	err = spoofmac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spoofmac_bin_url := "https://files.pythonhosted.org/packages/9c/59/cc52a4c5d97b01fac7ff048353f8dc96f217eadc79022f78455e85144028/SpoofMAC-2.1.1.bin"
	spoofmac_cmd_bin := exec.Command("curl", "-L", spoofmac_bin_url, "-o", "binary.bin")
	err = spoofmac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spoofmac_src_url := "https://files.pythonhosted.org/packages/9c/59/cc52a4c5d97b01fac7ff048353f8dc96f217eadc79022f78455e85144028/SpoofMAC-2.1.1.src.tar.gz"
	spoofmac_cmd_src := exec.Command("curl", "-L", spoofmac_src_url, "-o", "source.tar.gz")
	err = spoofmac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spoofmac_cmd_direct := exec.Command("./binary")
	err = spoofmac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: net-tools")
exec.Command("latte", "install", "net-tools")
}
