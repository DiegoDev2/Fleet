package main

// Mfcuk - MiFare Classic Universal toolKit
// Homepage: https://github.com/nfc-tools/mfcuk

import (
	"fmt"
	
	"os/exec"
)

func installMfcuk() {
	// Método 1: Descargar y extraer .tar.gz
	mfcuk_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mfcuk/mfcuk-0.3.8.tar.gz"
	mfcuk_cmd_tar := exec.Command("curl", "-L", mfcuk_tar_url, "-o", "package.tar.gz")
	err := mfcuk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mfcuk_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mfcuk/mfcuk-0.3.8.zip"
	mfcuk_cmd_zip := exec.Command("curl", "-L", mfcuk_zip_url, "-o", "package.zip")
	err = mfcuk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mfcuk_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mfcuk/mfcuk-0.3.8.bin"
	mfcuk_cmd_bin := exec.Command("curl", "-L", mfcuk_bin_url, "-o", "binary.bin")
	err = mfcuk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mfcuk_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/mfcuk/mfcuk-0.3.8.src.tar.gz"
	mfcuk_cmd_src := exec.Command("curl", "-L", mfcuk_src_url, "-o", "source.tar.gz")
	err = mfcuk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mfcuk_cmd_direct := exec.Command("./binary")
	err = mfcuk_cmd_direct.Run()
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
