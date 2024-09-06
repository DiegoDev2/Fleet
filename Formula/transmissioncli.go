package main

// TransmissionCli - Lightweight BitTorrent client
// Homepage: https://www.transmissionbt.com/

import (
	"fmt"
	
	"os/exec"
)

func installTransmissionCli() {
	// Método 1: Descargar y extraer .tar.gz
	transmissioncli_tar_url := "https://github.com/transmission/transmission/releases/download/4.0.6/transmission-4.0.6.tar.xz"
	transmissioncli_cmd_tar := exec.Command("curl", "-L", transmissioncli_tar_url, "-o", "package.tar.gz")
	err := transmissioncli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	transmissioncli_zip_url := "https://github.com/transmission/transmission/releases/download/4.0.6/transmission-4.0.6.tar.xz"
	transmissioncli_cmd_zip := exec.Command("curl", "-L", transmissioncli_zip_url, "-o", "package.zip")
	err = transmissioncli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	transmissioncli_bin_url := "https://github.com/transmission/transmission/releases/download/4.0.6/transmission-4.0.6.tar.xz"
	transmissioncli_cmd_bin := exec.Command("curl", "-L", transmissioncli_bin_url, "-o", "binary.bin")
	err = transmissioncli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	transmissioncli_src_url := "https://github.com/transmission/transmission/releases/download/4.0.6/transmission-4.0.6.tar.xz"
	transmissioncli_cmd_src := exec.Command("curl", "-L", transmissioncli_src_url, "-o", "source.tar.gz")
	err = transmissioncli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	transmissioncli_cmd_direct := exec.Command("./binary")
	err = transmissioncli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: miniupnpc")
	exec.Command("latte", "install", "miniupnpc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
