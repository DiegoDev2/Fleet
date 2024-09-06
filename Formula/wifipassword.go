package main

// WifiPassword - Show the current WiFi network password
// Homepage: https://github.com/rauchg/wifi-password

import (
	"fmt"
	
	"os/exec"
)

func installWifiPassword() {
	// Método 1: Descargar y extraer .tar.gz
	wifipassword_tar_url := "https://github.com/rauchg/wifi-password/archive/refs/tags/0.1.0.tar.gz"
	wifipassword_cmd_tar := exec.Command("curl", "-L", wifipassword_tar_url, "-o", "package.tar.gz")
	err := wifipassword_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wifipassword_zip_url := "https://github.com/rauchg/wifi-password/archive/refs/tags/0.1.0.zip"
	wifipassword_cmd_zip := exec.Command("curl", "-L", wifipassword_zip_url, "-o", "package.zip")
	err = wifipassword_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wifipassword_bin_url := "https://github.com/rauchg/wifi-password/archive/refs/tags/0.1.0.bin"
	wifipassword_cmd_bin := exec.Command("curl", "-L", wifipassword_bin_url, "-o", "binary.bin")
	err = wifipassword_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wifipassword_src_url := "https://github.com/rauchg/wifi-password/archive/refs/tags/0.1.0.src.tar.gz"
	wifipassword_cmd_src := exec.Command("curl", "-L", wifipassword_src_url, "-o", "source.tar.gz")
	err = wifipassword_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wifipassword_cmd_direct := exec.Command("./binary")
	err = wifipassword_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
