package main

// Tsduck - MPEG Transport Stream Toolkit
// Homepage: https://tsduck.io/

import (
	"fmt"
	
	"os/exec"
)

func installTsduck() {
	// Método 1: Descargar y extraer .tar.gz
	tsduck_tar_url := "https://github.com/tsduck/tsduck/archive/refs/tags/v3.38-3822.tar.gz"
	tsduck_cmd_tar := exec.Command("curl", "-L", tsduck_tar_url, "-o", "package.tar.gz")
	err := tsduck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tsduck_zip_url := "https://github.com/tsduck/tsduck/archive/refs/tags/v3.38-3822.zip"
	tsduck_cmd_zip := exec.Command("curl", "-L", tsduck_zip_url, "-o", "package.zip")
	err = tsduck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tsduck_bin_url := "https://github.com/tsduck/tsduck/archive/refs/tags/v3.38-3822.bin"
	tsduck_cmd_bin := exec.Command("curl", "-L", tsduck_bin_url, "-o", "binary.bin")
	err = tsduck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tsduck_src_url := "https://github.com/tsduck/tsduck/archive/refs/tags/v3.38-3822.src.tar.gz"
	tsduck_cmd_src := exec.Command("curl", "-L", tsduck_src_url, "-o", "source.tar.gz")
	err = tsduck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tsduck_cmd_direct := exec.Command("./binary")
	err = tsduck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: dos2unix")
exec.Command("latte", "install", "dos2unix")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: grep")
exec.Command("latte", "install", "grep")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: qpdf")
exec.Command("latte", "install", "qpdf")
	fmt.Println("Instalando dependencia: librist")
exec.Command("latte", "install", "librist")
	fmt.Println("Instalando dependencia: libvatek")
exec.Command("latte", "install", "libvatek")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: srt")
exec.Command("latte", "install", "srt")
}
