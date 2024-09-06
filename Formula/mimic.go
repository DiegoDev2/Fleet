package main

// Mimic - Lightweight text-to-speech engine based on CMU Flite
// Homepage: https://github.com/MycroftAI/mimic1

import (
	"fmt"
	
	"os/exec"
)

func installMimic() {
	// Método 1: Descargar y extraer .tar.gz
	mimic_tar_url := "https://github.com/MycroftAI/mimic1/archive/refs/tags/1.3.0.1.tar.gz"
	mimic_cmd_tar := exec.Command("curl", "-L", mimic_tar_url, "-o", "package.tar.gz")
	err := mimic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mimic_zip_url := "https://github.com/MycroftAI/mimic1/archive/refs/tags/1.3.0.1.zip"
	mimic_cmd_zip := exec.Command("curl", "-L", mimic_zip_url, "-o", "package.zip")
	err = mimic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mimic_bin_url := "https://github.com/MycroftAI/mimic1/archive/refs/tags/1.3.0.1.bin"
	mimic_cmd_bin := exec.Command("curl", "-L", mimic_bin_url, "-o", "binary.bin")
	err = mimic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mimic_src_url := "https://github.com/MycroftAI/mimic1/archive/refs/tags/1.3.0.1.src.tar.gz"
	mimic_cmd_src := exec.Command("curl", "-L", mimic_src_url, "-o", "source.tar.gz")
	err = mimic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mimic_cmd_direct := exec.Command("./binary")
	err = mimic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: portaudio")
	exec.Command("latte", "install", "portaudio").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
