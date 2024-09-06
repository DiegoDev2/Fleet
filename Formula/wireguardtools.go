package main

// WireguardTools - Tools for the WireGuard secure network tunnel
// Homepage: https://www.wireguard.com/

import (
	"fmt"
	
	"os/exec"
)

func installWireguardTools() {
	// Método 1: Descargar y extraer .tar.gz
	wireguardtools_tar_url := "https://git.zx2c4.com/wireguard-tools/snapshot/wireguard-tools-1.0.20210914.tar.xz"
	wireguardtools_cmd_tar := exec.Command("curl", "-L", wireguardtools_tar_url, "-o", "package.tar.gz")
	err := wireguardtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wireguardtools_zip_url := "https://git.zx2c4.com/wireguard-tools/snapshot/wireguard-tools-1.0.20210914.tar.xz"
	wireguardtools_cmd_zip := exec.Command("curl", "-L", wireguardtools_zip_url, "-o", "package.zip")
	err = wireguardtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wireguardtools_bin_url := "https://git.zx2c4.com/wireguard-tools/snapshot/wireguard-tools-1.0.20210914.tar.xz"
	wireguardtools_cmd_bin := exec.Command("curl", "-L", wireguardtools_bin_url, "-o", "binary.bin")
	err = wireguardtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wireguardtools_src_url := "https://git.zx2c4.com/wireguard-tools/snapshot/wireguard-tools-1.0.20210914.tar.xz"
	wireguardtools_cmd_src := exec.Command("curl", "-L", wireguardtools_src_url, "-o", "source.tar.gz")
	err = wireguardtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wireguardtools_cmd_direct := exec.Command("./binary")
	err = wireguardtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: wireguard-go")
	exec.Command("latte", "install", "wireguard-go").Run()
}
