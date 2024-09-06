package main

// Nfcutils - Near Field Communication (NFC) tools under POSIX systems
// Homepage: https://github.com/nfc-tools/nfcutils

import (
	"fmt"
	
	"os/exec"
)

func installNfcutils() {
	// Método 1: Descargar y extraer .tar.gz
	nfcutils_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/nfc-tools/nfcutils-0.3.2.tar.gz"
	nfcutils_cmd_tar := exec.Command("curl", "-L", nfcutils_tar_url, "-o", "package.tar.gz")
	err := nfcutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nfcutils_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/nfc-tools/nfcutils-0.3.2.zip"
	nfcutils_cmd_zip := exec.Command("curl", "-L", nfcutils_zip_url, "-o", "package.zip")
	err = nfcutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nfcutils_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/nfc-tools/nfcutils-0.3.2.bin"
	nfcutils_cmd_bin := exec.Command("curl", "-L", nfcutils_bin_url, "-o", "binary.bin")
	err = nfcutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nfcutils_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/nfc-tools/nfcutils-0.3.2.src.tar.gz"
	nfcutils_cmd_src := exec.Command("curl", "-L", nfcutils_src_url, "-o", "source.tar.gz")
	err = nfcutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nfcutils_cmd_direct := exec.Command("./binary")
	err = nfcutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libnfc")
exec.Command("latte", "install", "libnfc")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
