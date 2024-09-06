package main

// LibsignalProtocolC - Signal Protocol C Library
// Homepage: https://github.com/signalapp/libsignal-protocol-c

import (
	"fmt"
	
	"os/exec"
)

func installLibsignalProtocolC() {
	// Método 1: Descargar y extraer .tar.gz
	libsignalprotocolc_tar_url := "https://github.com/signalapp/libsignal-protocol-c/archive/refs/tags/v2.3.3.tar.gz"
	libsignalprotocolc_cmd_tar := exec.Command("curl", "-L", libsignalprotocolc_tar_url, "-o", "package.tar.gz")
	err := libsignalprotocolc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsignalprotocolc_zip_url := "https://github.com/signalapp/libsignal-protocol-c/archive/refs/tags/v2.3.3.zip"
	libsignalprotocolc_cmd_zip := exec.Command("curl", "-L", libsignalprotocolc_zip_url, "-o", "package.zip")
	err = libsignalprotocolc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsignalprotocolc_bin_url := "https://github.com/signalapp/libsignal-protocol-c/archive/refs/tags/v2.3.3.bin"
	libsignalprotocolc_cmd_bin := exec.Command("curl", "-L", libsignalprotocolc_bin_url, "-o", "binary.bin")
	err = libsignalprotocolc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsignalprotocolc_src_url := "https://github.com/signalapp/libsignal-protocol-c/archive/refs/tags/v2.3.3.src.tar.gz"
	libsignalprotocolc_cmd_src := exec.Command("curl", "-L", libsignalprotocolc_src_url, "-o", "source.tar.gz")
	err = libsignalprotocolc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsignalprotocolc_cmd_direct := exec.Command("./binary")
	err = libsignalprotocolc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
