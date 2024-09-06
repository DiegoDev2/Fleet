package main

// Srt - Secure Reliable Transport
// Homepage: https://www.srtalliance.org/

import (
	"fmt"
	
	"os/exec"
)

func installSrt() {
	// Método 1: Descargar y extraer .tar.gz
	srt_tar_url := "https://github.com/Haivision/srt/archive/refs/tags/v1.5.3.tar.gz"
	srt_cmd_tar := exec.Command("curl", "-L", srt_tar_url, "-o", "package.tar.gz")
	err := srt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	srt_zip_url := "https://github.com/Haivision/srt/archive/refs/tags/v1.5.3.zip"
	srt_cmd_zip := exec.Command("curl", "-L", srt_zip_url, "-o", "package.zip")
	err = srt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	srt_bin_url := "https://github.com/Haivision/srt/archive/refs/tags/v1.5.3.bin"
	srt_cmd_bin := exec.Command("curl", "-L", srt_bin_url, "-o", "binary.bin")
	err = srt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	srt_src_url := "https://github.com/Haivision/srt/archive/refs/tags/v1.5.3.src.tar.gz"
	srt_cmd_src := exec.Command("curl", "-L", srt_src_url, "-o", "source.tar.gz")
	err = srt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	srt_cmd_direct := exec.Command("./binary")
	err = srt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
