package main

// WireguardGo - Userspace Go implementation of WireGuard
// Homepage: https://www.wireguard.com/

import (
	"fmt"
	
	"os/exec"
)

func installWireguardGo() {
	// Método 1: Descargar y extraer .tar.gz
	wireguardgo_tar_url := "https://git.zx2c4.com/wireguard-go/snapshot/wireguard-go-0.0.20230223.tar.xz"
	wireguardgo_cmd_tar := exec.Command("curl", "-L", wireguardgo_tar_url, "-o", "package.tar.gz")
	err := wireguardgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wireguardgo_zip_url := "https://git.zx2c4.com/wireguard-go/snapshot/wireguard-go-0.0.20230223.tar.xz"
	wireguardgo_cmd_zip := exec.Command("curl", "-L", wireguardgo_zip_url, "-o", "package.zip")
	err = wireguardgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wireguardgo_bin_url := "https://git.zx2c4.com/wireguard-go/snapshot/wireguard-go-0.0.20230223.tar.xz"
	wireguardgo_cmd_bin := exec.Command("curl", "-L", wireguardgo_bin_url, "-o", "binary.bin")
	err = wireguardgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wireguardgo_src_url := "https://git.zx2c4.com/wireguard-go/snapshot/wireguard-go-0.0.20230223.tar.xz"
	wireguardgo_cmd_src := exec.Command("curl", "-L", wireguardgo_src_url, "-o", "source.tar.gz")
	err = wireguardgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wireguardgo_cmd_direct := exec.Command("./binary")
	err = wireguardgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
